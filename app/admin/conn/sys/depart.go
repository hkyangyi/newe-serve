package sys

import (
	core "newe-serve/app/admin/core/sys"
	"newe-serve/common/app"

	"github.com/gin-gonic/gin"
)

func DepartAdd(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysDepart
	)

	if err := app.BindAndValid(c, &d); err != nil {
		g.Error(err)
	}
	err := d.Add()
	if err != nil {
		g.Error(err)
		return
	}
	g.Success(nil)
	return
}

func DepartEdit(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysDepart
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

func DepartDel(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysDepart
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

func DepartGetList(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysDepart
	)

	if err := app.BindAndValid(c, &d); err != nil {
		g.Error(err)
	}
	items := d.GetList()

	g.Success(items)
	return
}
