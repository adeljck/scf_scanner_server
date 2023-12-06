package ip

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scf_scanner_server/Modules/ip"
)

func GetIPInfo(c *gin.Context) {
	var IPInfo ip.IPInfo
	IPInfo.GetCurrentRequestIPInfo()
	c.JSON(http.StatusOK, IPInfo)
}
