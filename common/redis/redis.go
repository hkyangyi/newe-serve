package redis

import (
	"encoding/json"
	"newe-serve/common/setting"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool
var RedisStatus bool

func Setup() error {
	RedisStatus = false
	RedisConn = &redis.Pool{
		MaxIdle:     setting.Rediscfg.MaxIdle,
		MaxActive:   setting.Rediscfg.MaxActive,
		IdleTimeout: time.Duration(int64(setting.Rediscfg.IdleTimeout)),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.Rediscfg.Host)
			if err != nil {
				return nil, err
			}
			if setting.Rediscfg.Password != "" {
				if _, err := c.Do("AUTH", setting.Rediscfg.Password); err != nil {
					c.Close()
					RedisStatus = false
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	RedisStatus = true
	return nil
}

//设置缓存
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

//检测缓存中是否有KEY
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

//获取数据
func Get(key string, v interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return err
	}

	err = json.Unmarshal(reply, &v)
	return err
}

//删除
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
