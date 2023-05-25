package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scf_scanner_server_gin/Modules"
)

func GetIPInfo(c *gin.Context) {
	var IPInfo Modules.IPInfo
	IPInfo.GetCurrentRequestIPInfo()
	c.JSON(http.StatusOK, IPInfo)
}
