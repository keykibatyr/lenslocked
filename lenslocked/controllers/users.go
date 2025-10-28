package controllers

import (
	"fmt"	
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/keykibatyr/lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		Signin Template
	}

	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email     string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.ExecuteTemp(w, r, data)
	fmt.Println("CSRF token:", csrf.Token(r))
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := u.UserService.Create(email, password)
	if err != nil {
		http.Error(
			w,
			"Could not create the user",
			http.StatusBadRequest,
		)
	}

	fmt.Fprintf(w, "User was created: %v", user)
	fmt.Fprintln(w, "Terms: ", r.FormValue("checkbox1"))

	// adding the picture
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(
			w,
			"File has to be <=10mb",
			http.StatusBadRequest,
		)
	}

	file, handler, err := r.FormFile("profile_picture")
	if err != nil {
		http.Error(
			w,
			"Could not the parse the file",
			http.StatusInternalServerError,
		)
	} else {
		defer file.Close()
		log.Print("the file was succesfuly parsed")

		dst, err := os.Create(handler.Filename)
		if err != nil {
			log.Print("could not create a file")
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			log.Print("could not copy the file to the destination")
		}
	}
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")

	u.Templates.Signin.ExecuteTemp(w, r, data)
}

func (u Users) ProcessSignin(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}

	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")

	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		http.Error(
			w,
			"Could not authenticate",
			http.StatusBadRequest,
		)
	}

	cookie := http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "User successfuly logged in: %v", user)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")
	if err != nil {
		fmt.Fprint(w, "Cookie doesnt exist")
		return
	}
	fmt.Fprintf(w, "User cookie: %v\n", email.Value)
}
