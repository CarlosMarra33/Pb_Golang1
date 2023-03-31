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

type Alunorepository struct{}

// checarEmail implements ProfessorRepository
func (a *Alunorepository) ChecarEmailAluno(email string) bool {
	db := database.GetDatabase()
	var aluno models.Aluno
	fmt.Println(email)
	dberr := db.Where("email = ?", email).First(&aluno).Error
	if dberr != nil {
		return true
	}

	if aluno.Email != email {
		return true
	}
	return false
}

// salvar implements ProfessorRepository
func (a *Alunorepository) SalvarAluno(aluno models.Aluno) {
	db := database.GetDatabase()
	var save = aluno
	db.Create(&save)
}

func (a *Alunorepository) LoginAluno(login dtos.Login) (bool, models.Aluno) {
	db := database.GetDatabase()
	var aluno models.Aluno

	dberr := db.Where("email =? AND senha =?", login.Email, login.Password).First(&aluno).Error
	if dberr != nil {
		return false, aluno
	}
	// se for diferente da false
	if aluno.Email != login.Email || aluno.Password != login.Password {
		return false, aluno
	}
	return true, aluno
}
