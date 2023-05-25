package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scf_scanner_server/Modules"
)

func FScan(c *gin.Context) {
	var Scanner Modules.Scanner
	if err := c.BindQuery(&Scanner); err == nil {
		Scanner.FScan()
		c.String(http.StatusOK, Scanner.Results)
	} else {
		c.JSON(http.StatusOK, Scanner)
	}
}
