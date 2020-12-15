package main

import (
	"io"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	log.Infof("Request %V", req)
	io.WriteString(w, "hello, world!\n")
}

func echoHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Receive: %s\n", msg[:n])

	m, err := ws.Write(msg[:n])
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Send: %s\n", msg[:m])
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting...")
	http.Handle("/hello", http.HandlerFunc(HelloHandler))
	http.Handle("/ws", websocket.Handler(echoHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
