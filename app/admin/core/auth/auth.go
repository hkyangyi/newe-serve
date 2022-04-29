package auth

import (
	"errors"
	"newe-serve/app/model"
	"newe-serve/common/redis"
	"newe-serve/pkg/utils"
	"strconv"
)

type AuthFrom struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/**
* 用户登陆
 */
func (a *AuthFrom) Login() (map[string]interface{}, error) {
	if len(a.Username) < 5 {
		return nil, errors.New("请输入正确的账号")
	}

	rediskey := "AUTH_ERROR_COUNT_" + a.Username
	var LoginErrCount int
	if redis.Exists(rediskey) {
		redis.Get(rediskey, &LoginErrCount)
	} else {
		LoginErrCount = 0
	}
	if LoginErrCount >= 5 {
		return nil, errors.New("您今日输入错误次数过多，请24小时后再试")
	}

	var data = make(map[string]interface{})
	//查询用户信息
	merdb, err := model.FindMemberByUsername(a.Username)
	if err != nil {
		return nil, err
	}

	if len(merdb.ID) != 32 {
		return nil, errors.New("账号或密码错误")
	}

	md5ps := utils.EncodeMD5(a.Password)
	if md5ps != merdb.Password {
		LoginErrCount++
		redis.Set(rediskey, LoginErrCount, 60*60*24)
		return nil, errors.New("密码错误,您还有" + strconv.Itoa(5-LoginErrCount) + "次机会")
	}

	merdb.Password = ""
	//登陆成功
	data["usdb"] = merdb
	authkey := "AUTH_" + model.GetUUID()
	token, _ := utils.SetToken(authkey)
	redis.Set(authkey, merdb, 60*60)
	data["token"] = token

	return data, nil
}
