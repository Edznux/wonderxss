package websocket

import (
	"fmt"
	"net/http"

	"github.com/edznux/wonderxss/events"
	"github.com/gorilla/websocket"
)

type WSApi struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
}

func New() *WSApi {
	wsapi := WSApi{}
	wsapi.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	wsapi.clients = make(map[*websocket.Conn]bool)
	go wsapi.dispatcher()

	return &wsapi
}

func (wsapi *WSApi) Handle(w http.ResponseWriter, req *http.Request) {
	//allow cross origin websocket
	wsapi.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsapi.upgrader.Upgrade(w, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	wsapi.clients[conn] = true

}

func (wsapi *WSApi) dispatcher() {
	var err error
	ch := events.Events.Sub(events.TOPIC_PAYLOAD_DELIVERED)
	for {
		for conn := range wsapi.clients {
			if msg, ok := <-ch; ok {
				fmt.Println("Sending data to websocket !")
				err = conn.WriteJSON(msg)
				if err != nil {
					conn.Close()
					delete(wsapi.clients, conn)
				}
			}
		}
	}
}
