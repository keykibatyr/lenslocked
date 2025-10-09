package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/keykibatyr/lenslocked/controllers"
	"github.com/keykibatyr/lenslocked/templates"
	"github.com/keykibatyr/lenslocked/views"
)

type Info struct {
	Name string
}

func main() {
	// var router Router
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	tpl := views.Must(views.ParseFileSys(
		templates.FS, "home.gohtml", "tailwind.gohtml"))

	router.Get("/", controllers.StaticHandler(tpl, nil))

	contacttpl := views.Must(views.ParseFileSys(
		templates.FS, "contact.gohtml", "tailwind.gohtml"))

	router.Get("/contact/{userID}", controllers.StaticHandler(
		contacttpl, func(r *http.Request) interface{} {
		return &Info{Name: chi.URLParam(r, "userID")}
	}))

	faqtpl := views.Must(views.ParseFileSys(
		templates.FS,"faq.gohtml", "tailwind.gohtml"))

	router.Get("/faq", controllers.FAQ(faqtpl))

	fmt.Println("Listening to port :7000...")
	http.ListenAndServe(":7000", router)
}
