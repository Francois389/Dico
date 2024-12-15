package mot

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode/utf8"
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

	if utf8.RuneCountInString(firstLetter) != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": InvalidFirstLetter})
	} else {
		mot, err := GetMotFirstLetter(firstLetter)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No word that starts with a " + firstLetter})
			return
		}

		c.JSON(http.StatusOK, mot)
	}
}