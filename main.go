package main

import (
	"log"
	"encoding/json"

	"github.com/joho/godotenv"
	postgres "github.com/mehmetveyselsekendiz/homework-4-week-5/common/db"
	"github.com/mehmetveyselsekendiz/homework-4-week-5/model"
	"github.com/mehmetveyselsekendiz/homework-4-week-5/service/book"

	"net/http"
    "github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetBooksApi(w http.ResponseWriter, r *http.Request) {

	books, _ := book.GetBooks(db)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
  
	jsonResponse, err := json.Marshal(books)
	if err != nil {
	 return
	}
	w.Write(jsonResponse)
}

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

	router := mux.NewRouter()
	handlers.AllowedMethods([]string{"POST", "GET", "PUT", "PATCH"})

	router.HandleFunc("/books", GetBooksApi).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)


	//new_book := model.Book{Name: "BookNew", Page: 11, Stock: 111, Price: 11.11, StockCode: "B1", ISBN: "ISBN-B1", Author: model.Author{Name:"AuthorNew"}}
	//book.CreateBook(db,&new_book)
}