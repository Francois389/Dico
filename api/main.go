package main

import (
	"Dico/db"
	"Dico/word"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting...")

	url := "mongodb://localhost:27027"
	databaseName := "dico-db"
	collectionName := "mots"

	err := db.Init(url, databaseName, collectionName)
	if err != nil {
		fmt.Println("Error while connecting to the database")
		fmt.Println(err)
		return
	}

	r := gin.Default()

	word.SetUpRoutes(r)
	var port = ":4242"

	err = r.Run(port)
	if err != nil {
		fmt.Println("Error while starting the server")
		panic(err)
	}

	defer func() {
		db.Close()
	}()
}
