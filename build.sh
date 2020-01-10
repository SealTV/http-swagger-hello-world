#!/bin/bash
go build .

set -e
images=$(echo $IMAGES | tr " " "\n")

for image in $images
do
    echo $image
    docker build -t $image .
    if $PUSH_IMAGE
    then
        docker push $image
    fi
done