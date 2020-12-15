package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"
var url = "ws://localhost:8080/ws"

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for ; ; i++ {
		message := []byte(fmt.Sprintf("massage #%v", i))
		_, err = ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Sprintf()
		// format string, a ...interface{}
		log.Infof("Send: %s", message)

		var msg = make([]byte, 512)
		n, err := ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("Receive: %s", msg[:n])
	}

}
