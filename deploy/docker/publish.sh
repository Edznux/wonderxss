#!/usr/bin/env bash

# Print executed commands & exit on first error
set -ex;

# https://stackoverflow.com/a/29436423
function yes_or_no {
    while true; do
        read -p "$* [y/n]: " yn
        case $yn in
            [Yy]*) return 0  ;;  
            [Nn]*) echo "Aborted" ; return  1 ;;
        esac
    done
}

# Ensure we provide a version number
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <version>"
fi

VERSION=$1

# Login to dockerhub or any registry
docker login

echo "Building..."
docker build -t edznux/wonderxss:$VERSION .

yes_or_no "Publish $VERSION to dockerhub ?" || exit 1
docker push edznux/wonderxss:$VERSION