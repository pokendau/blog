package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"blog/api"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error getting the current working directory, %v", err)
	}

	fmt.Println("The server is running on port 7269")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fp_layout := filepath.Join(cwd, "web", "layouts", "base.html")
		fp_page := filepath.Join(cwd, "web", "pages", "homepage.html")

		funcmap := template.FuncMap{
			"WithComData": api.WithComData,
		}

		tmpl, err := template.New("base.html").Funcs(funcmap).ParseFiles(fp_layout, fp_page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir(filepath.Join(cwd, "web", "static"))),
		),
	)

	log.Fatal(http.ListenAndServe(":7269", nil))
}
