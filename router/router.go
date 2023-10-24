package router

import (
	"apiassignment/handler"

	"github.com/gin-gonic/gin"
)

func RouterStart() {
	r := gin.Default()
	r.POST("/pets", handler.CreatePet)
	r.GET("/pets", handler.GetPets)
	r.GET("/pets/:id", handler.GetPet)
	r.PUT("/pets/:id", handler.UpdatePet)
	r.DELETE("/pets/:id", handler.DeletePet)
	r.Run(":8080")
}
