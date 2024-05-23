package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"ginjwtauth/utils"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Name     string `gorm:"size:255;not null;" json:"name"`
	Email    string `gorm:"size:255;not null;" json:"email"`
}

// GetUserByID retrieves a user by ID from the database
func GetUserByID(uid uint) (User, error) {
	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil
}

// PrepareGive removes sensitive information before sending user details
func (u *User) PrepareGive() {
	u.Password = ""
}

// GetUserByUsername retrieves a user by username from the database
func GetUserByUsername(username string) (User, error) {
	var user User

	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, errors.New("User not found")
		}
		return User{}, err
	}

	user.PrepareGive()

	return user, nil
}

// VerifyPassword compares the provided password with the hashed password
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// LoginCheck validates user credentials and generates a token
func LoginCheck(username, password string) (string, error) {
	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

// SaveUser creates a new user in the database
func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// BeforeSave is a callback function called before saving a user
func (u *User) BeforeSave() error {
	// Turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	// Remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

// UpdateUser updates an existing user in the database
func (u *User) UpdateUser() (*User, error) {
	if u.ID == 0 {
		return u, errors.New("User not found!")
	}

	err := DB.Model(u).Updates(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
