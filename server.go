package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func scrapperHandler(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("missing URL argument")
		visitUrl := "https://asuu-news-update.blogspot.com"
		msg := fmt.Sprintf(`There's nothing here, please visit %s and come back here with the following format to enable ads less experience

""""

https://asuu-news-production.up.railway.app/?url=PAGE_URL


Example:

https://asuu-news-production.up.railway.app/?url=https://asuu-news-update.blogspot.com/2022/10/latest-update-on-asuu-resumption_11.html

""""
		`, visitUrl)
		w.Write([]byte(msg))
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
