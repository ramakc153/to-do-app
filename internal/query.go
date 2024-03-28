package internal

import (
	"fmt"
	"log"
	"os"
	"to-do-app/config"

	"github.com/joho/godotenv"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func SqlQuery() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username, password, address, port, db_name := os.Getenv("USNAME"), os.Getenv("PASSWD"), os.Getenv("DB_IP"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")
	db, err := config.Connect(username, password, address, port, db_name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}
