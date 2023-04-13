package controllers

import (
	"application/controllers/dtos"
	"application/models"
	"application/services"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/gin-gonic/gin"
)

type AlunoController struct {
	alunoService services.AlunoService
}

func NewAlunoContoller(service services.AlunoService) *AlunoController {
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(requestDurationHistogram)
	return &AlunoController{alunoService: service}
}

var (
	requestDurationHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "duracao_do_create",
		Help:    "Duration of HTTP requests in seconds.",
		Buckets: []float64{0.1, 0.5, 1, 2.5, 5, 10},
	})
)

var (
	RequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "todas_requests",
		Help: "Total number of HTTP requests.",
	})
)

 var resposnseError = "cannot bind JSON: "

func (ac *AlunoController) LoginAluno(ctx *gin.Context) {
	// service := services.NewAlunoService(repositories.Alunorepository{})
	var login dtos.Login
	err := ctx.ShouldBindJSON(&login)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": resposnseError + err.Error(),
		})
		return
	}
	token, err := ac.alunoService.LoginAluno(&login)

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

func (ac *AlunoController) CreateAluno(ctx *gin.Context) {
	// service := services.NewAlunoService(repositories.Alunorepository{})

	start := time.Now()
	var aluno models.Aluno
	err := ctx.ShouldBindJSON(&aluno)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Problema ao passar para JSON" + err.Error(),
		})
		return
	}

	

	ac.alunoService.CreateAluno(&aluno)

	ctx.Status(204)
	
	RequestsTotal.Inc()
	duration := time.Since(start).Seconds()
	requestDurationHistogram.Observe(duration)

}

func (ac *AlunoController) MarcarPresença(ctx *gin.Context) {
	url := "http://localhost:5001/api/presenca/presente"
	var presencaAluno dtos.PresencaAluno

	err := ctx.ShouldBindJSON(&presencaAluno)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": resposnseError + err.Error(),
		})
		return
	}

	fmt.Println(presencaAluno)

	jsonData, err := json.Marshal(presencaAluno)
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
			"error": "requisição falhou",
		})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusBadRequest {
		ctx.JSON(400, gin.H{
			"error": "NÃO FOI ",
		})
		return
	}

	if resp.StatusCode == http.StatusFound {
		ctx.JSON(http.StatusFound, gin.H{
			"error": "Presensa de hoje já foi marcada ",
		})
		return
	}

}

func (ac *AlunoController) GetPresencaAula(c *gin.Context) {
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
