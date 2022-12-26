package domain

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null;" json:"name" valid:"required"`
	SocialMediaURL string `gorm:"not null;" json:"social_media_url" valid:"required"`
	UserID         uint   `json:"user_id"`
	User           *User  `gorm:"foreignKey:UserID" json:"user"`
}

type SocialMediaUpdate struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

func (socilaMedia *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(socilaMedia)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
