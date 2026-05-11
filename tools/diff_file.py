import subprocess
cmd = [
    "git", 
    "--git-dir=/Users/keyurs/projects/dc/dcimport/.git", 
    "--work-tree=/Users/keyurs/projects/dc/dcimport", 
    "diff", 
    "pipeline/util/src/main/java/org/datacommons/ingestion/util/GraphReader.java"
]
res = subprocess.run(cmd, capture_output=True, text=True)
print(res.stdout)
