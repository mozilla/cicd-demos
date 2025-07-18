#!/bin/bash -ex
# Generate version.json file following our dockerflow standards
# "source" : "https://github.com/mozilla-services/Dockerflow",
# "version": "release tag or string for humans",
# "commit" : "<git hash>",
# "build"  : "uri to CI build job"
# "date"   : "build date of image"
# Exit immediately if a command exits with a non-zero status
set -e

# Variables
DATE_EPOCH=$(date +%s)
GITHUB_REPOSITORY_URL="${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY}"
GITHUB_ACTION_RUN_URL="${GITHUB_SERVER_URL}/${GITHUB_REPOSITORY}/actions/runs/${GITHUB_RUN_ID}"
if [ "$GITHUB_EVENT_NAME" == "pull_request" ]; then
    GITHUB_SHA=$(cat $GITHUB_EVENT_PATH | jq -r .pull_request.head.sha)
    GITHUB_REF_NAME=$(cat $GITHUB_EVENT_PATH | jq -r .pull_request.head.sha)
fi

# Generate version.json
cat <<EOF > version.json
{
  "source": "$GITHUB_REPOSITORY_URL",
  "version": "$GITHUB_REF_NAME",
  "commit": "$GITHUB_SHA",
  "build": "$GITHUB_ACTION_RUN_URL",
  "date": "$DATE_EPOCH"
}
EOF
