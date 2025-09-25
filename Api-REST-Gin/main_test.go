package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/MarceloStos/Go-Alura/controllers"
	"github.com/MarceloStos/Go-Alura/database"
	"github.com/MarceloStos/Go-Alura/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", Idade: 99, Curso: "Curso Aluno Teste", Matricula: "45896584"}
	database.DB.Create(&aluno)
	fmt.Println("Aluno mock criado")
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, "matricula", "45896584")
	fmt.Println("Aluno mock deletado")
}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Marcelo", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz:":"E ai Marcelo, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody), "Deveriam ser iguais")
}

func TestListarTodosOsAlunosHandler(t *testing.T) {
	database.ConectacomBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code, "Deveriam ser iguais")
}

func TestBuscaAlunoPorCursoHandler(t *testing.T) {
	database.ConectacomBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/curso/:curso", controllers.ExibeAlunoPorCurso)
	req, _ := http.NewRequest("GET", "/alunos/curso/Curso Aluno Teste", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code, "Deveriam ser iguais")
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectacomBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var alunoMock models.Aluno
	json.Unmarshal(w.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Aluno Teste", alunoMock.Nome, "Deveriam ser iguais")
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectacomBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code, "Deveriam ser iguais")
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectacomBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	aluno := models.Aluno{Nome: "Aluno Teste Editado", Idade: 99, Curso: "Curso Aluno Teste", Matricula: "45896584"}
	alunoJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(alunoJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var alunoEditado models.Aluno
	json.Unmarshal(w.Body.Bytes(), &alunoEditado)
	assert.Equal(t, "Aluno Teste Editado", alunoEditado.Nome, "Deveriam ser iguais")
}
