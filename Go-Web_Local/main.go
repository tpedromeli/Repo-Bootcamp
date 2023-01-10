package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Saludo struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {

	router := gin.Default()

	router.POST("/saludo", func(ctx *gin.Context) {
		var saludo Saludo

		ctx.BindJSON(&saludo)
		response := fmt.Sprintf("Hola %v %v ", saludo.Nombre, saludo.Apellido)

		ctx.String(200, response)
		//ctx.JSON(200, gin.H{"message": response})

	})
	router.Run()
}
