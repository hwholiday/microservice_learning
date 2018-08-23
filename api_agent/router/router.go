package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"microservice_learning/api_agent/controller"
)

func SetRouters(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	test := &controller.TestController{}
	v1 := r.Group("/api/v1")
	{
		v1.GET("/test", test.Test)
		v1.GET("/any", test.Anything)
	}
}
/*func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.Request.Header.Get("Authorization")
		if len(authString) == 0 {
			c.JSON(401, gin.H{"code": 401, "data": "", "msg": "未获取到授权信息"})
			c.Abort()
		} else {
			_, err := utils.ValidToken(authString)
			if err != nil {
				c.JSON(401, gin.H{"code": 401, "data": "", "msg": "会话已经过期,请重新登录"})
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}*/
