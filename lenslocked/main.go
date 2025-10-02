package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Info struct {
	Name string
}

func executeTemplate(w http.ResponseWriter, path string, info *Info) { // I could add intrafce instead of *Info. Probably will do that
	tpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(
			w,
			"Pasring the templatewas unsuccessful",
			http.StatusInternalServerError,
		)
		return
	}

	err = tpl.Execute(w, info)
	if err != nil {
		log.Printf("executing: %v", err)
		http.Error(
			w,
			"Failed at executing the page",
			http.StatusInternalServerError,
		)
		return
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, "templates/home.gohtml", nil)

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "userID")

	data := &Info{
		Name: user_id,
	}

	executeTemplate(w, "templates/contact.gohtml", data)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, "templates/faq.gohtml", nil)
}

func main() {
	// var router Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact/{userID}", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(
			w,
			"page not found",
			404,
		)
	})
	fmt.Println("Listening to port :7000...")
	http.ListenAndServe(":7000", r)
}
