package domain

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null;" json:"name"`
	SocialMediaURL string `gorm:"not null;" json:"social_media_url"`
	UserID         User   `gorm:"foreignKey:UserID" json:"user"`
}
