package Control

import (
	"github.com/gin-gonic/gin"
	"college_jouranlv2/Model"
	"net/http"
	"time"
	"fmt"
	"strconv"
	"crypto/md5"
	"encoding/hex"
)

func Postuser(c *gin.Context){
	var postdate Model.User
	if c.BindJSON(&postdate) == nil  {
		var querydata Model.User
		res,err := engine.Id(postdate.Id).Get(&querydata)
		if err != nil {
			fmt.Println("err:",err)
			c.JSON(http.StatusOK, gin.H{
				"code":    "503",
				"message": nil,
				"data":    nil,
				"err":     "服务器错误，请稍后！",
				"time":    time.Now().Format(_timestamp_format),
			})
		}else {
			fmt.Println("查询res",res)
			if postdate.Telephone == querydata.Telephone {
				c.JSON(http.StatusOK, gin.H{
					"code":    "200",
					"message": "您的手机已经注册！！",
					"data":    nil,
					"err":     nil,
					"time":    time.Now().Format(_timestamp_format),
				})
			}else {
				postdate.Createtime = time.Now().Format(_timestamp_format)
				h := md5.New()
				h.Write([]byte(postdate.Password))
				enRes := hex.EncodeToString(h.Sum(nil))
				postdate.Password = enRes
				ress,err := engine.InsertOne(postdate)
				if err != nil {
					fmt.Println("插入err:",err)
					c.JSON(http.StatusOK, gin.H{
						"code":    "503",
						"message": nil,
						"data":    nil,
						"err":     "服务器错误，请稍后再试！",
						"time":    time.Now().Format(_timestamp_format),
					})
				}else {
					fmt.Println("插入res",ress)
					c.JSON(http.StatusOK, gin.H{
						"code":    "200",
						"message": "注册成功！",
						"data":    nil,
						"err":     nil,
						"time":    time.Now().Format(_timestamp_format),
					})
				}
			}
		}
	}else {
		c.JSON(http.StatusOK, gin.H{
			"code":    "400",
			"message": nil,
			"data":    nil,
			"err":     "请求体有误",
			"time":    time.Now().Format(_timestamp_format),
		})
	}

}

func Getlogin(c *gin.Context){
	var postdata Model.GetUser
	if c.BindJSON(&postdata) == nil {
		var checkdata Model.User
		res, err := engine.Id(postdata.Id).Get(&checkdata)
		if err != nil {
			fmt.Println("err:",err)
			Returnjson(c,"503","",nil,"服务器错误，请稍后重试！")
		}else {
			fmt.Println("Res:",res)
			fmt.Println(checkdata)
			if postdata.Telephone == checkdata.Telephone {
				h := md5.New()
				h.Write([]byte(postdata.Password))
				enRes := hex.EncodeToString(h.Sum(nil))
				postdata.Password = enRes
				if postdata.Password == checkdata.Password {
					Returnjson(c,"200","登陆成功",nil,"")
				}else {
					Returnjson(c, "200","",nil,"密码错误！")
				}
			}else {
				Returnjson(c,"200","您还没有注册！",nil,"")
			}

		}
	}else {
		fmt.Println(c.BindJSON(&postdata))
		Returnjson(c,"400","", nil,  "请求体错误！")
	}

}

func Putuser(c *gin.Context){
	var postdata Model.User
	if c.BindJSON(&postdata) == nil {
		postdata.Createtime = time.Now().Format(_timestamp_format)
		res,err := engine.Id(postdata.Id).Update(postdata)
		if err != nil {
			fmt.Println("err:",err)
			Returnjson(c,"503","",nil,"服务器错误，请稍后重试！")
		}else {
			fmt.Println("Res:",res)
			Returnjson(c,"200","更新成功",nil,"")
		}
	}else {
		Returnjson(c,"400","",nil,"请求体错误")
	}

}

func Putuse(c *gin.Context){
	id := c.Param("id")
	var idd int;
	idd, _ = strconv.Atoi(id)
	var postdata Model.Putsign
	if c.BindJSON(&postdata) == nil {
		sql := "UPDATE user set signature = ? where id =?"
		res, err := engine.Exec(sql,postdata.Signature,idd)
		if err != nil {
			fmt.Println("err:", err)
			Returnjson(c, "503", "", nil, "服务器错误，请稍后重试！")
		} else {
			fmt.Println("res:", res)
			Returnjson(c, "200", "更新资料成功！", nil, "")
		}
	}else {
		Returnjson(c,"400","",nil,"请求体错误")
	}

}

func Putpwd(c *gin.Context){
	old := c.Param("old")
	new := c.Param("new")
	id := c.Param("id")
	var idd int;
	idd, _ = strconv.Atoi(id)
	var checkdata Model.User
	res, err := engine.Id(idd).Get(&checkdata)
	if err != nil {
		fmt.Println("queryerr:",err)
		Returnjson(c,"503","",nil,"服务器错误，请稍后重试！")
	}else {
		fmt.Println("res:",res)
		h := md5.New()
		h.Write([]byte(old))
		enRes := hex.EncodeToString(h.Sum(nil))
		if enRes == checkdata.Password {
			hh := md5.New()
			hh.Write([]byte(new))
			enNewRes := hex.EncodeToString(hh.Sum(nil))
			sql1 := "UPDATE user set password = ? where id = ?"
			 res, err := engine.Exec(sql1,enNewRes,idd)
			if err != nil {
				fmt.Println("err:",err)
				Returnjson(c,"503","",nil,"服务器错误，请稍后重试！")
			}else {
				fmt.Println("res:",res)
				Returnjson(c,"200","密码更新成功！",nil,"")
			}
		}else {
			Returnjson(c,"200","",nil,"密码错误！")
		}
	}
}

func Getuser(c *gin.Context){
	hander := c.Request.Header.Get("authority")
	fmt.Println(hander)
	if hander == "normal"{
		var userinfo []Model.User
		err := engine.Find(&userinfo)
		if err != nil {
			fmt.Println("err:",err)
			Returnjson(c,"503","",nil,"服务器错误，请稍后再试！")
		}else {
			s := len(userinfo)
			var ss string
			ss = strconv.Itoa(s)
			for  i:=0; i < len(userinfo) ;i++{
				userinfo[i].Password = "不能看!"
			}
			fmt.Println(userinfo)
			Returnjson(c,"200",ss, userinfo,"")
		}
	}else {
		Returnjson(c,"400","",nil,"错误的链接")
	}
}

