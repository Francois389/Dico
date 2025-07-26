package main

import (
	"api/db"
	"api/word"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting...")

	url := os.Getenv("MONGO_URI")
	if url == "" {
		url = "mongodb://localhost:27017/dico-db"
	}
	databaseName := os.Getenv("MONGO_DB")
	if databaseName == "" {
		databaseName = "dico-db"
	}
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
