package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
	"github.com/keykibatyr/lenslocked/controllers"
	"github.com/keykibatyr/lenslocked/migrations"
	"github.com/keykibatyr/lenslocked/models"
	"github.com/keykibatyr/lenslocked/templates"
	"github.com/keykibatyr/lenslocked/views"
)

type Info struct {
	Name string
}

type config struct {
	PSQL models.PostgresConfig
	SMTP models.SMTPConfig
	CSRF struct {
		Key    string
		Secure bool
	}
	Server struct {
		Address string
	}
}

func loadENVconfig() (config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	cfg.PSQL = models.DefaultConfig()

	cfg.CSRF.Key = "qRWLtI8k0q2kZ28nNsG32byMQoqOVmfKOhmLZgv6AD0"
	cfg.CSRF.Secure = false

	cfg.SMTP.Host = os.Getenv("SMTP_HOST")
	portSMTP, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic(err)
	}
	cfg.SMTP.Port = portSMTP
	cfg.SMTP.Username = os.Getenv("SMTP_USERNAME")
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")

	cfg.Server.Address = ":3000"

	return cfg, nil
}

func main() {
	cfg, err := loadENVconfig()
	if err != nil {
		panic(err)
	}
	// Opening the DB
	db, err := models.Open(cfg.PSQL)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.PSQL.String())

	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	userService := &models.UserService{
		DB: db,
	}

	sessionService := &models.SessionService{
		DB: db,
	}

	pwResetService := &models.PasswordResetService{
		DB: db,
	}

	emailServcie := models.NewEmailService(cfg.SMTP)


	umw := controllers.UserMiddleware{
		SessionService: sessionService,
	}

	csrfMW := csrf.Protect(
		[]byte(cfg.CSRF.Key),
		csrf.Secure(cfg.CSRF.Secure),
		csrf.Path("/"),
	)

	usersC := controllers.Users{
		UserService:    userService,
		SessionService: sessionService,
		PasswordResetService: pwResetService,
		EmailService: emailServcie,
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
	router.Route("/users/me", func(ro chi.Router) {
		ro.Use(umw.RequireUser)
		ro.Get("/", usersC.CurrentUser)
	})

	fmt.Printf("Listening to port %s...", cfg.Server.Address)
	http.ListenAndServe(cfg.Server.Address, router)
}
