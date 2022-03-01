package main

import (
	"github.com/techarm/go-bookings/pkg/handlers"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("start server and listen on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("サーバーが起動できませんでした。", err)
	}
}
