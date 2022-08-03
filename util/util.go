package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// Getwd 获得项目的根路径
func Getwd() string {
	dir, _ := os.Getwd()
	return dir
}

// PrintJson 打印Json
func PrintJson(args interface{}) string {
	b, err := json.Marshal(args)
	if err != nil {
		return fmt.Sprintf("%+v", args)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", args)
	}
	return out.String()
}
