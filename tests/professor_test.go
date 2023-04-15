package services_test

import (
	"application/models"
	"application/repositories"
	"application/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProfessor(t *testing.T) {
	// Criar o repositório e o serviço
	repo := repositories.ProfessorRepository{}
	service := services.NewProfessorService(repo)

	// Criar um professor
	professor := &models.Professor{
		ProfessorId: 1,
		Pessoa: models.Pessoa{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
	}

	// Testar a criação de um professor
	result := service.CreateProfessor(professor)
	assert.Equal(t, "ok", result)

	// Testar a criação de um professor com o mesmo e-mail
	result = service.CreateProfessor(professor)
	assert.Equal(t, "user ja exieste", result)
}


