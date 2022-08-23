package model

import "gorm.io/gorm"

// swagger:model
type Users struct {
	// required: true
	// min: 1
	UserID uint `gorm:"primarykey;autoIncrement;not null" json:"user_id"`
	// the fullname of the user
	// required: true
	// min length: 15
	Fullname *string `gorm:"not null" json:"fullname"`
	// the username of the user
	// required: true
	// min length: 15
	Username *string `gorm:"not null" json:"username"`
	// the password of the user
	// required: true
	// min length: 15
	// max length: 20
	Password *string `gorm:"not null" json:"password"`
}

func MigrateStruct(db *gorm.DB) error {
	err := db.AutoMigrate(&Users{})
	return err
}
