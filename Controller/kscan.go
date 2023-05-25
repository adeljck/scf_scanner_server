package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scf_scanner_server_gin/Modules"
)

func KScan(c *gin.Context) {
	var Scanner Modules.Scanner
	if err := c.BindQuery(&Scanner); err == nil {
		Scanner.KScan()
		c.String(http.StatusOK, Scanner.Results)
	} else {
		c.JSON(http.StatusOK, Scanner)
	}
}
