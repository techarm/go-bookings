package main

import (
	"github.com/techarm/go-bookings/pkg/handler"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	log.Println("start server and listen on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("サーバーが起動できませんでした。", err)
	}
}
