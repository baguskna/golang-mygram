package domain

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserID  uint   `json:"user_id"`
	User    *User  `gorm:"foreignKey:UserID" json:"user"`
	PhotoID uint   `json:"photo_id"`
	Photo   *Photo `gorm:"foreignKey:PhotoID" json:"photo"`
	Message string `gorm:"not null;" json:"message" valid:"required"`
}

type CommentUpdate struct {
	Message string `json:"message"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
