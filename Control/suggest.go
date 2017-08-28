package Control

import (
	"github.com/gin-gonic/gin"
	"college_jouranlv2/Model"
	"net/http"
	"time"
	"fmt"
)

func Postsuggest(c *gin.Context)  {
	var suggest Model.Suggest
	if c.BindJSON(&suggest) == nil {
		suggest.Createtime = time.Now().Format(_timestamp_format)
		affect ,err := engine.InsertOne(&suggest)
		if err != nil {
			fmt.Println("err:",err)
			c.JSON(http.StatusOK, gin.H{
				"code"		:	"400",
				"message"	:	nil,
				"data"		:   nil,
				"err" 		:	"数据不正确",
				"time"		: 	time.Now().Format(_timestamp_format),
			})
		}else {
			fmt.Println("affect:",affect)
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "谢谢您的意见！",
				"data":    nil,
				"err":     nil,
				"time":    time.Now().Format(_timestamp_format),
			})
		}
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":    "400",
			"message": nil,
			"data":    nil,
			"err":     "错误的链接",
			"time":    time.Now().Format(_timestamp_format),
		})
	}
}
