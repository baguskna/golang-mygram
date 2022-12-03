package domain

type Comment struct {
	GormModel
	UserID  User   `gorm:"foreignKey:UserID" json:"user"`
	PhotoID Photo  `gorm:"foreignKey:PhotoID" json:"photo"`
	Message string `gorm:"not null;" json:"message"`
}
