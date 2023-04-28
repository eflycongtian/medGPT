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
	content = "生成一个治疗" + content + "的西药处方，表头包含：药品名称、规格、厂家、用法用量、天数、数量、注意事项，输出html表格"
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
	content = "患者主诉为：" + content + "，请为该主诉生成一个病历记录，病历记录包含：主诉、现病史、既往史、查体、诊断、处理意见，输出格式为一个html div片段，其中诊断内容的<p>标签加上id='diagnose-text'"
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
