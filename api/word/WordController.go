package word

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"strconv"
)

func SetUpRoutes(c *gin.Engine) {
	c.GET("/mots/:firstLetter", getWordsFirstLetter)
	c.GET("/mot/:firstLetter", getWordFirstLetter)
	c.GET("/mot/length/:length", getWordLength)
	c.GET("/anagrams/:word", getAnagrams)
}

const InvalidFirstLetter = "invalid first letter. Expected one character"

func getWordsFirstLetter(c *gin.Context) {
	firstLetter := c.Param("firstLetter")

	words, err := GetWordsFirstLetter(firstLetter)

	if err != nil {
		c.JSON(http.StatusBadRequest, getErrorInvalidFirstLetter())
		return
	}

	if len(words) == 0 {
		c.JSON(http.StatusNotFound, getErrorNoWordStartWith(firstLetter))
		return
	}

	c.JSON(http.StatusOK, words)

}

func getWordFirstLetter(c *gin.Context) {
	firstLetter := c.Param("firstLetter")

	word, err := GetWordFirstLetter(firstLetter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, getErrorNoWordStartWith(firstLetter))
		} else {
			c.JSON(http.StatusBadRequest, getErrorInvalidFirstLetter())
		}
		return
	}

	c.JSON(http.StatusOK, word)
}

func getWordLength(c *gin.Context) {
	length, err := strconv.Atoi(c.Param("length"))

	if err != nil {
		c.JSON(http.StatusBadRequest, getErrorInvalidLength())
		return
	}

	word, err := GetWordLength(length)

	if err != nil {
		c.JSON(http.StatusNotFound, getErrorNoWordWithLength(length))
		return
	}

	c.JSON(http.StatusOK, word)
}

func getAnagrams(c *gin.Context) {
	givenWord := c.Param("word")

	words, err := GetAnagrams(givenWord)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, getErrorNoAnagramFound(givenWord))
		} else {
			c.JSON(http.StatusBadRequest, getErrorNoWordLike(givenWord))
		}
		return
	}

	c.JSON(http.StatusOK, words)
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

func getErrorNoAnagramFound(word string) gin.H {
	return gin.H{"error": fmt.Sprintf("No anagram found for this word (%s)", word)}
}

func getErrorNoWordLike(word string) gin.H {
	return gin.H{"error": fmt.Sprintf("No match found for this word (%s)", word)}
}
