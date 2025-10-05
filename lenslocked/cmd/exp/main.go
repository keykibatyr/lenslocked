package main

import (
	"html/template"
	"os"
)

type Address struct{
	Street string
	HouseNum int
} 
type User struct{
	Name string
	Age int
	Bio string
	Address *Address
}
type UserMeta struct{
	Visits int
}

func main() {
	t , err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	address := &Address{
		Street: "Umai Ana",
		HouseNum: 15,
	}

	user := &User{
		Name: "Alisher",
		Age: 19,
		Bio: `<script>alert("ITS PATRICK")</script>`,
		Address: address,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil{
		panic(err)
	}
}
