// backend/main.go
package main

import (
	"fmt"
	"velomanager/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var sql_password, sql_login string
	sql_address := "127.0.0.1:3306"
	fmt.Println("Golang x DataBase")
	// Setup connection for my db
	fmt.Scanln(&sql_login)
	fmt.Println("Enter your login")
	fmt.Scanln(&sql_password)
	fmt.Println("Enter your password")
	dsn := sql_login + ":" + sql_password + "@tcp(" + sql_address + ")/velodb?parseTime=true"
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
