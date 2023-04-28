package route

import (
	"github.com/gin-gonic/gin"
	"sinohealth.com/medGPT/handler"
)

func initRoute(engine *gin.Engine) {
	engine.GET("/ping", handler.GetPing)
	engine.GET("gpt/home", handler.Home)
	engine.GET("gpt/genRx", handler.GenRx)
	engine.GET("gpt/genMedicalRecord", handler.GenMedicalRecord)
}
