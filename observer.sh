#!/bin/bash

DIRECTORY_TO_OBSERVE=$(dirname "$0")  # current script directory
echo "Watching ${DIRECTORY_TO_OBSERVE} for changes in *.go files"
while true; do
    change=$(find "${DIRECTORY_TO_OBSERVE}" -name "*.go" -print0 | xargs -0 inotifywait -e modify,create,delete,move --format '%w%f')
    if [[ $change = *.go ]]; then
        echo "Detected ${change} - restarting docker container"
        docker-compose restart app-api
        docker-compose restart app-redirect
    fi
done

#sudo apt-get install inotify-tools