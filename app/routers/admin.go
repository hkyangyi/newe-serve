package routers

import (
	"newe-serve/app/admin/conn/auth"
	"newe-serve/app/admin/conn/sys"
	"newe-serve/app/middle"

	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.RouterGroup) {
	r.POST("login", auth.Login)
	r.GET("logout", auth.LoginOut)
	r.Use(middle.AdminAuth())
	r.GET("getUserInfo", auth.GetUserInfo)
	r.GET("getPermCode", auth.GetPermCode)
	r.GET("getMenuList", auth.GetMenuList)
	r.GET("verifysole", auth.Verifysole)
	sys := r.Group("sys")
	NeweSys(sys)
}

//系统设置
func NeweSys(r *gin.RouterGroup) {
	r.POST("meadd", sys.MenuAdd)
	r.GET("megetlist", sys.MenuGetList)
	r.PUT("meedit", sys.MenuEdit)
	r.DELETE("medel", sys.MenuDel)
	//组织结构
	r.POST("deptadd", sys.DepartAdd)
	r.PUT("deptedit", sys.DepartEdit)
	r.DELETE("deptdel", sys.DepartDel)
	r.GET("deptgetlist", sys.DepartGetList)
	r.GET("deptgetrules", sys.DepartRulesGet)
	r.POST("deptrulessave", sys.DepartRulesSave)
	//账号管理
	r.POST("memberadd", sys.MemberAdd)
	r.PUT("memberedit", sys.MemberEdit)
	r.GET("membergetlist", sys.MemberGetList)
	r.DELETE("memberdel", sys.MemberDel)
}
