package main

import (
	"github.com/IgnacioNCM/go-first-api/controllers"
	"github.com/IgnacioNCM/go-first-api/initializers"
	"github.com/gin-gonic/gin"
)

// acá se inicializan las variables de entorno y hace la conexion a la bd
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/clients", controllers.GetClients)
	r.GET("/clients/:id", controllers.GetClientById)
	r.POST("/clients", controllers.CreateClient)
	r.PUT("/clients/:id", controllers.UpdateClients)
	r.Run()
}
