package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name        string  `json:"name"`
	Books		[]Book  `json:"books"`
}

type Book struct {
	gorm.Model
	Name        string   `json:"name"`
	Page        int      `json:"page"`
	Stock       int      `json:"stock"`
	Price       float32  `sql:"type:decimal(10,2);json:price"`
	StockCode   string   `json:"stock_code"`
	ISBN        string   `json:"isbn"`
	AuthorID    uint     `json:"author_id"`
	Author      Author   `json:"author"`
}

func (Author) TableName() string {
	return "Author"
}

func (Book) TableName() string {
	return "Book"
}

func (a *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, CreatedAt : %s", a.ID, a.Name, a.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, Page : %d, Stock : %d, Price : %g, StockCode : %s, ISNB : %s, AuthorID : %d, CreatedAt : %s",
	 b.ID, b.Name, b.Page, b.Stock, b.Price, b.StockCode, b.ISBN, b.AuthorID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}