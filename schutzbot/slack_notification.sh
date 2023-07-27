#!/bin/bash

set -eux

if [ -z "${SLACK_WEBHOOK_URL:-}" ]; then
    echo "INFO: Variable SLACK_WEBHOOK_URL is undefined"
    exit 0
fi

COMPOSE_ID=$(cat COMPOSE_ID)
COMPOSER_NVR=$(cat COMPOSER_NVR)
MESSAGE="\"Nightly pipeline execution on *$COMPOSE_ID* with *$COMPOSER_NVR* finished with status *$1* $2 \n QE: @atodorov, @jrusz\n Link to results: $CI_PIPELINE_URL\n For edge testing status please see https://url.corp.redhat.com/edge-pipelines \""

curl \
    -X POST \
    -H 'Content-type: application/json' \
    --data '{"text": "test", "blocks": [ { "type": "section", "text": {"type": "mrkdwn", "text":'"$MESSAGE"'}}]}' \
    "$SLACK_WEBHOOK_URL"
