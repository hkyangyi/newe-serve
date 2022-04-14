package routers

import (
	"newe-serve/app/admin/conn/auth"
	"newe-serve/app/admin/conn/sys"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.RouterGroup) {
	r.POST("login", auth.Login)
	r.GET("getUserInfo", auth.GetUserInfo)
	r.GET("getPermCode", auth.GetPermCode)
	r.GET("getMenuList", auth.GetMenuList)
	sys := r.Group("sys")
	NeweSys(sys)
}

//系统设置
func NeweSys(r *gin.RouterGroup) {
	r.POST("meadd", sys.MenuAdd)
	r.GET("megetlist", sys.MenuGetList)
	r.PUT("meedit", sys.MenuEdit)
	r.DELETE("medel", sys.MenuDel)
}
