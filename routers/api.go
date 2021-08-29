package routers

import (
	"github.com/gin-gonic/gin"
	"github/demo/mq"
	"github/demo/utils"
	"net/http"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.POST("/send-mq", func(c *gin.Context) {
		data := utils.BackData{}
		err, msg := mq.SendMQ{}.Send()

		if err != nil {
			data.Status = http.StatusBadRequest
			data.Message = err.Error()
			c.JSON(http.StatusBadRequest, data)
			c.JSON(http.StatusOK, "失败")
			return
		}
		data.Status = http.StatusOK
		data.Message = "成功"
		data.Data = msg
		c.JSON(http.StatusOK, data)

	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	return r
}
