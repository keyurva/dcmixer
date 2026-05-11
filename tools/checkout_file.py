import subprocess

# Run git checkout on the file directly using --git-dir and --work-tree
cmd = [
    "git", 
    "--git-dir=/Users/keyurs/projects/dc/dcimport/.git", 
    "--work-tree=/Users/keyurs/projects/dc/dcimport", 
    "checkout", 
    "pipeline/util/src/main/java/org/datacommons/ingestion/util/GraphReader.java"
]
res = subprocess.run(cmd, capture_output=True, text=True)
print(f"Exit code: {res.returncode}")
print(f"Stdout: {res.stdout}")
print(f"Stderr: {res.stderr}")
