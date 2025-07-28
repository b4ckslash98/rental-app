package entity

type User struct {
	ID           int    `gorm:"primaryKey"`
	Email        string `gorm:"uniqueIndex;size:100"`
	PasswordHash string `gorm:"size:255"`
	Role         string `gorm:"type:enum('admin','customer')"`
}
