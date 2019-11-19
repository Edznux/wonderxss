#!/usr/bin/env sh

run_with_doker()
exit 0

APP_NAME="wonderxss"

chmod +x $APP_NAME
# Add the capability to run on port 80 and 443 without root privileges.
# Principle of least privilege. We don't need nor want root.
sudo setcap 'cap_net_bind_service=+ep' $APP_NAME

./$APP_NAME


run_with_doker(){
    # docker build -t edznux/wonderxss deploy/docker/
    docker run -v wonderxss.conf:/wonderxss.conf -it edznux/wonderxss
}
