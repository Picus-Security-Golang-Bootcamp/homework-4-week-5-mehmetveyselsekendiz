package book

import (
	"github.com/mehmetveyselsekendiz/homework-4-week-5/model"
	"gorm.io/gorm"
	"errors"
)


func BatchCreateBook(db *gorm.DB){

	books := []model.Book{
		{Name: "Book1", Page: 11, Stock: 111, Price: 11.11, StockCode: "B1", ISBN: "ISBN-B1", Author: model.Author{Name:"Author1"}},
		{Name: "Book2", Page: 22, Stock: 222, Price: 22.22, StockCode: "B2", ISBN: "ISBN-B2", Author:model.Author{Name:"Author2"}}}
		
	db.Create(&books)

}

func GetBookById(db *gorm.DB, id int)(*model.Book, error){
	var book model.Book
	result := db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &book, nil
}

func GetBooks(db *gorm.DB)(*[]model.Book, error){
	var book []model.Book
	result := db.Find(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &book, nil
}

func CreateBook(db *gorm.DB, book *model.Book){
	db.Create(&book)
}