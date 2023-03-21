package controllers

import (
	"application/controllers/dtos"
	"application/database"
	"application/models"
	"application/services"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoginAluno(ctx *gin.Context) {
	db := database.GetDatabase()
	var login dtos.Login
	err := ctx.ShouldBindJSON(&login)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var aluno models.Aluno
	dberr := db.Where("email = ?", login.Email).First(&aluno).Error
	if dberr != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot find user",
		})
		return
	}

	if login.Password != aluno.Password {
		ctx.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(aluno.AlunoId)
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

func CreateAluno(ctx *gin.Context) {
	db := database.GetDatabase()
	var aluno models.Aluno
	erro := ctx.ShouldBindJSON(&aluno)

	if erro != nil {
		ctx.JSON(400, gin.H{
			"error": "Problema ao passar para JSON" + erro.Error(),
		})
		return
	}

	err := db.Create(&aluno).Error

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Problema ao criar aluno" + erro.Error(),
		})
		return
	}

	ctx.Status(204)
}

func MarcarPresença(ctx *gin.Context) {
	url := "http://localhost:5001/api/aula/presente"
	var presencaAluno dtos.PresencaAluno

	// validate := middlewares.ValidateAlunoRole(ctx)
	// if !validate {
	// 	ctx.JSON(401, gin.H{
	// 		"error": "",
	// 	})
	// 	return
	// }

	err := ctx.ShouldBindJSON(&presencaAluno)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	fmt.Println(presencaAluno)

	jsonData, err := json.Marshal(presencaAluno)
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
			"error": "requisição falhou",
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

func GetPresencaAula(c *gin.Context) {
	url := "http://localhost:5001/api/presenca/getPresenca"
	var response dtos.GetPresencaAula

	idAula, _ := strconv.ParseInt(c.Param("aula_id"), 10, 64)
	idAluno, _ := strconv.ParseInt(c.Param("aluno_id"), 10, 64)

	fmt.Println(idAluno)
	fmt.Println(idAula)
	resp, err := http.Get(fmt.Sprintf(url+"/%d/%d", int(idAluno), int(idAula)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	fmt.Println(resp.Body)
	var r interface{}
	err = json.Unmarshal(body, &r)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possível mapear o JSON para a struct Response"})
		return
	}
	fmt.Println(response)

	c.JSON(200, r)

}
