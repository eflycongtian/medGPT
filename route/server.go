package route

import (
	"github.com/gin-gonic/gin"
)

func RunServer(addr string) {
	engine := gin.Default()
	engine.LoadHTMLGlob("statics/*")
	initRoute(engine)
	engine.Run(addr)
}
