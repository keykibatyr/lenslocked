package main

import (
	"html/template"
	"os"
)

type User struct{
	Name string
	Age int
	Bio string
}

type UserMeta struct{
	Visits int
}

func main() {
	t , err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Alisher",
		Age: 19,
		Bio: `<script>alert("ITS PATRICK")</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil{
		panic(err)
	}
}
