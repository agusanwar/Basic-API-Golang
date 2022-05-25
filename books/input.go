package books

import "encoding/json"

// POST API
type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
	Desc  string      `json:"desc" binding:"required"`
}