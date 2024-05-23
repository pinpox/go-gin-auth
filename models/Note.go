package models

import (
	"errors"
	// "html"
	// "strings"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model

	User   User
	UserID int

	Title  string `gorm:"size:255;not null;" json:"title"`
	Text   string `gorm:"size:255;not null;" json:"text"`
}

// GetNoteByID retrieves a user by ID from the database
func GetNoteByID(uid uint) (Note, error) {
	var n Note

	if err := DB.First(&n, uid).Error; err != nil {
		return n, errors.New("Note not found!")
	}

	return n, nil
}

// SaveNote creates a new user in the database
func (n *Note) SaveNote() (*Note, error) {
	var err error
	err = DB.Model(&n).Create(&n).Error
	if err != nil {
		return &Note{}, err
	}
	return n, nil
}
