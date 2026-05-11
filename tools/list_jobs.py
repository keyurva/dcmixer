import subprocess
import json
from datetime import datetime

# Get fresh token
token_proc = subprocess.run(["gcloud", "auth", "application-default", "print-access-token"], capture_output=True, text=True)
token = token_proc.stdout.strip()

# Fetch Dataflow jobs
curl_proc = subprocess.run([
    "curl", "-s", 
    "-H", f"Authorization: Bearer {token}",
    "https://dataflow.googleapis.com/v1b3/projects/datcom-store/locations/us-central1/jobs"
], capture_output=True, text=True)

print("=== Raw Output ===")
print(curl_proc.stdout)

