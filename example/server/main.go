package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/app/") {

			http.ServeFile(w, r, "example/static/routing.html")
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
