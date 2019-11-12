#!/usr/bin/env sh

APP_NAME="wonder-xss"

chmod +x $APP_NAME
# Add the capability to run on port 80 and 443 without root privileges.
# Principle of least privilege. We don't need nor want root.
sudo setcap 'cap_net_bind_service=+ep' $APP_NAME

./$APP_NAME
