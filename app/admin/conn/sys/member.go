package sys

import (
	core "newe-serve/app/admin/core/sys"
	"newe-serve/common/app"

	"github.com/gin-gonic/gin"
)

func MemberAdd(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysMember
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

func MemberEdit(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysMember
	)

	if err := app.BindAndValid(c, &d); err != nil {
		g.Error(err)
	}
	err := d.Update()
	if err != nil {
		g.Error(err)
		return
	}
	g.Success(nil)
	return
}

func MemberDel(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysMember
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

func MemberGetList(c *gin.Context) {
	var (
		g = app.Gin{C: c}
		d core.SysMember
	)

	if err := app.BindAndValid(c, &d); err != nil {
		g.Error(err)
	}
	items := d.GetList()

	g.Success(items)
	return
}
