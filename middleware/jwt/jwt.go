package jwt


import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"ginAdmin/pkg/e"
	"ginAdmin/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
	//	c.Request.Header.Get("Authorization")
		//token := c.Query("token")
		token :=c.Request.Header.Get("Authorization")
		fmt.Println(token)
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUIH_CHECK_TOKE_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKE_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}