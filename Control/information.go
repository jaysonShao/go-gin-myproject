package Control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"college_jouranlv2/Model"
	"fmt"
	"time"
)



func Postinfo(c *gin.Context) {
	var info Model.Information
	if c.BindJSON(&info) == nil {
		fmt.Println(info.School, info.Id, info.Title, info.Createtime, info.Usernickname, info.Userid, info.Usertelephone, info.Content, info.Praise)
		if len(info.Title) > 1 && len(info.School) > 1&& len(info.Usernickname) > 0 && len(info.Usertelephone) == 11 && len(info.Content) > 5 && info.Praise == "0"{
			info.Id = 0
			info.Createtime = time.Now().Format(_timestamp_format)
			info.Praise = "0"
			affect, err := engine.Insert(&info)
			if err != nil {
				panic(err)
			}
			fmt.Println(affect)
			c.JSON(http.StatusOK, gin.H{
				"code"		:	"200",
				"message"	:	"发布成功！",
				"data"		:	nil,
				"err"		:	nil,
				"time"		:    time.Now().Format(_timestamp_format),
			})
		}else {
			c.JSON(http.StatusOK, gin.H{
				"code"		:	"400",
				"message" 	:   nil,
				"err"		:	"请填写完整信息！",
				"data"		:	nil,
				"time"		: 	time.Now().Format(_timestamp_format),
			})
		}
	} else {
		err := c.BindJSON(&info)
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code"		:	"400",
			"message"	:	 nil,
			"err"		:    "请先填写信息！",
			"data"		:	nil,
			"time"		:    time.Now().Format(_timestamp_format),
		})
	}
}

func Getinfo(c *gin.Context){
	key := c.Param("key")
	if key == "master" {
		var info []Model.Information
		count,_ := engine.Count(Model.Information{})
		err := engine.Desc("id").Find(&info)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"code"		:	"200",
			"message"	:	count,
			"data"		:   info,
			"err" 		:	nil,
			"time"		: 	time.Now().Format(_timestamp_format),

		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code"		:	"401",
			"message"	: 	nil,
			"data"		:	nil,
			"err" 		:	"错误的请求",
			"time"		: 	time.Now().Format(_timestamp_format),
		})
	}
	
}

func Postpraise(c *gin.Context){
	var postdata Model.Information
	h := c.Request.Header.Get("table")
	if c.BindJSON(&postdata) == nil {
		fmt.Println(postdata.Praise, 	"as" ,postdata.Id, postdata.Createtime)
		if h == "information" {
			sql := "UPDATE information set praise = ? WHERE id = ?"
			res, err := engine.Exec(sql, postdata.Praise, postdata.Id)
			fmt.Println("information", res,err)
			c.JSON(http.StatusOK, gin.H{
				"code"		:	"200",
				"message"	:	"赞",
				"data"		:   nil,
				"err" 		:	nil,
				"time"		: 	time.Now().Format(_timestamp_format),
			})
		} else if h == "user" {
			sql := "UPDATE information set praise = ? WHERE id = ?"
			res, err := engine.Exec(sql, postdata.Praise, postdata.Id)
			fmt.Println("user", res,err)
			c.JSON(http.StatusOK, gin.H{
				"code"		:	"200",
				"message"	:	"赞",
				"data"		:   nil,
				"err" 		:	nil,
				"time"		: 	time.Now().Format(_timestamp_format),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code"		:	"400",
				"message"	:	nil ,
				"data"		:   nil,
				"err":     "错误的请求",
				"time":    time.Now().Format(_timestamp_format),
			})
		}
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code"		:	"400",
			"message"	:	nil ,
			"data"		:   nil,
			"err"		:   "错误的请求",
			"time"		:   time.Now().Format(_timestamp_format),
		})
	}

}
