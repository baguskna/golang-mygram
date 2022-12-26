package domain

import (
	"errors"
	"golang-mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserUpdate struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex;size:255" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex;size:255" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null;" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null;" json:"age" valid:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errInit := govalidator.ValidateStruct(u)

	if errInit != nil {
		err = errInit
		return
	}

	if u.Age <= 8 {
		err = errors.New("minimum age is 8 years old")
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
