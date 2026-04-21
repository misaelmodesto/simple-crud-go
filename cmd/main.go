package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misaelmodesto/simple-crud-go/controller"
	"github.com/misaelmodesto/simple-crud-go/db"
	"github.com/misaelmodesto/simple-crud-go/repository"
	"github.com/misaelmodesto/simple-crud-go/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de Repositorio
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Camada UseCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//Camada de Controllers
	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)

	server.GET("/product/:productId", productController.GetProductById)

	server.Run(":8000")

}
