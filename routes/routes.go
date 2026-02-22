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
	stocks := r.Group("/stocks")
	{
		stocks.GET("", handlers.GetStocks)
		stocks.GET("/:id", handlers.GetStock)
		stocks.POST("", handlers.CreateStock)
		stocks.PUT("/:id", handlers.UpdateStock)
		stocks.DELETE("/:id", handlers.DeleteStock)
	}
}
