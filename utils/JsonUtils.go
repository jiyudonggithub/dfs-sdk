package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// StrToStruct 字符串转对象
func StrToStruct(str string, its interface{}) error {
	if len(str) < 1 {
		return nil
	}
	//json str 转struct
	err := json.Unmarshal([]byte(str), &its)
	if err != nil {
		//fmt.Println("对象转换错误", str, reflect.TypeOf(its).Kind())
	}
	return err
}

// ToJsonString 输出json字符串(处理特殊符>,&等HTML特殊字符被转义的)
func ToJsonString(obj interface{}) string {
	var bf bytes.Buffer
	enc := json.NewEncoder(&bf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(obj)
	if err != nil {
		fmt.Println("对象转字符串出错：", err)
		return ""
	} else {
		return bf.String()
	}
}
