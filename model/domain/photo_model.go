package domain

type Photo struct {
	GormModel
	Title    string `gorm:"not null;" json:"title"`
	Caption  string `gorm:"not null;" json:"caption"`
	PhotoURL string `gorm:"not null;" json:"photo_url"`
	UserID   User   `gorm:"foreignKey:UserID" json:"user"`
}
