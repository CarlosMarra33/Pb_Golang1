package controllers

import (
	"application/controllers/dtos"
	"application/models"
	"application/services"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfessorController struct {
	profService services.ProfessorService
}

func NewProfessorController(service services.ProfessorService) *ProfessorController {
	return &ProfessorController{

		profService: service,
	}
}

func (pc *ProfessorController) LoginProfessor(ctx *gin.Context) {
	var login dtos.Login
	err := ctx.ShouldBindJSON(&login)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": resposnseError + err.Error(),
		})
		return
	}

	token, err := pc.profService.LoginProfessor(&login)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot find user",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})

}

func (pc *ProfessorController) CreateProfessor(c *gin.Context) {
	var professor models.Professor
	err := c.ShouldBindJSON(&professor)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Problema ao passar para JSON" + err.Error(),
		})
		return
	}
	pc.profService.CreateProfessor(&professor)

	c.Status(204)
}

func (pc *ProfessorController) CreateAula(ctx *gin.Context) {

	url := "http://localhost:5001/api/aula/criar"

	var aula dtos.Aula_dto

	err := ctx.ShouldBindJSON(&aula)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": resposnseError + err.Error(),
		})
		return
	}

	isValid := pc.profService.ValidarLista(aula.Alunos)
	if isValid != "ok" {
		ctx.JSON(400, gin.H{
			"error": isValid,
		})
		return
	}

	jsonData, err := json.Marshal(aula)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": resposnseError + err.Error(),
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
			"error": "requisição falhou" + err.Error(),
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

func (pc *ProfessorController) AtualizarPresença(ctx *gin.Context) {
	url := "http://localhost:5001/api/presenca/atualizar"

	var p dtos.PresencaAluno
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": resposnseError + err.Error(),
		})
		return
	}
	jsonData, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("PUT", url, buffer)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		ctx.Status(200)
	} else {
		ctx.Status(400)
	}
}
