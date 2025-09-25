package routes

import (
	"github.com/MarceloStos/Go-Alura/controllers"
	"github.com/gin-gonic/gin"
)

func HandlerRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/curso/:curso", controllers.ExibeAlunoPorCurso)
	r.GET("/", controllers.Index)
	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run()
}
