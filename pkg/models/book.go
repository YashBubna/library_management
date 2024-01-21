package models

import (
	"github.com/YashBubna/library_management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) //db.NewRecord and db.Create are all functions in gorm
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBookByID(ID int64) Book {
	var deleteBook Book
	db.Where("id=?", ID).Delete(&deleteBook)
	return deleteBook
}
