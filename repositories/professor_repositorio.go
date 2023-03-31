package repositories

import (
	"application/controllers/dtos"
	"application/database"
	"application/models"
	"fmt"
)

// type IProfessorRepository interface {
// 	Salvar(professor models.Professor)
// 	ChecarEmail(email string) bool
// 	ChecarEmailSenha(dtos.Login) (bool, models.Professor)
// }

type ProfessorRepository struct{}

// checarEmail implements ProfessorRepository
func (p *ProfessorRepository) ChecarEmail(email string) bool {
	db := database.GetDatabase()
	var professor models.Professor
	// fmt.Println(email)
	dberr := db.Where("email = ?", email).First(&professor).Error
	if dberr != nil {
		fmt.Println("erro da chamada",dberr, "professor  ", professor )
		return true
	}
	// fmt.Println(professor)
	if professor.Email == email {
		return true
	}
	return false
}

// salvar implements ProfessorRepository
func (p *ProfessorRepository) Salvar(professor models.Professor) {
	db := database.GetDatabase()
	var prof = professor
	db.Create(&prof)
}

func (p *ProfessorRepository) ChecarEmailSenha(login dtos.Login) (bool, models.Professor) {
	db := database.GetDatabase()
	var professor models.Professor

	dberr := db.Where("email =? AND senha =?", login.Email, login.Password).First(&professor).Error
	if dberr != nil {
		return false, professor
	}
	// se for diferente da false
	if professor.Email != login.Email || professor.Password != login.Password {
		return false, professor
	}
	return true, professor
}
