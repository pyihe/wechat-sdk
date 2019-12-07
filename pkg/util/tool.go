package util

import (
	"os"
	"reflect"
)

//判断array中是否有元素value
func HaveInArray(array interface{}, value interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

//将嵌套的map[string]interface全部转换成一层
func Interface2Map(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range data.(map[string]interface{}) {
		switch v := v.(type) {
		case map[string]interface{}:
			for i, u := range v {
				result[i] = u
			}
		default:
			result[k] = v
		}
	}
	return result
}

//判断目录是否存在，如果不存在，则新建一个目录
func MakeNewPath(targetPath string) error {
	if _, err := os.Stat(targetPath); err != nil {
		if !os.IsExist(err) {
			//创建目录
			if mErr := os.MkdirAll(targetPath, os.ModePerm); mErr != nil {
				return mErr
			}
		}
	}
	return nil
}
