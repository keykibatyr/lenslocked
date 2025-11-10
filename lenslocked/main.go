package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/keykibatyr/lenslocked/controllers"
	"github.com/keykibatyr/lenslocked/migrations"
	"github.com/keykibatyr/lenslocked/models"
	"github.com/keykibatyr/lenslocked/templates"
	"github.com/keykibatyr/lenslocked/views"
)

type Info struct {
	Name string
}

func main() {
		// Opening the DB
	cfg := models.DefaultConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.String())

	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}	

	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}


	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	csrf_key := "qRWLtI8k0q2kZ28nNsG32byMQoqOVmfKOhmLZgv6AD0"
	csrfMW := csrf.Protect(
		[]byte(csrf_key), 
		csrf.Secure(false), 
		csrf.Path("/"),
	)
	

	usersC := controllers.Users{
		UserService: &userService,
		SessionService: &sessionService,
	}
 
	usersC.Templates.New = views.Must(views.ParseFileSys(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))

	usersC.Templates.Signin = views.Must(views.ParseFileSys(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))

	usersC.Templates.ForgotPassword = views.Must(views.ParseFileSys(
		templates.FS, "forgot-pw.gohtml", "tailwind.gohtml"))

	// var router Router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(csrfMW)
	router.Use(umw.SetUser)

	tpl := views.Must(views.ParseFileSys(
		templates.FS, "home.gohtml", "tailwind.gohtml"))

	router.Get("/", controllers.StaticHandler(tpl))

	contacttpl := views.Must(views.ParseFileSys(
		templates.FS, "contact.gohtml", "tailwind.gohtml"))

	router.Get("/contact", controllers.StaticHandler(contacttpl))

	faqtpl := views.Must(views.ParseFileSys(
		templates.FS, "faq.gohtml", "tailwind.gohtml"))


	router.Get("/faq", controllers.FAQ(faqtpl))

	router.Get("/signUp", usersC.New)
	router.Post("/signUp", usersC.Create)
	router.Get("/signIn", usersC.SignIn)
	router.Post("/signIn", usersC.ProcessSignin)
	router.Post("/signOut", usersC.ProcessSignOut)
	router.Get("/forgot-pw", usersC.ForgotPassword)
	router.Post("/forgot-pw", usersC.ProcessForgotPassword)
	router.Route("/users/me", func(ro chi.Router){
		ro.Use(umw.RequireUser)
		ro.Get("/", usersC.CurrentUser)
	})


	fmt.Println("Listening to port :7000...")
	http.ListenAndServe(":7000", router)
}
