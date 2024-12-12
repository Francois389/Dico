package mot

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRoutes(c *gin.Engine) {
	c.GET("/mots/:firstLetter", GetMotsFirsLetter)
	c.GET("/mot/:firstLetter", GetMotFirsLetter)
}

const InvalidFirstLetter = "Invalid first letter. Expected one character."


func GetMotsFirsLetter( c *gin.Context) {
	firstLetter := c.Param("firstLetter")
	if len(firstLetter) != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidFirstLetter})
	}

	mots := GetMotsFirstLetter(firstLetter)

	c.JSON(http.StatusOK, mots)
}

func GetMotFirsLetter( c *gin.Context) {
	firstLetter := c.Param("firstLetter")

	if len(firstLetter) != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidFirstLetter})
	}

	mot := GetMotFirstLetter(firstLetter)

	c.JSON(http.StatusOK, mot)
}