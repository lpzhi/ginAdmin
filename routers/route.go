package routers

import (
	_ "ginAdmin/docs"
	"ginAdmin/middleware/jwt"
	"ginAdmin/pkg/setting"
	"ginAdmin/routers/api"
	"ginAdmin/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	//"github.com/swaggo/gin-swagger"

	//"github.com/swaggo/gin-swagger"

)

func InitRouter() *gin.Engine  {
	r := gin.New()
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	url := ginSwagger.URL("http://localhost:8090/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

    r.GET("/auth",api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 :=r.Group("/api/v1").Use(jwt.JWT())
	{

		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/roles",v1.GetRoles)

		//创建数据表
		apiv1.GET("roletotal", v1.CreateRoleTotal)
		apiv1.GET("enterlog", v1.CreateEnterLog)
	}


	return r
}