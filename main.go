package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func main() {
	fmt.Println("The server is running on port 7269")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fp_layout := path.Join("layouts", "base.html")
		fp_page := path.Join("pages", "homepage.html")

		tmpl, err := template.ParseFiles(fp_layout, fp_page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":7269", nil))
}
