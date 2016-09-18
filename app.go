package main

import (
    "os"
    "log"
    "net/http"

    "gopkg.in/igm/sockjs-go.v2/sockjs"
)

func main() {
    //
    port := os.Getenv("PORT")
    if port == "" {
      log.WithField("PORT", port).Fatal("$PORT must be set")
    }
    //
    handler := sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler) 
    log.Fatal(http.ListenAndServe(":" + port, handler))
}

func echoHandler(session sockjs.Session) {
    for {
        if msg, err := session.Recv(); err == nil {
            session.Send(msg)
            continue
        }
        break
    }
}
