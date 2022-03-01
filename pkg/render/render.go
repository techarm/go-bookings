package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap {}

func Render(w http.ResponseWriter, name string) {
	cache, err := createTemplateCache()
	if err != nil {
		log.Fatal("Error getting template cache", err)
	}

	t, ok := cache[name]
	if !ok {
		log.Fatalln("Can not parse template file", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln("Can not execute template file", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		t, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		match, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return cache, err
		}

		if len(match) > 0 {
			t, err = t.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = t
	}

	return cache, nil
}