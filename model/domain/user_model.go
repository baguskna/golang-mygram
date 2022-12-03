package domain

type User struct {
	GormModel
	Username string `gorm:"not null;" json:"username"`
	Email    string `gorm:"not null;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
	Age      int    `gorm:"not null;" json:"age"`
}
