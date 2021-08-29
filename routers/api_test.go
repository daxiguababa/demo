package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github/demo/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 真正的测试单元
func TestAPI(t *testing.T) {
	router := gin.Default()
	SetupRouter(router)
	//address, _ := net.InterfaceAddrs()
	//for _, v := range address {
	//	// 检查ip地址判断是否回环地址
	//	if ip, ok := v.(*net.IPNet); ok && !ip.IP.IsLoopback() {
	//		if ip.IP.To4() != nil {
	//			fmt.Println(ip.IP.String())
	//		}
	//	}
	//}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/send-mq", nil)
	router.ServeHTTP(w, req)
	data := utils.BackData{}

	_ = json.Unmarshal(w.Body.Bytes(), &data)
	fmt.Println(data)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "成功", data.Message)

}
