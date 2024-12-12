package main

import (
	"API-Mot/db"
	"API-Mot/mot"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Sarting...")
	db.Init("mongodb://localhost:27017", "api-mot", "mots")

	r := gin.Default()

	mot.SetUpRoutes(r)

	err := r.Run()
	if err != nil {
		fmt.Println("Error while starting the server")
		panic(err)
	}

	defer func() {
		db.Close()
	}()
}

