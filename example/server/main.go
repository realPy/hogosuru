package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {

			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {

	http.HandleFunc("/echo", echo)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/app/") {

			http.ServeFile(w, r, "example/static/loading.html")
		} else {
			http.ServeFile(w, r, "example/static/"+r.URL.Path[1:])
		}

	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
