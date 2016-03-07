package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("name"))

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println(string(p))

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/", wsHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
