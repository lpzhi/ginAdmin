package api

import (
	"ginAdmin/models"
	"ginAdmin/pkg/e"
	"ginAdmin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type  auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}



func GetAuth(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	//a := auth{username,password}

	a := auth{Username: username, Password: password}

	log.Println(a)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if 	ok,_:=valid.Valid(&a);ok {

		if isExit := models.CheckAuth(username,password);isExit {
			 if token,err := util.GenerateToke(username,password);err==nil{
			 	data["token"] = token
			 	code = e.SUCCESS
			 }else {
				code = e.ERROR_AUIH_TOKEN
			 }


		}
	}else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}