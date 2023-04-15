package repositories

import (
	"application/controllers/dtos"
	"application/database"
	"application/models"
	"fmt"

	"gorm.io/gorm"
)

type ProfessorRepository struct {
	db *gorm.DB
}

func NewProfessorRepository() *ProfessorRepository {
	return &ProfessorRepository{
		db: database.GetDatabase()}

}

func (pr *ProfessorRepository) ChecarEmail(email string) bool {
	var professor models.Professor
	pr.db.Where("email = ?", email).First(&professor)

	if professor.Email == email {
		return true
	}
	return false
}

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
	if professor.Email != login.Email || professor.Password != login.Password {
		return false, professor
	}
	return true, professor
}

func (p *ProfessorRepository) VarificarListaAlunos(alunosid []uint) (bool, error) {

	// var aluno models.Aluno
	var alunos []models.Aluno
	p.db.Find(&alunos, "aluno_id IN (?)", alunosid)

	fmt.Println(len(alunos))

	if len(alunos) != len(alunosid) {
		return false, nil
	}
	return true, nil
}
