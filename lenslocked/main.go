package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/keykibatyr/lenslocked/controllers"
	"github.com/keykibatyr/lenslocked/views"
)

type Info struct {
	Name string
}

func main() {
	// var router Router
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	tpl, err := views.Parse("templates/home.gohtml")
	if err != nil {
		panic(err)
	}

	router.Get("/", controllers.StaticHandler(tpl, nil))

	contacttpl, err := views.Parse("templates/contact.gohtml")

	if err != nil {
		panic(err)
	}

	router.Get("/contact/{userID}", controllers.StaticHandler(contacttpl, func(r *http.Request) interface{} {
		return &Info{Name: chi.URLParam(r, "userID")}
	}))

	faqtpl, err := views.Parse("templates/faq.gohtml")

	if err != nil {
		panic(err)
	}

	router.Get("/faq", controllers.StaticHandler(faqtpl, nil))

	fmt.Println("Listening to port :7000...")
	http.ListenAndServe(":7000", router)
}
