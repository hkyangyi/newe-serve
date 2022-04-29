package sys

import (
	"errors"
	"newe-serve/app/model"
	"newe-serve/pkg/utils"
	"strings"
	"time"
)

type SysDepart struct {
	ID         string      `gorm:"primary_key" json:"id" form:"id"` //uuid
	Pid        string      `json:"pid"`                             //父级ID
	Name       string      `json:"name"`                            //分组名称（机构名称）
	Code       string      `json:"code"`                            //分组编码
	Type       int         `json:"type"`                            //类型（1集团，2公司，3部门，4服务门店）
	Telephone  string      `json:"telephone"`                       //联系电话
	Phone      string      `json:"phone"`                           //联系手机
	Address    string      `json:"address"`                         //地址
	SortNo     int         `json:"sortNo"`                          //排序
	CreateTime int64       `json:"createTime"`                      //创建时间
	UpdateTime int64       `json:"updateTime"`                      //更新时间
	List       []SysDepart `gorm:"-" json:"children"`
	Menus      string      `json:"menus"`
}

//添加
func (a *SysDepart) Add() error {
	var data model.SysDepart
	a.CreateTime = time.Now().Unix()
	a.UpdateTime = time.Now().Unix()
	utils.StAtoB(*a, data, &data)
	err := data.Add()
	return err
}

//获取列表
func (a *SysDepart) GetList() []model.SysDepart {
	var params []interface{}
	var where []string
	if len(a.Name) > 0 {
		w := "name like ?"
		params = append(params, "%"+a.Name+"%")
		where = append(where, w)
	}

	var data model.SysDepart

	ws := strings.Join(where, " AND ")

	items := data.GetList(ws, params...)
	return items
}

//更新数据
func (a *SysDepart) Edit() error {
	if a.ID == "" || len(a.ID) < 32 {
		return errors.New("缺少参数ID")
	}
	var data model.SysDepart
	a.UpdateTime = time.Now().Unix()
	utils.StAtoB(*a, data, &data)
	err := data.Edit()
	return err
}

//删除数据
func (a *SysDepart) Del() error {
	if a.ID == "" || len(a.ID) < 32 {
		return errors.New("缺少参数ID")
	}
	var data model.SysDepart
	data.ID = a.ID
	err := data.Del()

	return err
}

//获取权限列表
func (a *SysDepart) GetRules() ([]string, error) {
	if a.ID == "" || len(a.ID) < 32 {
		return []string{}, errors.New("缺少参数ID")
	}
	var data = model.SysDepart{
		ID: a.ID,
	}
	items := data.GetRules()
	var ids []string
	for _, v := range items {
		ids = append(ids, v.MenuId)
	}

	return ids, nil
}

func (a *SysDepart) SaveRules() error {
	if a.ID == "" || len(a.ID) < 32 {
		return errors.New("缺少参数ID")
	}

	items, err := a.GetRules()
	if err != nil {
		return err
	}

	farr := strings.Split(a.Menus, ",")

	var delarr, addarr []string
	for i := 0; i < len(items); i++ {
		b := utils.InArrar(items[i], farr)
		if !b {
			delarr = append(delarr, items[i])
		}
	}

	for _, v := range farr {
		if b := utils.InArrar(v, items); !b {
			addarr = append(addarr, v)
		}
	}

	var data = model.SysDepart{
		ID: a.ID,
	}
	if len(addarr) > 0 {
		err = data.AddRules(addarr)
		if err != nil {
			return err
		}
	}

	if len(delarr) > 0 {
		err = data.DelRules(delarr)
	}
	return err
}
