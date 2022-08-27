package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jose077/api-go-gin/controllers"
)

func HandleRequests() {
	log.Println("subiu!")
	r := gin.Default()

	r.Run(":5000") //-> caso não definido o servidor será rodado na porta 8080

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCPF)

	r.Run()
}
