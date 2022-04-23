package main

import (
	"example/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Example struct {
	ID	int `json:"id"`
	Title	string `json:"title"`
}

func main() {

	db := database.Connect()

	defer db.Close()

	err := db.Ping()

	if (err != nil) {
		fmt.Println("DB Connection Unsuccessful ->", err.Error())
		return
	} else {
		fmt.Println("DB Connection Successful")
	}

	r := gin.Default()
	r.GET("/example/:id", func(c *gin.Context) {

		rows, err := db.Query("SELECT * FROM example WHERE id = ?", c.Param("id"))
		if err != nil {
			panic(err.Error())
		}

		exampleArgs := make([]Example, 0)
		for rows.Next() {
			var example Example
			err = rows.Scan(&example.ID, &example.Title)
			if (err != nil) {
				panic(err.Error())
			}
			exampleArgs = append(exampleArgs, example)
		}

		c.JSON(200, exampleArgs)
	})
	// listem and serve on 0:8080
	r.Run()
}

