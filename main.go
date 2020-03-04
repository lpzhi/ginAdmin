package main

import (
	"fmt"
	"ginAdmin/pkg/e"
	"ginAdmin/pkg/setting"
	"ginAdmin/routers"
	"github.com/gin-gonic/gin"

	"net/http"
)


// @host localhost:8090
// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main()  {


	router := routers.InitRouter()//gin.Default()

	router.GET("/test", func(context *gin.Context) {
		context.JSON(e.SUCCESS,gin.H{
			"message":"tes1t",
		})
	})


	s := &http.Server{
		Addr:fmt.Sprintf(":%d",setting.ServerSetting.HttpPort),
		Handler:router,
		//ReadTimeout:setting.ReadTimeout,
		//WriteTimeout:setting.WriteTimeout,
		MaxHeaderBytes:1<<20,
	}

	s.ListenAndServe()
}
