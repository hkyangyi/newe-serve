package middle

import (
	"fmt"
	"newe-serve/app/model"
	"newe-serve/common/redis"
	"newe-serve/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Result    interface{} `json:"result"`
	Success   string      `json:"success"`
	Timestamp int64       `json:"timestamp"`
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		uuid, err := utils.AuthToken(token)
		if !err {
			c.JSON(401, Response{
				Code:      10000,
				Message:   "登录超时",
				Result:    nil,
				Success:   "fail",
				Timestamp: time.Now().Unix(),
			})
			c.Abort()
			return
		}
		fmt.Printf("middle AdminMid uuid：%s uid \n", uuid)
		//检测KEY是否在缓存中
		res := redis.Exists(uuid)

		if !res {
			c.JSON(401, Response{
				Code:      10000,
				Message:   "登录超时",
				Result:    nil,
				Success:   "fail",
				Timestamp: time.Now().Unix(),
			})
			c.Abort()
			return
		}
		//从缓存中拿取数据
		var data model.SysMember
		errs := redis.Get(uuid, &data)
		if errs != nil {
			c.JSON(401, Response{
				Code:      10000,
				Message:   "登录超时",
				Result:    nil,
				Success:   "fail",
				Timestamp: time.Now().Unix(),
			})
			c.Abort()
			return
		}
		redis.Set(uuid, data, 3600)
		c.Set("AdminAuthData", data)
		c.Next()
	}
}
