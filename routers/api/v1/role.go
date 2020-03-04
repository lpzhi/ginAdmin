package v1

import (
	"ginAdmin/models"
	"ginAdmin/pkg/app"
	"ginAdmin/pkg/e"
	"ginAdmin/pkg/setting"
	"ginAdmin/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)



// @Summary Get multiple roles
// @Produce  json

// @Security ApiKeyAuth
// @Param name query string false "Name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/roles [get]
func GetRoles(c *gin.Context)  {
	appG := app.Gin{C:c}

	roleName := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if roleName!="" {
		maps["role_name"] = roleName
	}

	data["lists"] = models.GetRoles(util.GetPage(c),setting.AppSetting.PageSize,maps)
	data["total"] = models.GetRoleTotal()

	appG.Response(http.StatusOK,e.SUCCESS,data)

}
