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
	c.GET("/words/:firstLetter", getWordsFirstLetter)
	c.GET("/word/:firstLetter", getWordFirstLetter)
	c.GET("/word/length/:length", getWordLength)
	c.GET("/anagrams/:word", getAnagrams)
	c.GET("/words-batch/:letters", getWordsBatch)
}

const InvalidFirstLetter = "invalid first letter. Expected one character"

func getWordsFirstLetter(c *gin.Context) {
	firstLetter := c.Param("firstLetter")

	words, err := GetWordsFirstLetter(firstLetter)

	if err != nil {
		c.JSON(http.StatusBadRequest, invalidFirstLetter())
		return
	}

	if len(words) == 0 {
		c.JSON(http.StatusNotFound, noWordStartWith(firstLetter))
		return
	}

	c.JSON(http.StatusOK, words)

}

func getWordFirstLetter(c *gin.Context) {
	firstLetter := c.Param("firstLetter")

	word, err := GetWordFirstLetter(firstLetter)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, noWordStartWith(firstLetter))
		} else {
			c.JSON(http.StatusBadRequest, invalidFirstLetter())
		}
		return
	}

	c.JSON(http.StatusOK, word)
}

func getWordLength(c *gin.Context) {
	length, err := strconv.Atoi(c.Param("length"))

	if err != nil {
		c.JSON(http.StatusBadRequest, invalidLength())
		return
	}

	word, err := GetWordLength(length)

	if err != nil {
		c.JSON(http.StatusNotFound, noWordWithLength(length))
		return
	}

	c.JSON(http.StatusOK, word)
}

func getAnagrams(c *gin.Context) {
	givenWord := c.Param("word")

	words, err := GetAnagrams(givenWord)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, noAnagramFound(givenWord))
		} else {
			c.JSON(http.StatusBadRequest, noWordLike(givenWord))
		}
		return
	}

	c.JSON(http.StatusOK, words)
}

func getWordsBatch(c *gin.Context) {
	letters := c.Param("letters")

	words := GetWordsBatch(letters)

	if len(words) == 0 {
		c.JSON(http.StatusNotFound, noWordsStartWithGivenLetter(letters))
		return
	}

	c.JSON(http.StatusOK, words)
}

func noWordStartWith(firstLetter string) gin.H {
	return gin.H{"error": fmt.Sprintf("No words start with a (%s)", firstLetter)}
}

func invalidFirstLetter() gin.H {
	return gin.H{"error": InvalidFirstLetter}
}

func noWordWithLength(length int) gin.H {
	return gin.H{"error": fmt.Sprintf("No words with length (%d)", length)}
}

func invalidLength() gin.H {
	return gin.H{"error": "Please give a number"}
}

func noAnagramFound(word string) gin.H {
	return gin.H{"error": fmt.Sprintf("No anagram found for this word (%s)", word)}
}

func noWordLike(word string) gin.H {
	return gin.H{"error": fmt.Sprintf("No match found for this word (%s)", word)}
}

func noWordsStartWithGivenLetter(letters string) gin.H {
	return gin.H{"error": fmt.Sprintf("No words start with given letters (%s)", letters)}

}