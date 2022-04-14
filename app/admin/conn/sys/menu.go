package sys

import (
	core "newe-serve/app/admin/core/sys"
	"newe-serve/common/app"

	"github.com/gin-gonic/gin"
)

func MenuAdd(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysMenus
	)

	if err := app.BindAndValid(c, &d); err != nil {
		g.Error(err)
	}
	err := d.Create()
	if err != nil {
		g.Error(err)
		return
	}
	g.Success(nil)
	return
}

func MenuGetList(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysMenus
	)
	if err := app.BindAndValid(c, &d); err != nil {
		g.Error(err)
	}
	items := d.GetList()
	g.Success(items)
	return
}

func MenuEdit(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysMenus
	)

	if err := app.BindAndValid(c, &d); err != nil {
		g.Error(err)
	}
	err := d.Edit()
	if err != nil {
		g.Error(err)
		return
	}
	g.Success(nil)
	return
}

func MenuDel(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysMenus
	)

	if err := app.BindAndValid(c, &d); err != nil {
		g.Error(err)
	}
	err := d.Del()
	if err != nil {
		g.Error(err)
		return
	}
	g.Success(nil)
	return
}
