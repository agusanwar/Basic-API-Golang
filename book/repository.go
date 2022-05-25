package book

import "gorm.io/gorm"

// devinisi interface untuk repository
type Reposiroty interface{
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book Book) (Book, error)
}

// Struct Book
type repository struct {
	// Access database for gorm DB
	db *gorm.DB
}

func NewReposiroty(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error){
	var books [] Book
	
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindByID(ID int) (Book, error){
	var book Book

	err := r.db.Find(&book, ID ).Error
	return book, err
}

func (r *repository) Create(book Book) (Book, error){
	err := r.db.Create(&book).Error
	return book, err
}