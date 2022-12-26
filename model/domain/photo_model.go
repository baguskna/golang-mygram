package domain

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null;" json:"title" valid:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `gorm:"not null;" json:"photo_url" valid:"required"`
	UserID   uint   `json:"user_id"`
	User     *User  `json:"user" gorm:"ForeignKey:UserID"`
}

type PhotoUpdate struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
