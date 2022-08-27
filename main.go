package main

import (
	"fmt"

	"github.com/jose077/api-go-gin/database"
	"github.com/jose077/api-go-gin/routes"
)

func main() {

	database.ConectaComBancoDeDados()

	// sobre servidor e rotas da aplicação
	routes.HandleRequests()

	x := 10
	res := abc(&x)
	fmt.Println(res)

	z := &x

	*z = 100

	fmt.Println(*z, x)

}

func abc(a *int) int {
	return *a
}
