package services

import (
	"application/controllers/dtos"
	"application/models"
	"application/repositories"
)

type AlunoService struct {
	repo repositories.Alunorepository
}

// CreateProfessor implements ProfessorService
func (a *AlunoService) CreateAluno(aluno *models.Aluno) string {
	var email = aluno.Email
	check := a.repo.ChecarEmailAluno(email)
	if check {
		return "user ja exieste"
	}
	a.repo.SalvarAluno(*aluno)
	return "ok"
}

// LoginProfessor implements ProfessorService
func (a *AlunoService) LoginAluno(login *dtos.Login) (string, error) {
	chek, aluno := a.repo.LoginAluno(*login)

	if !chek {
		return "usuário ou senha incorrect", nil
	}
	token, err := NewJWTService().GenerateToken(aluno.AlunoId)
	if err != nil {

		return "", err
	}

	return token, nil
}

func NewAlunoService(repo repositories.Alunorepository) *AlunoService {
	return &AlunoService{
		repo: repo,
	}
}
