package controllers

import (
	"github.com/MarceloStos/Go-Alura/database"
	"github.com/MarceloStos/Go-Alura/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"mensagem": "Olá " + nome + ", tudo bem?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var novoAluno models.Aluno
	if err := c.ShouldBindJSON(&novoAluno); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidaDadosDeAlunos(&novoAluno); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&novoAluno)
	c.JSON(200, novoAluno)
}

func ExibeAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"mensagem": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"mensagem": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, gin.H{
		"mensagem": "O aluno foi deletado com sucesso!",
	})

}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidaDadosDeAlunos(&aluno); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(200, aluno)
}

func ExibeAlunoPorCurso(c *gin.Context) {
	var alunos []models.Aluno
	curso := c.Params.ByName("curso")
	database.DB.Where(&models.Aluno{Curso: curso}).Find(&alunos)
	if len(alunos) == 0 {
		c.JSON(404, gin.H{
			"mensagem": "Nenhum aluno encontrado para esse curso",
		})
		return
	}
	c.JSON(200, alunos)
}

func Index(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(200, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(404, "404.html", nil)
}
