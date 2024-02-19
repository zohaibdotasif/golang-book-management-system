package models

import (
	"golang-book-management-system/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `gorm:"" json:"author"`
	Publication string `gorm:"" json:"publication"`
}

type Books []*Book

func init() {
	config.Connect()
	db = config.GetDb()
	if err := db.AutoMigrate(&Book{}); err != nil {
		panic(err)
	}
}

func (b *Book) Create() error {
	result := db.Create(b).Scan(b)
	return result.Error
}

func (b *Books) List() error {
	result := db.Find(b)
	return result.Error
}

func (b *Book) GetByBookId(bookId int) error {
	result := db.Where("ID=?", bookId).First(b)
	return result.Error
}

func (b *Book) Update(bookId int, updates *Book) error {
	if updates.Name != "" {
		b.Name = updates.Name
	}
	if updates.Author != "" {
		b.Author = updates.Author
	}
	if updates.Publication != "" {
		b.Publication = updates.Publication
	}
	result := db.Save(b).Scan(b)
	return result.Error
}

func (b *Book) Delete(bookId int) error {
	result := db.Where("ID=?", bookId).First(b).Delete(b)
	return result.Error
}
