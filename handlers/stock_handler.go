package handlers

import (
	"example-crud/config"
	"example-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStocks(c *gin.Context) {
	var stocks []models.Stock
	config.DB.Find(&stocks)
	c.JSON(http.StatusOK, stocks)
}

func GetStock(c *gin.Context) {
	var stock models.Stock
	if err := config.DB.First(&stock, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}
	c.JSON(http.StatusOK, stock)
}

func CreateStock(c *gin.Context) {
	var stock models.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&stock)
	c.JSON(http.StatusCreated, stock)
}

func UpdateStock(c *gin.Context) {
	var stock models.Stock
	if err := config.DB.First(&stock, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}
	c.ShouldBindJSON(&stock)
	config.DB.Save(&stock)
	c.JSON(http.StatusOK, &stock)
}

func DeleteStock(c *gin.Context) {
	config.DB.Delete(&models.Stock{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Stock deleted"})
}
