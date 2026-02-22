package routes

import (
	"example-crud/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	items := r.Group("/items")
	{
		items.GET("", handlers.GetItems)
		items.GET("/:id", handlers.GetItem)
		items.POST("", handlers.CreateItem)
		items.PUT("/:id", handlers.UpdateItem)
		items.DELETE("/:id", handlers.DeleteItem)
	}
}
