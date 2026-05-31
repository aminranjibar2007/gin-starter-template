package main

// Libs Import
import (
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
)

// main Block - Starter Block
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO If you want to use this server in public and release mode, it is better to set it to gin.SetMode(gin.ReleaseMode)

	gin.SetMode(gin.DebugMode)

	// Setting gin engine
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.RedirectFixedPath = true
	engine.HandleMethodNotAllowed = true
	engine.ForwardedByClientIP = true

	// Set LOG Middleware Plugin
	engine.Use(func(c *gin.Context) {
		date := c.GetHeader("Date")
		ip := c.ClientIP()

		log.Println(date, " ", ip)
	})

	// Set HOST Check Middleware Plugin
	engine.Use(func(c *gin.Context) {
		host := c.Request.Host

		// TODO If you want to increase the security of your server and prevent DDos operations with the IP domains that connect to your server

		switch host {
		case "www.baidu.com":
			c.Next()
		default:
			c.AbortWithStatus(403)
		}
	})

	// TODO You can set up your Jain server routes with these commands. What do you want to do ?
	//engine.GET("/", func(c *gin.Context) {})
	//engine.Static("/static", "./static")
	//engine.Group("/", func(c *gin.Context) {})
	//engine.POST("/", func(c *gin.Context) {})
	//engine.Any("/", func(c *gin.Context) {})

	// Run server on port 80 on all ip
	err := engine.Run(":80")

	if err != nil {
		panic(err)
	}

	// TODO If you want to have a secure server, it is better to use TLS and SSL
	//err := engine.RunTLS(":443", "./cert/server.pem", "./cert/server.key")
	//if err != nil {
	//	panic(err)
	//}
}
