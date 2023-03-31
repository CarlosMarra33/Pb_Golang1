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

// CreateProfessor implements ProfessorService
func (p *ProfessorService) CreateProfessor(professor *models.Professor) string {
	fmt.Println(professor)
	check := p.repo.ChecarEmail(professor.Email)
	fmt.Println(check)
	if check {
		return "user ja exieste"
	}
	p.repo.Salvar(*professor)
	return "ok"
}

// LoginProfessor implements ProfessorService
func (p *ProfessorService) LoginProfessor(login *dtos.Login) (string, error) {
	chek, professor := p.repo.ChecarEmailSenha(*login)

	if !chek {
		return "usuário ou senha incorrect", nil
	}
	token, err := NewJWTService().GenerateToken(professor.ProfessorId)
	if err != nil {

		return "", err
	}

	return token, nil
}


