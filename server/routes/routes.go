package routes

import (
	// "os/user"

	"github.com/gin-gonic/gin"
	// "application/controllers"
	"application/controllers"
	"application/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api")
	{
		user := main.Group("user")
		{
			user.POST("/createAluno", controllers.CreateAluno)
			user.POST("/createProfessor", controllers.CreateProfessor)

			user.GET("/loginAluno", controllers.LoginAluno)
			user.GET("/loginProfessor", controllers.LoginProfessor)
		}
		aluno := main.Group("aluno", middlewares.Auth())
		{
			aluno.POST("/marcarPresenca", controllers.MarcarPresença)
			aluno.GET("/get/aulas/:aula_id/:aluno_id", controllers.GetPresencaAula)
		}

		professor := main.Group("/professor", middlewares.Auth())
		{
			professor.POST("/create", controllers.CreateProfessor)
			professor.POST("/create/aula", controllers.CreateAula)
			
		}
	}
	return router
}
