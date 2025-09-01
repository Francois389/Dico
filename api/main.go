package main

import (
	"api/db"
	"api/word"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	verbose := flag.Bool("v", false, "Display all errors")
	flag.Parse()

	fmt.Println("Starting...")

	url := os.Getenv("MONGO_URI")
	if url == "" {
		url = "mongodb://localhost:27017/dico-db"
	}
	databaseName := os.Getenv("MONGO_DB")
	if databaseName == "" {
		databaseName = "dico-db"
	}
	fmt.Println("Connecting to " + url + "/" + databaseName)
	collectionName := "mots"

	fmt.Println("Connecting to the database...")
	err := db.Init(url, databaseName, collectionName)
	if err != nil {
		fmt.Println("Error while connecting to the database")
		if *verbose {
			fmt.Println(err)
		} else {
			fmt.Println("Use -v flag for detailed error information")
		}
		return
	}
	fmt.Println("Connected to the database")

	r := gin.Default()

	word.SetUpRoutes(r)
	var port = ":4242"

	err = r.Run(port)
	if err != nil {
		fmt.Println("Error while starting the server")
		if *verbose {
			panic(err)
		}
		fmt.Println("Server failed to start. Use -v flag for detailed error information")
	}

	defer func() {
		db.Close()
	}()
}
