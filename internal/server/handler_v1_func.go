// Copyright 2023 Google LLC
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

// This file includes functions used by handler_v1 that can be mocked in tests.

package server

import (
	"net/http"

	pbv1 "github.com/datacommonsorg/mixer/internal/proto/v1"
	"github.com/datacommonsorg/mixer/internal/server/resource"
	"github.com/datacommonsorg/mixer/internal/server/v1/info"
	"github.com/datacommonsorg/mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var apiPaths = map[string]string{
	protoName(&pbv1.BulkVariableInfoRequest{}):      "/v1/bulk/info/variable",
	protoName(&pbv1.BulkVariableGroupInfoRequest{}): "/v1/bulk/info/variable-group",
}

var localBulkVariableInfoFunc = info.BulkVariableInfo

var remoteBulkVariableInfoFunc = func(
	s *Server,
	remoteReq *pbv1.BulkVariableInfoRequest,
) (*pbv1.BulkVariableInfoResponse, error) {
	return fetchRemote(s.metadata, s.httpClient,
		remoteReq,
		&pbv1.BulkVariableInfoResponse{},
	)
}

var localBulkVariableGroupInfoFunc = info.BulkVariableGroupInfo

var remoteBulkVariableGroupInfoFunc = func(
	s *Server,
	remoteReq *pbv1.BulkVariableGroupInfoRequest,
) (*pbv1.BulkVariableGroupInfoResponse, error) {
	return fetchRemote(s.metadata, s.httpClient,
		remoteReq,
		&pbv1.BulkVariableGroupInfoResponse{},
	)
}

func fetchRemote[T protoreflect.ProtoMessage](metadata *resource.Metadata, httpClient *http.Client, in protoreflect.ProtoMessage, out T) (T, error) {
	if apiPath, ok := apiPaths[protoName(in)]; !ok {
		return out, status.Error(codes.Internal, "Unsupported request.")
	} else {
		return out, util.FetchRemote(metadata, httpClient, apiPath, in, out)
	}
}

func protoName(in protoreflect.ProtoMessage) string {
	return string(in.ProtoReflect().Descriptor().FullName())
}
