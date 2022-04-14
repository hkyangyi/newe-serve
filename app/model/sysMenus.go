package model

import (
	"newe-serve/common/db"
	"newe-serve/pkg/utils"
)

type SysMenus struct {
	ID         int64      `gorm:"primary_key" json:"id"` //
	Pid        int64      `json:"pid"`                   //
	Name       string     `json:"name"`                  //
	Component  string     `json:"component"`             //组件地址
	Icon       string     `json:"icon"`                  //图片
	IsExt      int        `json:"isExt"`                 //是否外链
	IsIframe   int        `json:"isIframe"`              //是否嵌套
	Keepalive  int        `json:"keepalive"`             //是否缓存
	Show       int        `json:"show"`                  //是否显示
	Type       int        `json:"type"`                  //类型 1目录2菜单 3按钮
	SortNo     int        `json:"sortNo"`                //排序
	RouteName  string     `json:"routeName"`             //路由名称
	RoutePath  string     `json:"routePath"`             //路由地址
	Permission string     `json:"permission"`            //权限编码
	Status     int        `json:"status"`                //是否启用 0启用1禁用
	CreateTime int64      `json:"createTime"`            //创建时间
	List       []SysMenus `gorm:"-" json:"children"`
}

//添加菜单
func SysMenusAdd(data SysMenus) error {
	err := db.Db.Create(&data).Error
	return err
}

//获取列表
func SysMenusList() []SysMenus {
	var items []SysMenus
	db.Db.Model(&items).Order("sort_no asc").Find(&items)
	return items
}

//获取列表
func SysMenusGetList(where string, v ...interface{}) []SysMenus {
	var items []SysMenus
	db.Db.Table("sys_menus").Where(where, v...).Order("sort_no asc").Find(&items)
	if len(v) > 0 {
		return items
	}
	var list []SysMenus
	list = SysMenusDigui(items, 0, list)
	return list
}

//digui
func SysMenusDigui(items []SysMenus, pid int64, list []SysMenus) []SysMenus {
	var item []SysMenus
	for _, v := range items {
		if v.Pid == pid {
			var ls SysMenus
			utils.StAtoB(v, ls, &ls)
			ls.List = SysMenusDigui(items, v.ID, item)
			list = append(list, ls)
		}

	}
	return list
}

//删除列表
func SysMenusDel(data SysMenus) error {
	err := db.Db.Table("sys_menus").Where("id = ? or pid = ?", data.ID, data.ID).Delete(&data).Error
	return err
}

//编辑菜单
func SysMenusEdit(data SysMenus) error {
	err := db.Db.Save(&data).Error
	return err
}
