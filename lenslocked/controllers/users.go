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
	SessionService *models.SessionService
}

	func (u Users) New(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Email   string
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
	
	// fmt.Fprintf(w, "User was created: %v", user)
	// fmt.Fprintln(w, "Terms: ", r.FormValue("checkbox1"))
	// fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)

	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return 
	}

	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
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
		return
	}

	session, err := u.SessionService.Create(user.ID)
	
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return 
	}

	setCookie(w, CookieSession, session.Token)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := readCookie(r, CookieSession)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}

	user , err := u.SessionService.User(tokenCookie)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	
	fmt.Fprintf(w, "Current User is: %v", user)
}

