package Control

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine
var _timestamp_format = "2006-01-02 15:04:05"

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "username:password@tcp(127.0.0.1:3306)/college_journal?timeout=90s&charset=utf8")
	if err != nil {
		panic(err)
	}
}

func CloseEngine() {
	defer engine.Close()
}

func Returnjson(c *gin.Context, code string, message string, data interface{}, err string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
		"err":     err,
		"time":    time.Now().Format(_timestamp_format),
	})
}
