package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello Goland!")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a about page.")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("サーバーが起動できませんでした。")
	}
}
