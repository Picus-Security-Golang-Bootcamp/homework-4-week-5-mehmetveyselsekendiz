package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name        *string
	Books		[]Book `gorm:"foreignKey:AuthorID;references:ID"`
}

type Book struct {
	gorm.Model
	Name        *string
	Page        int
	Stock       int
	Price       float32   `sql:"type:decimal(10,2);"`
	StockCode   string
	ISBN        string
	AuthorID    int
	Author      Author
}

func (Author) TableName() string {
	return "Author"
}

func (Book) TableName() string {
	return "Book"
}

func (a *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, CreatedAt : %s", a.ID, *a.Name, a.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, Page : %d, Stock : %d, Price : %g, StockCode : %s, ISNB : %s, AuthorID : %d, CreatedAt : %s",
	 b.ID, *b.Name, b.Page, b.Stock, b.Price, b.StockCode, b.ISBN, b.AuthorID, b.CreatedAt.Format("2006-01-02 15:04:05"))
}