package main

import (
	"Dico/db"
	"Dico/mot"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Sarting...")
	db.Init("mongodb://localhost:27027", "dico-db", "mots")

	r := gin.Default()

	mot.SetUpRoutes(r)
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
