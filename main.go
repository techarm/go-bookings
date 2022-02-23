package main

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render(w, "about.page.html")
}

func render(w http.ResponseWriter, name string) {
	t, err := template.ParseFiles("./templates/" + name)
	if err != nil {
		log.Fatalln("Can not parse template file", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln("Can not execute template file", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("サーバーが起動できませんでした。", err)
	}
}
