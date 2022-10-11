package main

import (
	"log"
	"net/http"
	"text/template"
)

func scrapperHandler(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}

	page := scrapAndDo(URL)

	// dump results
	// b, err := json.Marshal(page)
	// if err != nil {
	// 	log.Println("failed to serialize response:", err)
	// 	return
	// }
	var indexTmpl = "./templates/index.html"
	tmpl := template.Must(template.ParseFiles(indexTmpl))

	tmpl.Execute(w, page)
}
