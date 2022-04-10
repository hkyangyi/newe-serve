package setting

import (
	"fmt"
	"log"
	"newe-serve/pkg/file"
	"reflect"
	"strconv"

	"gopkg.in/ini.v1"
)

type Httpcgf struct {
	RunMode         string //运行模式debug or release
	HttpPort        int    //http服务端口
	WsPort          int    //ws服务端口
	ReadTimeout     int    //读取时间
	WriteTimeout    int    //写入时间
	ServeUrl        string //服务地址
	RuntimeRootPath string //日志存储目录
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type Images struct {
	//图片访问URL
	ImagePrefixUrl string
	//图片上传地址
	ImageSavePath string
	//#图片最大
	ImageMaxSize int
	//#图片格式
	ImageAllowExts string
}

var SqlDb = &Database{}
var SYS = &Httpcgf{}
var Rediscfg = &Redis{}
var Imgcfg = &Images{}

func SetUp() {

	cfg, err := ini.Load("config/serve.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/serve.ini': %v", err)

	}

	err = cfg.Section("sys").MapTo(SYS)
	if err != nil {
		log.Fatalf("Cfg.MapTo Sys err: %v", err)
	}

	err = cfg.Section("database").MapTo(SqlDb)
	if err != nil {
		log.Fatalf("Cfg.MapTo Appdb err: %v", err)
	}

	err = cfg.Section("images").MapTo(Imgcfg)
	if err != nil {
		log.Fatalf("Cfg.MapTo images err: %v", err)
	}

	err = cfg.Section("redis").MapTo(Rediscfg)
	if err != nil {
		log.Fatalf("Cfg.MapTo redis err: %v", err)
	}
	// Rediscfg.IdleTimeout = Rediscfg.IdleTimeout * time.Second

}

func init() {
	//检查文件是否存在
	b := file.CheckNotExist("config/serve.ini")
	fmt.Println(b)
	if b {
		fmt.Println("开始创建文件")
		err := file.MkAll("config", "serve.ini")
		if err != nil {
			log.Fatalf("创建配置文件失败")
		}
		cfg, err := ini.Load("config/serve.ini")

		Http := Httpcgf{
			RunMode:         "debug",             //运行模式debug or release
			HttpPort:        80,                  //http服务端口
			WsPort:          808,                 //ws服务端口
			ReadTimeout:     60,                  //读取时间
			WriteTimeout:    60,                  //写入时间
			ServeUrl:        "http://localhost/", //服务地址
			RuntimeRootPath: "runtime",           //日志存储目录
		}

		dbm := Struct2Map(Http)
		sys, _ := cfg.NewSection("sys")
		WriteTo(sys, dbm)

		var mysql = Database{
			Type:        "mysql",          //数据链接类型
			User:        "root",           //用户名
			Password:    "newe123",        //newe123
			Host:        "127.0.0.1:3306", //链接地址
			Name:        "newe",           //数据库名
			TablePrefix: "",               //表前缀
		}
		sqlm := Struct2Map(mysql)
		sql, _ := cfg.NewSection("database")
		WriteTo(sql, sqlm)

		var redis = Redis{
			Host:        "127.0.0.1:6379",
			Password:    "",
			MaxIdle:     2,   //最大空闲连接数
			MaxActive:   10,  // #在给定时间内，允许分配的最大连接数（当为零时，没有限制）
			IdleTimeout: 200, // #在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
		}

		rdsm := Struct2Map(redis)
		rdss, _ := cfg.NewSection("redis")
		WriteTo(rdss, rdsm)

		var image = Images{
			//图片访问URL
			ImagePrefixUrl: "",
			//图片上传地址
			ImageSavePath: "",
			//#图片最大单位B
			ImageMaxSize: 2097152,
			//#图片格式
			ImageAllowExts: ".jpg,.jpeg,.png",
		}

		imgm := Struct2Map(image)
		imgs, _ := cfg.NewSection("image")
		WriteTo(imgs, imgm)

		// sys, _ := cfg.NewSection("sys")
		// sys.StrictMapTo(db.Http)
		fmt.Println("初始化配置文件结束")
		cfg.SaveTo("config/serve.ini")
	}
	return
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func WriteTo(s *ini.Section, obj map[string]interface{}) {
	for k, v := range obj {
		var val string

		if reflect.TypeOf(v).Name() == "int" {
			va := v.(int)
			val = strconv.Itoa(va)
		} else {
			val = v.(string)
		}
		s.NewKey(k, val)
	}
}
