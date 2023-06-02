package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer(addr string) {
	engine := gin.Default()
	engine.LoadHTMLGlob("statics/*")
	engine.Use(cors.Default())
	initRoute(engine)
	engine.Run(addr)
}
