package sys

import (
	"errors"
	"newe-serve/app/model"
	"newe-serve/pkg/utils"
	"strings"
	"time"
)

type SysMember struct {
	ID         string `json:"id" form:"id"`             //
	DepartId   string `json:"departId"`                 //组织结构ID
	UID        string `json:"uid"`                      //会员ID
	Username   string `json:"username" form:"username"` //登陆账号
	Password   string `json:"password"`                 //密码
	Nickname   string `json:"nickname" form:"nickname"` //昵称
	Realname   string `json:"realname" form:"realname"` //真实姓名
	Headimgurl string `json:"headimgurl"`               //头像
	Mp         string `json:"mp" form:"mp"`             //手机号
	Idcard     string `json:"idcard"`                   //身份证号码
	Sex        int    `json:"sex"`                      //性别 1男2女
	Status     int    `json:"status"`                   //1正常，2禁用
	OrgCode    string `json:"orgCode"`                  //组织结构编码
	CreateTime int64  `json:"createTime"`               //创建时间
	UpdateTime int64  `json:"updateTime"`               //更新时间
	Files      string `json:"files"`                    //附件
	utils.PageList
}

func (a *SysMember) GetList() utils.PageList {
	var data model.SysMember
	// err:=utils.StAtoB(*a, data, &data)
	// if err!=nil{
	// 	return err
	// }
	// data.CreateTime = time.Now().Unix()
	// data.UpdateTime = time.Now().Unix()
	var wheremap []string
	var params []interface{}
	if len(a.Username) > 0 {
		wheremap = append(wheremap, " username = ?")
		params = append(params, a.Username)
	}

	if len(a.Nickname) > 0 {
		wheremap = append(wheremap, "nickname like ?")
		params = append(params, "%"+a.Nickname+"%")
	}

	if len(a.Mp) > 0 {
		wheremap = append(wheremap, " mp = ?")
		params = append(params, "%"+a.Mp+"%")
	}

	if len(a.Realname) > 0 {
		wheremap = append(wheremap, " realname = ? ")
		params = append(params, a.Realname)
	}

	where := strings.Join(wheremap, " AND ")
	page := utils.PageList{
		Page:     a.Page,
		PageSize: a.PageSize,
	}
	items := data.GetList(page, where, params...)
	return items

}

func (a *SysMember) Create() error {
	var data model.SysMember
	err := utils.StAtoB(*a, data, &data)
	if err != nil {
		return err
	}
	data.Password = utils.EncodeMD5(a.Password)
	data.CreateTime = time.Now().Unix()
	data.UpdateTime = time.Now().Unix()
	err = data.Add()
	return err
}

func (a *SysMember) Update() error {
	if len(a.ID) != 32 {
		return errors.New("参数ID错误")
	}
	var data model.SysMember
	err := utils.StAtoB(*a, data, &data)
	if err != nil {
		return err
	}
	if len(a.Password) > 0 {
		data.Password = utils.EncodeMD5(a.Password)
	}
	data.UpdateTime = time.Now().Unix()
	err = data.Edit()
	return err
}

func (a *SysMember) Del() error {
	if len(a.ID) != 32 {
		return errors.New("参数ID错误")
	}
	var data = model.SysMember{
		ID: a.ID,
	}

	err := data.Del()
	return err
}
