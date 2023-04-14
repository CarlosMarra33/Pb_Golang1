package repositories

import (
	"application/controllers/dtos"
	"application/database"
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

func NewAlunorepository() *Alunorepository {
	return &Alunorepository{db: database.GetDatabase()}
}

// checarEmail implements ProfessorRepository
func (ar *Alunorepository) ChecarEmailAluno(email string) bool {
	// db := database.GetDatabase()
	var aluno models.Aluno
	fmt.Println(email)
	ar.db.Where("email = ?", email).First(&aluno)

	if aluno.Email == email {

		fmt.Println()
		return true
	}
	return false
}

// salvar implements ProfessorRepository
func (ar *Alunorepository) SalvarAluno(aluno models.Aluno) {
	// db := database.GetDatabase()
	var save = aluno
	ar.db.Create(&save)
}

func (ar *Alunorepository) LoginAluno(login dtos.Login) (bool, models.Aluno) {
	// db := database.GetDatabase()
	var aluno models.Aluno

	dberr := ar.db.Where("email =? AND password =?", login.Email, login.Password).First(&aluno).Error
	if dberr != nil {
		return false, aluno
	}
	// se for diferente da false
	if aluno.Email != login.Email || aluno.Password != login.Password {
		return false, aluno
	}
	return true, aluno
}

func (ar *Alunorepository) VerificarAlunoId(idAluno uint) (bool){
	var aluno models.Aluno
	response := ar.db.Where(idAluno).First(&aluno)
	if response != nil {
		return true
	}
	return false
}
