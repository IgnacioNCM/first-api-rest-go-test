package controllers

import (
	"github.com/IgnacioNCM/go-first-api/initializers"
	"github.com/IgnacioNCM/go-first-api/models"
	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {

	var body struct {
		Nombre_cliente   string
		Apellido_cliente string
		Email_cliente    string
	}

	c.Bind(&body)

	post := models.Client{Nombre_cliente: body.Nombre_cliente, Apellido_cliente: body.Apellido_cliente, Email_cliente: body.Email_cliente, Postal_code: "1240000"}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "created",
		"post":    post,
	})
}

// obtener todos los clientes
func GetClients(c *gin.Context) {
	//get data
	// Get all records
	var result []models.Client

	//fmt.Println("resultado: ", result)
	//result := models.Find(&)
	initializers.DB.Find(&result)

	count := len(result)

	// SELECT * FROM users;

	// result.RowsAffected // returns found records count, equals `len(users)`

	c.JSON(200, gin.H{
		"result": result,
		"count":  count,
	})
}

func GetClientById(c *gin.Context) {
	id := c.Param("id")

	var result models.Client
	initializers.DB.First(&result, id)

	c.JSON(200, gin.H{
		"result": result,
	})
}

func UpdateClients(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Nombre_cliente   string
		Apellido_cliente string
		Email_cliente    string
	}

	c.Bind(&body)

	var result models.Client
	initializers.DB.First(&result, id)
	//actualizar
	initializers.DB.Model(&result).Updates(models.Client{
		Nombre_cliente:   body.Nombre_cliente,
		Apellido_cliente: body.Apellido_cliente,
		Email_cliente:    body.Email_cliente,
	}).Where(result.Id_cliente, id)

	c.JSON(200, gin.H{
		"result": result,
	})
}
