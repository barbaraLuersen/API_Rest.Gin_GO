package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
)

// ExibeTodosAlunos godoc
// @Summary Exibe todos os alunos
// @Description Retorna uma lista de todos os alunos cadastrados
// @Tags alunos
// @Produce json
// @Success 200 {array} models.Aluno
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// Saudacao godoc
// @Summary Saúda o usuário
// @Description Retorna uma saudação personalizada para o usuário
// @Tags saudacoes
// @Produce json
// @Param nome path string true "Nome do usuário"
// @Success 200 {object} map[string]string
// @Router /{nome} [get]
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	// Verifica se a rota é "/index" e evita tratá-la como um nome de aluno
	if nome == "index" {
		ExibePaginaIndex(c)
		return
	}

	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo beleza?",
	})
}

// CriaNovoAluno godoc
// @Summary Cria um novo aluno
// @Description Adiciona um novo aluno ao banco de dados
// @Tags alunos
// @Accept json
// @Produce json
// @Param aluno body models.Aluno true "Novo aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} map[string]string
// @Router /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	//Validação antes do envio do dado ao banco de dados
	//Se houver erro na função ValidarDados
	if err := models.ValidaDados(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorID godoc
// @Summary Busca aluno por ID
// @Description Retorna um aluno pelo ID
// @Tags alunos
// @Produce json
// @Param id path string true "ID do aluno"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} map[string]string
// @Router /alunos/{id} [get]
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno godoc
// @Summary Deleta aluno por ID
// @Description Remove um aluno pelo ID
// @Tags alunos
// @Produce json
// @Param id path string true "ID do aluno"
// @Success 200 {object} map[string]string
// @Router /alunos/{id} [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditaAluno godoc
// @Summary Edita aluno por ID
// @Description Atualiza os dados de um aluno pelo ID
// @Tags alunos
// @Accept json
// @Produce json
// @Param id path string true "ID do aluno"
// @Param aluno body models.Aluno true "Dados do aluno atualizados"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} map[string]string
// @Router /alunos/{id} [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	//Busca um aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	//Valida se os dados passados são válidos
	if err := models.ValidaDados(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF godoc
// @Summary Busca aluno por CPF
// @Description Retorna um aluno pelo CPF
// @Tags alunos
// @Produce json
// @Param cpf path string true "CPF do aluno"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} map[string]string
// @Router /alunos/cpf/{cpf} [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// Funções para exibição de tela HTML com Gin

// Exibe a página inicial
func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

// Exibe a página 404
func ExibePagina404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
