package mot

import (
	"fmt"
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

	fmt.Println("firstLetter: ", firstLetter)
	fmt.Printf("firstLetter: %d\n", utf8.RuneCountInString(firstLetter))

	if utf8.RuneCountInString(firstLetter) != 1 {
		fmt.Println("Invalid first letter")
		c.JSON(http.StatusBadRequest, getErrorInvalidFirstLetter())
	} else {

		mots, err := GetMotsFirstLetter(firstLetter)

		if err != nil {
			c.JSON(http.StatusNotFound, getErrorNoWordStartWith(firstLetter))
			return
		}

		c.JSON(http.StatusOK, mots)
	}
}

func GetMotFirsLetter( c *gin.Context) {
	firstLetter := c.Param("firstLetter")

	if utf8.RuneCountInString(firstLetter) != 1 {
		c.JSON(http.StatusBadRequest, getErrorInvalidFirstLetter())
	} else {
		mot, err := GetMotFirstLetter(firstLetter)
		if err != nil {
			c.JSON(http.StatusNotFound, getErrorNoWordStartWith(firstLetter))
			return
		}

		c.JSON(http.StatusOK, mot)
	}
}

func getErrorNoWordStartWith(firstLetter string) gin.H {
	return gin.H{"error": fmt.Sprintf("No words start with a (%s)", firstLetter)}
}

func getErrorInvalidFirstLetter() gin.H {
	return gin.H{"error": InvalidFirstLetter}
}