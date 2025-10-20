package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/keykibatyr/lenslocked/models"
)

type Users struct {
	Templates struct {
		New Template
	}

	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")

	u.Templates.New.ExecuteTemp(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Email: ", r.FormValue("email"))
	fmt.Fprintln(w, "Password: ", r.FormValue("password"))
	fmt.Fprintln(w, "Terms: ", r.FormValue("checkbox1"))

	err := r.ParseMultipartForm(10 << 20)
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
