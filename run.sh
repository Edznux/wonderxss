#!/usr/bin/env sh

# Print executed commands & exit on first error
set -ex;

# global variables
APP_NAME="wonderxss"

run_with_docker(){
    echo "Building docker container"
    docker build -t edznux/$APP_NAME deploy/docker/
    echo "Running docker container"
    docker run -v $(pwd)/$APP_NAME.conf:/$APP_NAME.conf:ro -it edznux/$APP_NAME
}

run_with_docker

exit 1


# Add the capability to run on port 80 and 443 without root privileges.
# Principle of least privilege. We don't need nor want root.
# sudo setcap 'cap_net_bind_service=+ep' $APP_NAME



