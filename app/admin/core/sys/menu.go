package sys

import (
	"errors"
	"newe-serve/app/model"
	"newe-serve/pkg/utils"
	"strings"
	"time"
)

type SysMenus struct {
	ID         string `gorm:"primary_key" json:"id" form:"id"` //
	Pid        string `json:"pid"`                             //
	Name       string `json:"name" form:"name"`                //
	Component  string `json:"component"`                       //组件地址
	Icon       string `json:"icon"`                            //图片
	IsExt      int    `json:"isExt"`                           //是否外链
	IsIframe   int    `json:"isIframe"`                        //是否嵌套
	Keepalive  int    `json:"keepalive"`                       //是否缓存
	Show       int    `json:"show"`                            //是否显示
	Type       int    `json:"type"`                            //类型 1目录2菜单 3按钮
	SortNo     int    `json:"sortNo"`                          //排序
	RouteName  string `json:"routeName"`                       //路由名称
	RoutePath  string `json:"routePath"`                       //路由地址
	Permission string `json:"permission"`                      //权限编码
	Status     int    `json:"status" form:"status"`            //是否启用 0启用1禁用
	CreateTime int64  `json:"createTime"`                      //创建时间
}

//添加
func (a *SysMenus) Create() error {
	var data model.SysMenus
	a.CreateTime = time.Now().Unix()
	utils.StAtoB(*a, data, &data)
	err := model.SysMenusAdd(data)

	return err
}

//获取列表
func (a *SysMenus) GetList() []model.SysMenus {
	var params []interface{}
	var where []string
	if len(a.Name) > 0 {
		w := "name like ?"
		params = append(params, "%"+a.Name+"%")
		where = append(where, w)
	}

	if a.Status != 0 {
		w := "status = ?"
		params = append(params, a.Status)
		where = append(where, w)
	}
	var items []model.SysMenus

	ws := strings.Join(where, " AND ")

	items = model.SysMenusGetList(ws, params...)
	return items
}

//更新数据
func (a *SysMenus) Edit() error {
	if a.ID == "" {
		return errors.New("缺少参数ID")
	}
	var data model.SysMenus
	a.CreateTime = time.Now().Unix()
	utils.StAtoB(*a, data, &data)
	err := model.SysMenusEdit(data)

	return err
}

//删除数据
func (a *SysMenus) Del() error {
	if a.ID == "" {
		return errors.New("缺少参数ID")
	}
	var data model.SysMenus
	data.ID = a.ID
	err := model.SysMenusDel(data)

	return err
}

type Mate struct {
	Title           string `json:"title"`
	IgnoreKeepAlive bool   `json:"ignoreKeepAlive"` // 是否忽略KeepAlive缓存
	Affix           bool   `json:"affix"`           // 是否固定标签
	Icon            string `json:"icon"`            // 图标，也是菜单图标
	FrameSrc        string `json:"frameSrc"`        // 内嵌iframe的地址
	HideMenu        bool   `json:"hideMenu"`        // 当前路由不再菜单显示
	OrderNo         int    `json:"orderNo"`         // 菜单排序，只对第一级有效
}

type MenuItem struct {
	Id        string     `json:"id"`
	Pid       string     `json:"pid"`
	Path      string     `json:"path"`
	Name      string     `json:"name"`
	Component string     `json:"component"`
	Redirect  string     `json:"redirect"`
	Mate      Mate       `json:"meta"`
	Children  []MenuItem `json:"children"`
}

func GetMenuList(departId string) []MenuItem {
	menus := model.SysMenusListGetByDepart(departId)
	var items []MenuItem
	for _, v := range menus {
		if v.Type == 3 {
			break
		}
		if v.Status == 0 {
			break
		}
		var item MenuItem
		item.Id = v.ID
		item.Pid = v.Pid
		item.Path = v.RoutePath
		item.Name = v.RouteName
		if v.Type == 1 {
			item.Component = "LAYOUT"
		} else {
			item.Component = v.Component
		}

		var mate Mate
		mate.Title = v.Name
		mate.Icon = v.Icon
		if v.Keepalive == 1 {
			mate.IgnoreKeepAlive = false
		} else {
			mate.IgnoreKeepAlive = true
		}

		if v.Show == 1 {
			mate.HideMenu = false
		} else {
			mate.HideMenu = true
		}
		mate.OrderNo = v.SortNo

		if v.IsIframe == 1 {
			item.Component = "/sys/iframe/FrameBlank"
			mate.FrameSrc = v.Component
		}

		item.Mate = mate

		items = append(items, item)
	}
	var bitems []MenuItem
	//递归
	bitems = menuDigui(items, "", bitems)
	// for i:=0;i<len(bitems);i++{
	// 	if len(bitems[i].Children) >0{
	// 		bitems[i].Redirect = "/"+
	// 	}
	// }
	return bitems
}

func menuDigui(items []MenuItem, pid string, list []MenuItem) []MenuItem {
	var item []MenuItem
	for _, v := range items {
		if v.Pid == pid {
			var ls MenuItem
			utils.StAtoB(v, ls, &ls)
			ls.Children = menuDigui(items, v.Id, item)
			list = append(list, ls)
		}
	}
	return list
}
