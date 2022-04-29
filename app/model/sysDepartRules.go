package model

import (
	"newe-serve/common/db"
	"time"
)

type SysDepartRules struct {
	ID         string `gorm:"primary_key" json:"id"` //
	DepartId   string `json:"departId"`              //组织结构ID
	OrgCode    string `json:"orgCode"`               //组织结构编码
	MenuId     string `json:"menuId"`                //菜单ID
	CreateTime int64  `json:"createTime"`            //
}

//根据部门获取已有权限
func (a *SysDepart) GetRules() []SysDepartRules {
	var items []SysDepartRules
	db.Db.Model(&items).Where("depart_id = ?", a.ID).Find(&items)
	return items

}

func (a *SysDepart) DelRules(ids []string) error {
	err := db.Db.Table("sys_depart_rules").Where("menu_id IN ?", ids).Delete(&SysDepartRules{}).Error
	return err
}

func (a *SysDepart) AddRules(ids []string) error {
	var items []SysDepartRules
	for _, v := range ids {
		item := SysDepartRules{
			ID:         GetUUID(),
			OrgCode:    a.Code,
			DepartId:   a.ID,
			MenuId:     v,
			CreateTime: time.Now().Unix(),
		}
		items = append(items, item)
	}
	err := db.Db.Create(&items).Error
	return err
}
