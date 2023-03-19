package models

type Professor struct {
	ProfessorId uint   `gorm:"primaryKey;autoIncrement"`
	Nome        string `json:"name" gorm:"not null"`
	Email       string `json:"email" gorm:"uniqueIndex;not null"`
	Password    string `json:"password" gorm:"not null"`
}
