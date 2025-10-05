package main

import (
	"fmt"
	// "html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/keykibatyr/lenslocked/views"
)

type Info struct {
	Name string
}

func executeTemplate(w http.ResponseWriter, path string, data interface{}) { // I could add interafce instead of *Info. Probably will do that
	t, err := views.Parse(path, data)
	if err != nil {
		log.Printf("parsing template: %v", path)
		http.Error(
			w,
			"Failed at parsing the template",
			404,
		)
		return
	}

	t.ExecuteTemp(w)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	executeTemplate(w, "templates/home.gohtml",nil)

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
