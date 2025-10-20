package main

import (
	// "database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/keykibatyr/lenslocked/models"
)

type PostgresConfig struct{
	Host string
	Port string
	User string
	Password string
	Database string
	SSLMode string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host = %s  port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, 
	cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	db, err := models.Open(models.DefaultConfig())
	if err != nil {
		panic(err)
	}

	defer db.Close()

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("keykibatyr@gmail.com", "Alisher2005")
	if err != nil {
		panic(err)
	}
	fmt.Print(user)
	// cfg := PostgresConfig{
	// 	Host: "localhost",
	// 	Port: "5432",
	// 	User: "baloo",
	// 	Password: "Alisher2005",
	// 	Database: "lenslocked",
	// 	SSLMode: "disable",
	// }

	// db, err := sql.Open("pgx", cfg.String())

	// if err != nil {
	// 	panic(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Connected!")
	// defer db.Close()

	// _, err = db.Exec(`
	// CREATE TABLE IF NOT EXISTS users(
	//  id SERIAL PRIMARY KEY,
	//  name TEXT,
	//  email TEXT UNIQUE NOT NULL
	//  );

	// CREATE TABLE IF NOT EXISTS orders(
	// id SERIAL PRIMARY KEY,
	// user_id INT NOT NULL,
	// amount INT,
	// description TEXT);
	// `)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Tables are created!")

	// name := "Alisher"
	// email := "keykibatyr@gmail.com"

	// _, err = db.Exec(`
	// INSERT INTO users (name, email)
	// VALUES($1, $2);`, name, email)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Info inserted")

	// name = "Bob"
	// email = "Bob@ver.org"
	// row := db.QueryRow(`
	// INSERT INTO users (name, email)
	// VALUES($1, $2) RETURNING id, name;`, name, email)
	
	// var id int 
	// err = row.Scan(&id, &name)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("User created id:%v, name: %v", id, name)

	// id := 1
	// row := db.QueryRow(`Select name, email 
	// From users where id = $1`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("recieved user; name = %v, email=%v", name, email)

	// userID := 1
	// for i:=1; i < 6; i++{
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%v", i)
	// 	_, err = db.Exec(`
	// 	Insert into orders (user_id, description, amount)
	// 	values($1, $2, $3);`, userID, desc, amount)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// }

	// userID := 1
	// type Order struct{
	// 	Id int
	// 	UserId int
	// 	Amount int
	// 	Description string
	// }

	// var objects []Order
	// rows , err := db.Query(`
	// SELECT id, amount, description
	// FROM orders Where user_id = $1`, userID)

	// if err != nil{
	// 	panic(err)
	// }

	// defer rows.Close()
	
	// for rows.Next(){
	// 	var object Order
	// 	object.UserId = userID
	// 	err = rows.Scan(&object.Id, &object.Amount, &object.Description)
	// 	if err != nil{
	// 		panic(err)
	// 	}
	// 	objects = append(objects, object)
	// }

	// fmt.Print(objects)
}
