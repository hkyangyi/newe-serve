package model

import (
	"newe-serve/common/db"
	"newe-serve/pkg/utils"
	"strconv"
	"strings"
)

type SysDepart struct {
	ID         string      `gorm:"primary_key" json:"id"` //uuid
	Pid        string      `json:"pid"`                   //父级ID
	Name       string      `json:"name"`                  //分组名称（机构名称）
	Code       string      `json:"code"`                  //分组编码
	Type       int         `json:"type"`                  //类型（1集团，2公司，3部门，4服务门店）
	Telephone  string      `json:"telephone"`             //联系电话
	Phone      string      `json:"phone"`                 //联系手机
	Address    string      `json:"address"`               //地址
	SortNo     int         `json:"sortNo"`                //排序
	CreateTime int64       `json:"createTime"`            //创建时间
	UpdateTime int64       `json:"updateTime"`            //更新时间
	List       []SysDepart `gorm:"-" json:"children"`
}

//添加
func (a *SysDepart) Add() error {
	a.ID = GetUUID()
	if len(a.Pid) == 0 {
		a.Code = a.GetCode("")

	} else {
		a.Code = a.GetCode(a.Pid)
	}
	err := db.Db.Create(a).Error
	return err
}

func (a *SysDepart) GetCode(pid string) string {
	var data SysDepart
	db.Db.Model(&SysDepart{}).Where("pid = ?", pid).Order("create_time desc").First(&data)
	if len(data.Code) > 0 {
		lsobj := strings.Split(data.Code[1:], "A")
		lsis := lsobj[len(lsobj)-1]
		isi, _ := strconv.Atoi(lsis)
		isi++
		if isi < 10 {
			lsis = "0" + strconv.Itoa(isi)
		} else {
			lsis = strconv.Itoa(isi)
		}
		lsobj[len(lsobj)-1] = lsis
		code := "A" + strings.Join(lsobj, "A")
		return code
	} else {
		if len(pid) == 0 {
			return "A00"
		} else {
			var pdb SysDepart
			db.Db.Model(&SysDepart{}).Where("id = ?", pid).Order("create_time desc").First(&pdb)
			return pdb.Code + "A00"
		}
	}
}

//编辑
func (a *SysDepart) Edit() error {
	err := db.Db.Save(a).Error
	return err
}

//删除
func (a *SysDepart) Del() error {
	err := db.Db.Table("sys_depart").Where("id = ? or pid = ?", a.ID, a.ID).Delete(&SysDepart{}).Error
	return err
}

//获取列表
func (a *SysDepart) GetList(where string, v ...interface{}) []SysDepart {
	var items []SysDepart
	db.Db.Table("sys_depart").Where(where, v...).Order("sort_no asc").Find(&items)
	if len(v) > 0 {
		return items
	}
	var list []SysDepart
	list = SysDepartDigui(items, "", list)
	return list
}

func SysDepartDigui(items []SysDepart, pid string, list []SysDepart) []SysDepart {
	var item []SysDepart
	for _, v := range items {
		if v.Pid == pid {
			var ls SysDepart
			utils.StAtoB(v, ls, &ls)
			ls.List = SysDepartDigui(items, v.ID, item)
			list = append(list, ls)
		}

	}
	return list
}
