// backend/main.go
package main

import (
	"velomanager/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/velodb?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("échec de la connexion à la base MySQL")
	}

	db.AutoMigrate(&models.Bike{})

	r := gin.Default()

	// Routes API
	r.GET("/bikes", func(c *gin.Context) {
		var bikes []models.Bike
		db.Find(&bikes)
		c.JSON(200, bikes)
	})

	r.POST("/bikes", func(c *gin.Context) {
		var bike models.Bike
		if err := c.ShouldBindJSON(&bike); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&bike)
		c.JSON(201, bike)
	})

	r.Run(":8080")
}
