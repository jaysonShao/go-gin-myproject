package Control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"fmt"
	"college_jouranlv2/Model"
)

func Getupdate(c *gin.Context){
	var data []Model.Updatee
	err := engine.Find(&data)
		if err != nil{
			fmt.Println("err:",err)
			c.JSON(http.StatusOK, gin.H{
				"code"		:	"400",
				"message"	:	nil,
				"data"		:   nil,
				"err" 		:	"服务器错误，请稍后！",
				"time"		: 	time.Now().Format(_timestamp_format),
			})
		}else {
			fmt.Println("res",data)
			c.JSON(http.StatusOK, gin.H{
				"code"		:	"200",
				"message"	:	nil,
				"data"		:   data,
				"err" 		:	nil,
				"time"		: 	time.Now().Format(_timestamp_format),
			})

		}
}
