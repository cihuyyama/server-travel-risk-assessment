package models

import (
	"time"
	"travel-risk-assessment/helpers"
)

type User struct {
	ID           uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email" gorm:"unique"`
	Password     string    `json:"password"`
	Photos       []Photo   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role         string    `json:"role" gorm:"default:user"`
	UserSymptoms []Symptom `gorm:"many2many:user_symptoms;"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (user *User) HashPassword(password string) error {
	has, err := helpers.HashPassword(password)
	if err != nil {
		return err

	}
	user.Password = has
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	result, err := helpers.CheckPasswordHash(providedPassword, user.Password)
	if !result {
		return err
	}
	return nil
}
