package mot

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"strconv"
)

func SetUpRoutes(c *gin.Engine) {
	c.GET("/mots/:firstLetter", GetMotsFirsLetter)
	c.GET("/mot/:firstLetter", GetMotFirsLetter)
	c.GET("/mot/length/:length", getMotLength)
}

const InvalidFirstLetter = "invalid first letter. Expected one character"

func GetMotsFirsLetter(c *gin.Context) {
	firstLetter := c.Param("firstLetter")

	mots, err := GetMotsFirstLetter(firstLetter)

	if err != nil {
		c.JSON(http.StatusBadRequest, getErrorInvalidFirstLetter())
		return
	}

	if len(mots) == 0 {
		c.JSON(http.StatusNotFound, getErrorNoWordStartWith(firstLetter))
		return
	}

	c.JSON(http.StatusOK, mots)

}

func GetMotFirsLetter(c *gin.Context) {
	firstLetter := c.Param("firstLetter")

	mot, err := GetMotFirstLetter(firstLetter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, getErrorNoWordStartWith(firstLetter))
		} else {
			c.JSON(http.StatusBadRequest, getErrorInvalidFirstLetter())
		}
		return
	}

	c.JSON(http.StatusOK, mot)
}

func getMotLength(c *gin.Context) {
	length, err := strconv.Atoi(c.Param("length"))

	if err != nil {
		c.JSON(http.StatusBadRequest, getErrorInvalidLength())
		return
	}

	mot, err := GetMotLength(length)

	if err != nil {
		c.JSON(http.StatusNotFound, getErrorNoWordWithLength(length))
		return
	}

	c.JSON(http.StatusOK, mot)
}

func getErrorNoWordStartWith(firstLetter string) gin.H {
	return gin.H{"error": fmt.Sprintf("No words start with a (%s)", firstLetter)}
}

func getErrorInvalidFirstLetter() gin.H {
	return gin.H{"error": InvalidFirstLetter}
}

func getErrorNoWordWithLength(length int) gin.H {
	return gin.H{"error": fmt.Sprintf("No words with length (%d)", length)}
}

func getErrorInvalidLength() gin.H {
	return gin.H{"error": "Please give a number"}
}
