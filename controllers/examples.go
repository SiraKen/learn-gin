package controllers

import (
	"fmt"
	"lgin/database"
	m "lgin/models"

	"github.com/gin-gonic/gin"
)

func GetExample(c *gin.Context) {

	db := database.Connect()

	defer db.Close()

	err := db.Ping()

	if err != nil {
		fmt.Println("DB Connection Unsuccessful ->", err.Error())
		return
	} else {
		fmt.Println("DB Connection Successful")
	}

	rows, err := db.Query("SELECT * FROM example WHERE id = ?", c.Param("id"))
	if err != nil {
		panic(err.Error())
	}

	exampleArgs := make([]m.Example, 0)
	for rows.Next() {
		var example m.Example
		err = rows.Scan(&example.ID, &example.Title)
		if err != nil {
			panic(err.Error())
		}
		exampleArgs = append(exampleArgs, example)
	}

	c.JSON(200, exampleArgs)
}
