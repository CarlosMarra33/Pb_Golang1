package repositories

import (
	"application/controllers/dtos"
	"application/database"
	"application/models"

	"gorm.io/gorm"
)

// type IProfessorRepository interface {
// 	Salvar(professor models.Professor)
// 	ChecarEmail(email string) bool
// 	ChecarEmailSenha(dtos.Login) (bool, models.Professor)
// }

type ProfessorRepository struct {
	db *gorm.DB
}

func NewProfessorRepository() *ProfessorRepository {
	return &ProfessorRepository{
		db: database.GetDatabase()}

}

// checarEmail implements ProfessorRepository
func (pr *ProfessorRepository) ChecarEmail(email string) bool {
	var professor models.Professor
	// fmt.Println(email)
	pr.db.Where("email = ?", email).First(&professor)
	// if dberr != nil {
	// 	fmt.Println("erro da chamada", dberr, "professor  ", professor)
	// 	return true
	// }
	// fmt.Println(professor)
	if professor.Email == email {
		return true
	}
	return false
}

// salvar implements ProfessorRepository
func (pr *ProfessorRepository) Salvar(professor models.Professor) {
	var prof = professor
	pr.db.Create(&prof)
}

func (p *ProfessorRepository) ChecarEmailSenha(login dtos.Login) (bool, models.Professor) {
	var professor models.Professor

	dberr := p.db.Where("email =? AND password =?", login.Email, login.Password).First(&professor).Error
	if dberr != nil {
		return false, professor
	}
	// se for diferente da false
	if professor.Email != login.Email || professor.Password != login.Password {
		return false, professor
	}
	return true, professor
}
