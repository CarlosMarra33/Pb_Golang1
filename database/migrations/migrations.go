package migrations

import (
	"application/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Aluno{})
	db.AutoMigrate(models.Professor{})
}
