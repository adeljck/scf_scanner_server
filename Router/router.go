package Router

import (
	"github.com/gin-gonic/gin"
	"scf_scanner_server/Controller"
	"scf_scanner_server/Controller/ip"
)

func InitRouter() {
	router := gin.Default()
	router.Any("/", Controller.HomePage)
	router.POST("/scan", Controller.Scan)
	router.GET("/ip", ip.GetIPInfo)
	router.Run(":9000")
}
