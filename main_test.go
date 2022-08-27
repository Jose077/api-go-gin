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
	"github.com/jose077/api-go-gin/controllers"
	"github.com/jose077/api-go-gin/database"
	"github.com/jose077/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	// instancia roteador
	r := SetupDasRotasDeTeste()

	// instancia rota
	r.GET("/:nome", controllers.Saudacao)

	// cria modelo de requisicao -> http.NewRequest("METHOD", "ROTA", "PAYLOAD")
	req, _ := http.NewRequest("GET", "/gui", nil)

	// Cria modelo para resposta
	resposta := httptest.NewRecorder()

	// realiza requisicao
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API":"E aí gui, tudo blza?"}`

	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mockDaResposta, string(respostaBody))
}

var ID int

// Cria aluno de teste np db
func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno teste", CPF: "12345678912", RG: "123456789"}

	database.DB.Create(&aluno)

	ID = int(aluno.ID)
}

// deleta aluno teste do db
func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListandoTodosOsAlunos(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/22222222222", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	// Cria aluno db
	CriaAlunoMock()

	// ao finalizar deleta aluno do db
	defer DeletaAlunoMock()

	// Instancia roteador
	r := SetupDasRotasDeTeste()

	// declara rota
	r.GET("/alunos/:id", controllers.BuscarAlunoPorId)

	// Cria path com id
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)

	// Cria request
	req, _ := http.NewRequest("GET", pathDaBusca, nil)

	// Cria resposta
	resposta := httptest.NewRecorder()

	// Realiza requisicao
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno

	// recebe aluno da resposta e atribui dentro de alunoMock em json
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Aluno teste", alunoMock.Nome)
	assert.Equal(t, "12345678912", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	pathBusca := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", pathBusca, nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47123456789", RG: "123456700"}
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
}
