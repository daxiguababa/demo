package routers

import (
	"github.com/gin-gonic/gin"
	"github/demo/dao"
	"github/demo/mq"
	"github/demo/utils"
	"net/http"
	"time"
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

	r.GET("/redis", func(c *gin.Context) {
		err := dao.Redis.Set("demo", time.Now().UnixNano(), -1).Err()
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, "ok")
	})

	r.GET("/mysql", func(c *gin.Context) {

		var total int64
		err := dao.DB.Table("users").Select("id").Count(&total).Error

		if total > 0 && err == nil {
			c.JSON(http.StatusOK, total)
			return
		}
		c.JSON(http.StatusOK, err.Error())

	})

	r.GET("/mongo", func(c *gin.Context) {

		data := dao.MonConData{Name: "this is demo"}

		insertResult, err := dao.Mongo.Collection("new_demo").InsertOne(dao.MonCon{}.Con(), &data)
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
		}
		c.JSON(http.StatusOK, insertResult)
	})
	return r
}
