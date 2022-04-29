package model

import (
	"newe-serve/common/db"
	"newe-serve/pkg/utils"
)

type SysMember struct {
	ID         string `gorm:"primary_key" json:"id"` //
	DepartId   string `json:"departId"`              //组织结构ID
	UID        string `json:"uid"`                   //会员ID
	Username   string `json:"username"`              //登陆账号
	Password   string `json:"password"`              //密码
	Nickname   string `json:"nickname"`              //昵称
	Realname   string `json:"realname"`              //真实姓名
	Headimgurl string `json:"headimgurl"`            //头像
	Mp         string `json:"mp"`                    //手机号
	Idcard     string `json:"idcard"`                //身份证号码
	Sex        int    `json:"sex"`                   //性别 1男2女
	Status     int    `json:"status"`                //1正常，2禁用
	OrgCode    string `json:"orgCode"`               //组织结构编码
	CreateTime int64  `json:"createTime"`            //创建时间
	UpdateTime int64  `json:"updateTime"`            //更新时间
	Files      string `json:"files"`                 //附件
}

//添加
func (a *SysMember) Add() error {
	a.ID = GetUUID()
	err := db.Db.Create(a).Error
	return err
}

//编辑
func (a *SysMember) Edit() error {
	err := db.Db.Model(a).Updates(a).Error
	return err
}

//获取列表
func (a *SysMember) GetList(page utils.PageList, where string, v ...interface{}) utils.PageList {
	var items []SysMember
	db.Db.Model(&SysMember{}).Where(where, v...).Count(&page.Total).Order("create_time desc").Offset(page.GetOffice()).Limit(page.PageSize).Find(&items)
	page.List = items
	return page
}

//删除
func (a *SysMember) Del() error {
	err := db.Db.Model(a).Delete(a).Error
	return err
}

//用户名查询
func FindMemberByUsername(username string) (SysMember, error) {
	var data SysMember
	err := db.Db.Model(&data).Where("username = ?", username).First(&data).Error
	return data, err
}
