# wonderxss

Blind-XSS tool.
Work in progress

Features:
- [x] 0 runtime dependencies
- [ ] Realtime (websocket)
- [x] Extensible
- [ ] Notification services
  - [x] Slack
  - [ ] Email
  - [ ] Web UI (websocket)
- [ ] Payload Generator
- [ ] One click deploy


## Deploy & Run

Deployment methods are WIP
Currently, you will have to build the execute the project:

```bash
git clone https://github.com/Edznux/wonderxss
cd wonderxss
export WONDERXSS_HTTPS=true
export WONDERXSS_DOMAIN=example.com
go build
sudo setcap 'cap_net_bind_service=+ep' ./wonderxss
./wonderxss
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

GET  /loots
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
curl $DOMAIN/api/v1/alias
# Get all loots
curl $DOMAIN/api/v1/loots
```