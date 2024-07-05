package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/guilhermeonrails/api-go-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()

	// Configuração para carregar templates e assets
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	// Rotas para páginas específicas
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.ExibePagina404)

	// Rotas CRUD para alunos
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	// Rota para saudação
	r.GET("/:nome", controllers.Saudacao)

	// Configuração do Swagger
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Execução do servidor
	r.Run()
}
