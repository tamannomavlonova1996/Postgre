package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {

	db, err := sql.Open("pgx", "postgres://app:pass@localhost:5432/db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	//Установка данных
	//insert, err := db.Query(`INSERT INTO users (name,age) VALUES ('Alex', 35)`)
	//if err != nil {
	//log.Println(err)
	//}
	//defer insert.Close()

	//Выборка данных
	res, err := db.Query(`Select name, age FROM users`)
	if err != nil {
		log.Println(err)
	}
	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}
}
