package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")                 //Rota para entrar no pacote de templates (HTML)
	r.Static("/assets", "./assets")               //Rota para entrar no pacote de assets (CSS)
	r.GET("/index", controllers.ExibePaginaIndex) //Rota para a tela
	r.NoRoute(controllers.ExibePagina404)         //Rota para tela de quando o link não encontra a rota

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
