package main

import (
    "os"
    "io"
    "log"
    "net/http"

    "golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
    io.Copy(ws, ws)
}

func main() {
    //
    port := os.Getenv("PORT")
    if port == "" {
      log.Fatal("$PORT must be set")
    }
    //
    http.Handle("/echo", websocket.Handler(EchoServer))
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
