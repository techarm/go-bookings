package render

import (
	"html/template"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, name string) {
	t, err := template.ParseFiles("./templates/" + name)
	if err != nil {
		log.Fatalln("Can not parse template file", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln("Can not execute template file", err)
	}
}
