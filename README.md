# wonderxss

Blind-XSS tool.
**Work in progress**

Features:
- [x] 0 runtime dependencies
- [ ] Realtime (websocket)
- [ ] Authentication
  - [ ] JWT
  - [ ] Oauth2
- [x] Extensible
- [ ] Notification services
  - [x] Slack
  - [x] Discord
  - [ ] Email
  - [ ] Web UI (websocket)
- [x] Payload Generator
- [ ] One click deploy
  - [ ] Terraform
  - [ ] Docker
  - [ ] Vagrant
  - [ ] Github release (binary)

Roadmap:
- [] DNS Listenner
- [] Authorization

## How to

1. Start the application (edit the `wonderxss.conf` file if needed)
2. Create a new payload by visiting `https://yourdomain.tld/editor` and filling up the form.
   1. Add an optional alias. Let's say `a` so you get a short injection payload.
3. Go back to the home page and select the payload you want to create an injection for.
   1. Check or uncheck the use of the subdomain
   2. Copy paste the injection
4. Inject your payload in some vulnerable application
5. Wait to the application to trigger your payload.
   1. If you have set up a notification system, you will also get a notification on trigger

## Deploy & Run

Deployment methods are WIP
Currently, you will have to build the execute the project:

```bash
git clone https://github.com/Edznux/wonderxss
cd wonderxss
export WONDERXSS_DOMAIN=example.com
go build
# Only sets the cap_net_bind_service to bind.
# You can also directly use `sudo ./wonderxss`
sudo setcap 'cap_net_bind_service=+ep' ./wonderxss
./wonderxss
```

### Frontend

The frontend UI is really WIP. You will need to run the following commands to get it running:
```
cd ui/wonderxss/
npm install

# this next command will serve the app on a different port.
# You may have CORS and/or port issue
npm start # If you get some permission error, run as sudo ;)

# Later, we will bundle the frontend with the rest of the app. We will use:
npm run build
```

## API

Prefix: `/api/v1/`

Routes:

```bash
GET  /api/v1/healthz

GET  /api/v1/payloads
GET  /api/v1/payloads/{id}
POST /api/v1/payloads

GET  /api/v1/aliases
GET  /api/v1/aliases/{id}
POST /api/v1/aliases

GET  /api/v1/collectors
```

Examples:

```bash
export DOMAIN=example.com
# Create a new payload
curl -X POST $DOMAIN/api/v1/payloads --data '{"name":"Test 1", "content":"alert(1)"}'
# Create a new alias for a paylaod
curl -X POST $DOMAIN/api/v1/aliases --data '{"alias":"a", "payload_id":"b4221cb8-5ff8-4677-8a16-f567edd9d58d"}'
# Get all payloads
curl $DOMAIN/api/v1/payloads
# Get all aliases
curl $DOMAIN/api/v1/aliases
# Get all loots
curl $DOMAIN/api/v1/collectors
```


## Configuration

Example configuration file (`cp wonderxss.conf.example wonderxss.conf`)

```toml
# Domain to listen on. (Usefull only for sudomains paylaod aliases)
domain = "localhost"
# Type of database
database = "sqlite"
# true will enable listening on HTTPS (without a reverse proxy). Disabled by default.
# This will require certificates:
# self signed RSA: `openssl req -x509 -nodes -newkey rsa:2048 -keyout server.key -out server.crt -days 3650`
# Use letsencrypt or other trusted provider for production environment
standalone_https = true

# Lists of notification systems
[notifications]
  # slack-webhook is an arbitrary name. Use whatever you prefer.
  # TODO:
  # You can use multiple times the same kind of notification
  # For example, multiple slack channels
  [notifications.slack-webhook]
    name    = "slack"
    # Change this token
    token   = "test"
    # Globaly disable this kind of notifications
    enabled = true

# Databases definitions
[storages]
  [storages.sqlite]
    adapter = "sqlite"
    # name of the sqlite file
    file    = "db.sqlite"

```

Environment variables:

`WONDERXSS_DOMAIN`: Domain to listen on. Usefull for wildcard subdomain.
`WONDERXSS_HTTPS`: true will enable listening on HTTPS (without a reverse proxy). Disabled by default.
This will require certificates:
(self signed RSA: `openssl req -x509 -nodes -newkey rsa:2048 -keyout server.key -out server.crt -days 3650`)
(Use letsencrypt or other trusted provider for prod)
`WONDERXSS_STORE`: Select the database used by wonderxss. Default: sqlite
