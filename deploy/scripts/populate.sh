#!/usr/bin/env bash


set -e

BASE_DIR=`pwd`/deploy/scripts
INJECTIONS_FOLDER="injections"
PAYLOADS_FOLDER="payloads"

# Create all injections
for file in $BASE_DIR/$INJECTIONS_FOLDER/*
do
    echo "Injection: $file"
    if [[ -f "$file" ]]; then
        echo "Injection: $file"
        wonderxss injection create $file
    fi
done

# Create all payloads
for file in $BASE_DIR/$PAYLOADS_FOLDER/*
do
    if [[ -f "$file" ]]; then
        echo "Payload: $file"
        wonderxss payload create $file
    fi
done

