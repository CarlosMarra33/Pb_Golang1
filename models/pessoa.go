package models

type Pessoa struct {
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
}
