package tickets

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "create ticket"})
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "get tickets"})
}

func GetById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "get ticket by id"})
}
