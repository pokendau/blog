package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"blog/api"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Error getting the current working directory, %v", err)
	}

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
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

	http.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		fp_layout := filepath.Join(cwd, "web", "layouts", "base.html")
		fp_page := filepath.Join(cwd, "web", "pages", "articles.html")

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

	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != os.Getenv("USERNAME") || password != os.Getenv("PASSWORD") {
			w.Header().Set("WWW-AUTHENTICATE", "Basic realm=\"restricted\"")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("kebab")
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
