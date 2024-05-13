package controllers

import (
	"errors"
	"lgin/database"
	m "lgin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetExample(c *gin.Context) {

	db := database.Connect()

	res := db.Exec("SELECT * FROM examples WHERE id = ?", c.Param("id"))
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{
				"message": "Record not found",
			})
			return
		} else {
			panic(res.Error)
		}
	}

	var exampleArgs m.Example
	res.Scan(&exampleArgs)

	c.JSON(200, exampleArgs)
}
