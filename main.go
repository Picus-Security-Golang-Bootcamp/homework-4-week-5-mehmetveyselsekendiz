package main

import (
	"log"

	"github.com/joho/godotenv"
	postgres "github.com/mehmetveyselsekendiz/homework-4-week-5/common/db"
	"github.com/mehmetveyselsekendiz/homework-4-week-5/model"
)

func main(){

	//Set environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	db.AutoMigrate(&model.Author{}, &model.Book{})
}