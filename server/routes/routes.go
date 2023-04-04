package routes

import (
	// "os/user"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// "application/controllers"
	"application/controllers"
	"application/database"
	"application/repositories"
	"application/server/middlewares"
	"application/services"
)



func ConfigRoutes(router *gin.Engine) *gin.Engine {
	db := database.GetDatabase()
	alunoController := controllers.NewAlunoContoller(*services.NewAlunoService(*repositories.NewAlunorepository(db)))

	router.GET("/metrics", prometheusHandler())
	main := router.Group("/api")
	{
		user := main.Group("user")
		{

			user.POST("/createAluno", alunoController.CreateAluno)
			user.POST("/createProfessor", controllers.CreateProfessor)

			user.GET("/loginAluno", alunoController.LoginAluno)
			user.GET("/loginProfessor", controllers.LoginProfessor)
		}
		aluno := main.Group("aluno", middlewares.Auth())
		{
			aluno.POST("/marcarPresenca", alunoController.MarcarPresença)
			aluno.GET("/get/aulas/:aula_id/:aluno_id", alunoController.GetPresencaAula)
		}

		professor := main.Group("professor", middlewares.Auth())
		{
			// professor.POST("/create", controllers.CreateProfessor)
			professor.POST("/create/aula", controllers.CreateAula)
			professor.PUT("/atualizar", controllers.AtualizarPresença)
		}
	}
	return router
}


func prometheusHandler() gin.HandlerFunc {
    // Crie um http.Handler a partir da função Handler() do pacote promhttp.
    promHandler := promhttp.Handler()

    // Retorne um handler do tipo gin.HandlerFunc que chama o http.Handler criado acima.
    return func(c *gin.Context) {
        promHandler.ServeHTTP(c.Writer, c.Request)
    }
}
