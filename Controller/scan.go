package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"scf_scanner_server/Modules"
)

func Scan(c *gin.Context) {
	Scanner := new(Modules.Scanner)
	if err := c.MustBindWith(Scanner, binding.JSON); err == nil {
		Scanner.Scan()
		c.String(http.StatusOK, Scanner.Results)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusExpectationFailed, "msg": err})
	}
}
