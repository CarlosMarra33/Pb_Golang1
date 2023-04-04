package repositories

import (
	"application/controllers/dtos"
	"application/models"
	"fmt"

	"gorm.io/gorm"
)

// type IProfessorRepository interface {
// 	Salvar(professor models.Professor)
// 	ChecarEmail(email string) bool
// 	ChecarEmailSenha(dtos.Login) (bool, models.Professor)
// }

type Alunorepository struct {
	db *gorm.DB
}

func NewAlunorepository(db *gorm.DB) *Alunorepository {
	return &Alunorepository{db: db}
}

// checarEmail implements ProfessorRepository
func (a *Alunorepository) ChecarEmailAluno(email string) bool {
	// db := database.GetDatabase()
	var aluno models.Aluno
	fmt.Println(email)
	dberr := a.db.Where("email = ?", email).First(&aluno)
	fmt.Println(dberr)
	// if dberr != nil{

	// 	return true
	// }

	if aluno.Email == email {

		fmt.Println()
		return true
	}
	return false
}

// salvar implements ProfessorRepository
func (a *Alunorepository) SalvarAluno(aluno models.Aluno) {
	// db := database.GetDatabase()
	var save = aluno
	a.db.Create(&save)
}

func (a *Alunorepository) LoginAluno(login dtos.Login) (bool, models.Aluno) {
	// db := database.GetDatabase()
	var aluno models.Aluno

	dberr := a.db.Where("email =? AND senha =?", login.Email, login.Password).First(&aluno).Error
	if dberr != nil {
		return false, aluno
	}
	// se for diferente da false
	if aluno.Email != login.Email || aluno.Password != login.Password {
		return false, aluno
	}
	return true, aluno
}
