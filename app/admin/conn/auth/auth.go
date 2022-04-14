package auth

import (
	coresys "newe-serve/app/admin/core/sys"
	"newe-serve/common/app"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data = make(map[string]interface{})
	var g = app.Gin{C: c}

	data["roles"] = []int{100, 200, 300}
	data["userId"] = 10000
	data["username"] = "neweadmin"
	data["realName"] = "杨一"
	data["avatar"] = ""
	data["desc"] = "介绍介绍"
	data["token"] = "asdfasdfaeasdfawefasefawsef"
	g.Success(data)
}

func GetUserInfo(c *gin.Context) {
	var data = make(map[string]interface{})
	var g = app.Gin{C: c}
	data["userId"] = '1'
	data["username"] = "vben"
	data["realName"] = "Vben Admin"
	data["avatar"] = "https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640"
	data["desc"] = "manager"
	data["password"] = "123456"
	data["token"] = "fakeToken1"
	data["homePath"] = "/dashboard/analysis"
	//   roles: [
	//     {
	//       roleName: 'Super Admin',
	//       value: 'super',
	//     },
	//   ],
	g.Success(data)
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
	data := coresys.GetMenuList()
	g.Success(data)
}
