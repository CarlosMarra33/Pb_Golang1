package models

type Aluno struct {
	AlunoId  uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	// Pessoa   `gorm:"embedded"`
}
