package controllers

import (
	"application/controllers/dtos"
	"application/database"
	"application/models"
	"application/services"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)



func LoginProfessor(ctx *gin.Context) {
	db := database.GetDatabase()
	var login dtos.Login
	err := ctx.ShouldBindJSON(&login)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var prof models.Professor
	dberr := db.Where("email = ?", login.Email).First(&prof).Error
	if dberr != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot find user",
		})
		return
	}

	if login.Password != prof.Password {
		ctx.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(prof.ProfessorId)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})

}

func CreateProfessor(c *gin.Context) {
	db := database.GetDatabase()
	var professor models.Professor
	err := c.ShouldBindJSON(&professor)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Problema ao passar para JSON" + err.Error(),
		})
		return
	}

	err = db.Create(&professor).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Problema ao criar aluno" + err.Error(),
		})
		return
	}

	c.Status(204)
}

func CreateAula(ctx *gin.Context) {
	
	url := "http://localhost:5001/api/aula/criar"

	var aula dtos.Aula_dto


	err := ctx.ShouldBindJSON(&aula)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	jsonData, err := json.Marshal(aula)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	buffer := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "criação da requisição falhou",
		})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "requisição falhou"+ err.Error(), 
		})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		ctx.JSON(400, gin.H{
			"error": "NÃO FOI ",
		})
		return
	}

}

// func vincularAluno(c *gin.Context){
// 	url := "localhost:5001/api/criar/aula"

// }
