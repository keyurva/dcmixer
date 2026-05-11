// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	adminpb "cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
	instance "cloud.google.com/go/spanner/admin/instance/apiv1"
	instancepb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
)

func main() {
	project := flag.String("project", "datcom-store", "Spanner project ID")
	srcInstance := flag.String("src_instance", "dc-kg-test", "Source instance ID")
	destInstance := flag.String("dest_instance", "dc-kg-test2", "Destination instance ID")
	destDb := flag.String("dest_db", "dc_graph_2026_04_30", "Destination database ID")
	flag.Parse()

	ctx := context.Background()

	// 1. Initialize Instance Admin Client
	instanceAdmin, err := instance.NewInstanceAdminClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create instance admin client: %v", err)
	}
	defer instanceAdmin.Close()

	destInstanceName := fmt.Sprintf("projects/%s/instances/%s", *project, *destInstance)

	// 2. Delete the existing instance to reset its edition
	fmt.Printf("Deleting existing instance %s to reset its edition...\n", *destInstance)
	err = instanceAdmin.DeleteInstance(ctx, &instancepb.DeleteInstanceRequest{
		Name: destInstanceName,
	})
	if err != nil {
		fmt.Printf("Delete returned error (it might not exist): %v\n", err)
	}

	// 3. Fetch source instance config to replicate exactly
	srcInstanceName := fmt.Sprintf("projects/%s/instances/%s", *project, *srcInstance)
	srcInst, err := instanceAdmin.GetInstance(ctx, &instancepb.GetInstanceRequest{
		Name: srcInstanceName,
	})
	if err != nil {
		log.Fatalf("Failed to get source instance %s: %v", srcInstanceName, err)
	}

	fmt.Printf("Replicating source instance config %s, edition %s, and node count %d...\n", 
		srcInst.Config, srcInst.Edition, srcInst.NodeCount)

	// 4. Recreate destination instance with 1 node
	fmt.Printf("Recreating new instance %s with 1 node to fit in quota...\n", *destInstance)
	op, err := instanceAdmin.CreateInstance(ctx, &instancepb.CreateInstanceRequest{
		Parent:     fmt.Sprintf("projects/%s", *project),
		InstanceId: *destInstance,
		Instance: &instancepb.Instance{
			Config:      srcInst.Config,
			DisplayName: *destInstance,
			NodeCount:   1, // Set to 1 node to fit in quota
			Edition:     srcInst.Edition,   // Set to Enterprise edition
		},
	})
	if err != nil {
		log.Fatalf("Failed to create instance %s: %v", *destInstance, err)
	}

	fmt.Println("Waiting for instance creation to complete...")
	_, err = op.Wait(ctx)
	if err != nil {
		log.Fatalf("Instance creation failed: %v", err)
	}
	fmt.Printf("Successfully created instance %s with 1 node!\n", *destInstance)

	// 5. Initialize Database Admin Client
	dbAdmin, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create database admin client: %v", err)
	}
	defer dbAdmin.Close()

	// 6. Create destination database
	fmt.Printf("Creating new database %s...\n", *destDb)
	dbOp, err := dbAdmin.CreateDatabase(ctx, &adminpb.CreateDatabaseRequest{
		Parent:          destInstanceName,
		CreateStatement: fmt.Sprintf("CREATE DATABASE `%s`", *destDb),
	})
	if err != nil {
		log.Fatalf("Failed to create database %s: %v", *destDb, err)
	}

	fmt.Println("Waiting for database creation to complete...")
	_, err = dbOp.Wait(ctx)
	if err != nil {
		log.Fatalf("Database creation failed: %v", err)
	}
	fmt.Printf("Successfully created database %s\n", *destDb)

	// 7. Execute DDL to create table and indexes
	fmt.Println("Updating database schema with FlatTimeSeries table and indexes...")
	ddlOp, err := dbAdmin.UpdateDatabaseDdl(ctx, &adminpb.UpdateDatabaseDdlRequest{
		Database: fmt.Sprintf("projects/%s/instances/%s/databases/%s", *project, *destInstance, *destDb),
		Statements: []string{
			`CREATE TABLE FlatTimeSeries (
			  variable_measured STRING(1024) NOT NULL,
			  entity1 STRING(1024) NOT NULL,
			  entity2 STRING(1024),
			  entity3 STRING(1024),
			  dates_and_values JSON,
			  time_series_attributes JSON,
			  id STRING(1024) NOT NULL
			) PRIMARY KEY(variable_measured, entity1, id)`,
			`CREATE INDEX FlatTimeSeriesEntity1 ON FlatTimeSeries(entity1)`,
			`CREATE INDEX FlatTimeSeriesEntity2 ON FlatTimeSeries(entity2)`,
			`CREATE INDEX FlatTimeSeriesEntity3 ON FlatTimeSeries(entity3)`,
		},
	})
	if err != nil {
		log.Fatalf("Failed to enqueue DDL: %v", err)
	}

	fmt.Println("Waiting for DDL execution to complete...")
	err = ddlOp.Wait(ctx)
	if err != nil {
		log.Fatalf("DDL execution failed: %v", err)
	}
	fmt.Println("Successfully created table and secondary indexes!")
}
