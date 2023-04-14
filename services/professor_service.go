package services

import (
	"application/controllers/dtos"
	"application/models"
	"application/repositories"
	"fmt"
)

type ProfessorService struct {
	repo repositories.ProfessorRepository
}

func NewProfessorService(repo repositories.ProfessorRepository) *ProfessorService {
	return &ProfessorService{
		repo: repo,
	}
}

func (ps *ProfessorService) CreateProfessor(professor *models.Professor) string {
	fmt.Println(professor)
	check := ps.repo.ChecarEmail(professor.Email)
	fmt.Println(check)
	if check {
		return "user ja exieste"
	}
	ps.repo.Salvar(*professor)
	return "ok"
}

func (ps *ProfessorService) LoginProfessor(login *dtos.Login) (string, error) {
	chek, professor := ps.repo.ChecarEmailSenha(*login)

	if !chek {
		return "usuário ou senha incorrect", nil
	}
	token, err := NewTokenService().GenerateToken(professor.ProfessorId)
	if err != nil {

		return "", err
	}

	return token, nil
}
