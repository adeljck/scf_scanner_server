package Router

import (
	"github.com/gin-gonic/gin"
	"scf_scanner_server_gin/Controller"
)

func InitRouter() {
	router := gin.Default()
	router.GET("/f", Controller.FScan)
	router.GET("/k", Controller.KScan)
	router.GET("/ip", Controller.GetIPInfo)
	router.Run(":9000")
}
