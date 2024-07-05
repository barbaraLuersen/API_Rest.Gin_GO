package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriarAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "123.456.789-10", RG: "12.345.678-9"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletarAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

// Toda função de teste deve iniciar com a palavra -> "Test"
// Toda função de teste deve receber como parâmetro -> (t * testing.T)
func TestVerificaStatusCode(t *testing.T) {
	r := SetupRotasTeste()                                               // Cria uma rota do Gin
	r.GET("/:nome", controllers.Saudacao)                                // Define um endpoint GET na rota com um handler específico
	req, _ := http.NewRequest("GET", "/Barbara", nil)                    // Cria um novo request HTTP GET para o endpoint "/nome"
	resposta := httptest.NewRecorder()                                   // Cria um gravador de resposta HTTP para capturar a resposta do servidor
	r.ServeHTTP(resposta, req)                                           // Serve o request HTTP criado usando o gravador de resposta criado anteriormente
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais") // Verifica se o código de status HTTP é 200 (OK)
	mockDaResposta := `{"API diz":"E ai Barbara, tudo beleza?"}`         // Define um mock da resposta esperada
	respostaBody, _ := ioutil.ReadAll(resposta.Body)                     // Lê o corpo da resposta HTTP capturada
	assert.Equal(t, mockDaResposta, string(respostaBody))                // Verifica se o corpo da resposta é igual ao mock esperado
}

func TestListarTodosAlunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscarPorCpf(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/123.456.789-10", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "CPF não encontrado")
}

func TestBuscarPorID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), alunoMock)
	assert.Equal(t, "Aluno Teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "123.456.789-10", alunoMock.CPF, "Os CPFs devem ser iguais")
	assert.Equal(t, "12.345.678-9", alunoMock.RG, "Os RGs devem ser iguais")
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletarAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	r := SetupRotasTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditarAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriarAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotasTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Aluno Testi", CPF: "123.456.789-22", RG: "12.345.678-0"}
	valorJson, _ := json.Marshal(aluno)
	pathDeEdicao := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathDeEdicao, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockEditado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockEditado)
	assert.Equal(t, "Aluno Testi", alunoMockEditado.Nome)
	assert.Equal(t, "123.456.789-22", alunoMockEditado.CPF)
	assert.Equal(t, "12.345.678-0", alunoMockEditado.RG)
}
