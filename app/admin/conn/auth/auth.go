package auth

import (
	"errors"
	coreauth "newe-serve/app/admin/core/auth"
	coresys "newe-serve/app/admin/core/sys"
	"newe-serve/app/model"
	"newe-serve/common/app"
	"newe-serve/common/db"
	"newe-serve/common/redis"
	"newe-serve/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var (
		a coreauth.AuthFrom
		g = app.Gin{C: c}
	)
	if e := app.BindAndValid(c, &a); e != nil {
		g.Error(errors.New("参数错误"))
		return
	}
	data, err := a.Login()
	if err != nil {
		g.Error(err)
		return
	}
	g.Success(data)
	return

	// data["roles"] = []int{100, 200, 300}
	// data["userId"] = 10000
	// data["username"] = "neweadmin"
	// data["realName"] = "杨一"
	// data["avatar"] = ""
	// data["desc"] = "介绍介绍"
	// data["token"] = "asdfasdfaeasdfawefasefawsef"

}

func LoginOut(c *gin.Context) {
	var g = app.Gin{C: c}
	token := c.GetHeader("Authorization")
	uuid, err := utils.AuthToken(token)
	if !err {
		g.LoginError(nil)
		return
	}

	res := redis.Exists(uuid)
	if !res {
		g.LoginError(nil)
		return
	}
	redis.Delete(uuid)
	g.Success(nil)
	return

}

func GetUserInfo(c *gin.Context) {

	var g = app.Gin{C: c}
	merdata, b := c.Get("AdminAuthData")
	if !b {
		g.LoginError(errors.New("登陆超时"))
		return
	}

	mer := merdata.(model.SysMember)
	g.Success(mer)
	return

	// var data = make(map[string]interface{})
	// var g = app.Gin{C: c}
	// data["userId"] = '1'
	// data["username"] = "admin"
	// data["realName"] = "neweadmin Admin"
	// data["avatar"] = "https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640"
	// data["desc"] = "manager"
	// data["password"] = "123456"
	// data["token"] = "fakeToken1"
	// data["homePath"] = "/dashboard/analysis"
	// //   roles: [
	// //     {
	// //       roleName: 'Super Admin',
	// //       value: 'super',
	// //     },
	// //   ],
	//g.Success(data)
}

func GetPermCode(c *gin.Context) {

	var g = app.Gin{C: c}

	//   roles: [
	//     {
	//       roleName: 'Super Admin',
	//       value: 'super',
	//     },
	//   ],['1000', '3000', '5000']
	var data = []string{"1000", "3000", "5000"}
	g.Success(data)
}

func GetMenuList(c *gin.Context) {
	var g = app.Gin{C: c}
	merdata, b := c.Get("AdminAuthData")
	if !b {
		g.LoginError(errors.New("登陆超时"))
		return
	}

	mer := merdata.(model.SysMember)
	data := coresys.GetMenuList(mer.DepartId)
	g.Success(data)
}

type verifyform struct {
	TableName  string `form:"tablename" valid:"Required; MaxSize(50)"`
	FieldName  string `form:"fieldname" valid:"Required; MaxSize(50)"`
	Tablevalue string `form:"tablevalue" valid:"Required; MaxSize(50)"`
	TableId    string `form:"tableid"`
}

// @验证字段唯一

// @Tags Base
// @Description 验证字段是否存在
// @Accept  json
// @Produce json
// @Param tablename  path   string    true   "表名字"
// @Param fieldname  path   string    true   "字段名字"
// @Param tablevalue  path   string    true   "字段值"
// @Param tableid  path   int    true  "过滤ID 在编辑时传入ID"
// @Success 200 {string} string	"ok"
// @Router /api/base/verifysole [get]
func Verifysole(c *gin.Context) {
	var (
		a verifyform
		g = app.Gin{C: c}
	)
	if e := app.BindAndValid(c, &a); e != nil {
		g.Error(errors.New("参数错误"))
		return
	}

	where := make(map[string]interface{})
	where[a.FieldName] = a.Tablevalue

	res := db.VerifyOnly(a.TableName, a.TableId, where)
	if res {
		g.Error(errors.New("已存在"))
		return
	}
	g.Success(nil)

}
