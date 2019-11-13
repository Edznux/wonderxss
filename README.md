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


# Fact
Steve was blind, this may help.