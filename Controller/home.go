package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "!!!!!Hello World!!!!!")
}
