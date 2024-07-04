package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/stretchr/testify/assert"
)

func SetupRotasTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

// Toda função de teste deve iniciar com a palavra -> "Test"
// Toda função de teste deve receber como parâmetro -> (t * testing.T)
func TestVerificaStatusCode(t *testing.T) {
	r := SetupRotasTeste()                                               // Cria uma rota do Gin
	r.GET("/:nome", controllers.Saudacao)                                // Define um endpoint GET na rota com um handler específico
	req, _ := http.NewRequest("GET", "/nome", nil)                       // Cria um novo request HTTP GET para o endpoint "/nome"
	resposta := httptest.NewRecorder()                                   // Cria um gravador de resposta HTTP para capturar a resposta do servidor
	r.ServeHTTP(resposta, req)                                           // Serve o request HTTP criado usando o gravador de resposta criado anteriormente
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais") // Verifica se o código de status HTTP é 200 (OK)
	mockDaResposta := `{"API diz":"E ai Barbara, tudo beleza?"}`         // Define um mock da resposta esperada
	respostaBody, _ := ioutil.ReadAll(resposta.Body)                     // Lê o corpo da resposta HTTP capturada
	assert.Equal(t, mockDaResposta, string(respostaBody))                // Verifica se o corpo da resposta é igual ao mock esperado
}
