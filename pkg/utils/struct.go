package utils

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func StAtoB(a, b, c interface{}) error {
	//第一步,先将结构体转化为map方便后续遍历
	amap := Struct2Map(a)
	bmap := Struct2Map(b)
	for k1, v1 := range amap {
		if _, ok := bmap[k1]; ok {
			typea := reflect.TypeOf(v1)
			typeb := reflect.TypeOf(bmap[k1])
			if typea == typeb {
				bmap[k1] = v1
			}
		}
		//bmap[k1] = v1
		// for k2, v2 := range bmap {
		// 	typea := reflect.TypeOf(v1)
		// 	typeb := reflect.TypeOf(v2)
		// 	if k1 == k2 && typea == typeb {
		// 		bmap[k2] = v1
		// 	}
		// }
	}

	if err := mapstructure.Decode(bmap, c); err != nil {
		fmt.Println("the err when unmarshal mapstructure is:", err)
		return err
	}
	return nil

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

// func Digui(items []interface{}, list interface{}, key string, val interface{}) error {
// 	var item []interface{}
// 	items = acdigui(items, key, val, item)

// 	if err := mapstructure.Decode(items, list); err != nil {
// 		fmt.Println("the err when unmarshal mapstructure is:", err)
// 		return err
// 	}

// 	return nil

// }

// func acdigui(items []interface{}, key string, val interface{}, list []interface{}) []interface{} {
// 	var item []interface{}
// 	for _, v := range items {
// 		mp := Struct2Map(v)
// 		if mp[key] == val {
// 			var ls map[string]interface{}
// 			for k, v := range mp {
// 				ls[k] = v
// 			}

// 			ls["list"] = acdigui(items, key, ls["id"], item)
// 			var b interface{}
// 			mapstructure.Decode(ls, b)
// 			list = append(list, ls)
// 		}

// 	}
// 	return list
// }
