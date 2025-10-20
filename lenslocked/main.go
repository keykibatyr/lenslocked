package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/keykibatyr/lenslocked/controllers"
	"github.com/keykibatyr/lenslocked/models"
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

	router.Get("/", controllers.StaticHandler(tpl))

	contacttpl := views.Must(views.ParseFileSys(
		templates.FS, "contact.gohtml", "tailwind.gohtml"))

	router.Get("/contact", controllers.StaticHandler(contacttpl))

	faqtpl := views.Must(views.ParseFileSys(
		templates.FS,"faq.gohtml", "tailwind.gohtml"))

	router.Get("/faq", controllers.FAQ(faqtpl))
	
	db, err := models.Open(models.DefaultConfig())
	if err != nil {
		panic(err)
	}

	defer db.Close()

	userService := models.UserService{
		DB: db,
	}
	usersC := controllers.Users{
		UserService: &userService,
	}

	usersC.Templates.New = views.Must(views.ParseFileSys(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))

	router.Get("/signUp", usersC.New)
	router.Post("/users", usersC.Create)

	fmt.Println("Listening to port :7000...")
	http.ListenAndServe(":7000", router)
}
