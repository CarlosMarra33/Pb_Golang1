package services

import (
	"application/controllers/dtos"
	"application/models"
	"application/repositories"
	"fmt"
)

type AlunoService struct {
	repo repositories.Alunorepository
}

func NewAlunoService(repo repositories.Alunorepository) *AlunoService {
	return &AlunoService{
		repo: repo,
	}
}

// CreateProfessor implements ProfessorService
func (as *AlunoService) CreateAluno(aluno *models.Aluno) string {
	var email = aluno.Email
	check := as.repo.ChecarEmailAluno(email)

	fmt.Println("teste do service  ",check)
	if check {
		return "user ja exieste"
	}
	as.repo.SalvarAluno(*aluno)
	return "ok"
}

// LoginProfessor implements ProfessorService
func (as *AlunoService) LoginAluno(login *dtos.Login) (string, error) {
	chek, aluno := as.repo.LoginAluno(*login)

	if !chek {
		return "usuário ou senha incorrect", nil
	}
	token, err := NewJWTService().GenerateToken(aluno.AlunoId)
	if err != nil {

		return "", err
	}

	return token, nil
}

func (as *AlunoService) VerificarAluno(idAluno uint) (string) {
	response := as.repo.VerificarAlunoId(idAluno)
	if response{
		return "ok"
	}
	return "Aluno não encontrado"

}

