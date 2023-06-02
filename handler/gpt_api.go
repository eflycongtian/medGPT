package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sinohealth.com/medGPT/model"
)

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chat.tmpl", gin.H{
		"code":  http.StatusOK,
		"title": "OpenAI ChatGPT 医生接诊助手",
	})
}

func GenRx(ctx *gin.Context) {
	content, _ := ctx.GetQuery("content")
	content = "生成一个治疗" + content + "的西药处方，表头包含：药品名称、规格、厂家、用法用量、天数、数量、注意事项,直接只输出json数据,数据格式:['drugName':'','spec':'','company':'','usage':'','days':'','quantity':'','notice':''],不要多余文本内容,注意json格式正确"
	resp, err := model.CreateChatCompletion(content)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": resp.Choices[0].Message.Content,
			"usage":   resp.Usage,
		})
	}
}

func GenMedicalRecord(ctx *gin.Context) {
	content, _ := ctx.GetQuery("content")
	content = "患者主诉为：" + content + "，请为该患者生成一个病历记录，病历记录包含：主诉、现病史、既往史、查体、诊断、处理意见,直接只输出json数据结构,数据格式:{'complain':'','PIH':'','PMH':'','bodyCheck':'','diagnose':'','suggest':''}不要多余文本内容,注意json格式正确"
	resp, err := model.CreateChatCompletion(content)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": resp.Choices[0].Message.Content,
			"usage":   resp.Usage,
		})
	}
}
