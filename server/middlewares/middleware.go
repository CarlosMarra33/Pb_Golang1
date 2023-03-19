package middlewares

import (
	"application/database"
	"application/models"
	"application/services"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}

func ValidateAlunoRole(c *gin.Context) bool {
	Token := c.GetHeader("Authorization")
	var email string
	var err error
	db := database.GetDatabase()
	email, err = services.NewJWTService().ExtractEmailFromToken(Token)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return false
	}
	var aluno models.Aluno
	dbErr := db.Where("Email = ?", email).First(&aluno).Error
	if dbErr != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return false
	}

	if aluno.Email == email {
		return true
	}
	return false
}

func ValidateProfessorRole(c *gin.Context) bool {
	Token := c.GetHeader("Authorization")
	var email string
	var err error
	db := database.GetDatabase()
	email, err = services.NewJWTService().ExtractEmailFromToken(Token)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return false
	}
	var professor models.Professor
	dbErr := db.Where("Email = ?", email).First(&professor).Error
	if dbErr != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return false
	}

	if professor.Email == email {
		return true
	}
	return false
}
