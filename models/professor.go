package models

type Professor struct {
	ProfessorId uint `gorm:"primaryKey;autoIncrement"`
	// Name        string `json:"name" gorm:"not null"`
	// Email       string `json:"email" gorm:"uniqueIndex;not null"`
	// Password    string `json:"password" gorm:"not null"`
	Pessoa `gorm:"embedded"`
}
