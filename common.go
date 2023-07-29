// @Title
// @Author  zls  2023/7/28 21:58
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 处理异常
func HandleError(err error) {
	if err != nil {
		panic("程序发生异常: " + err.Error())
	}
}

func CloseRspBody(resp *http.Response) {
	err := resp.Body.Close()
	HandleError(err)
}

func MapToJson(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}

func GetRespBodyStr(r io.Reader) string {
	body, err := io.ReadAll(r) // 读取响应 body, 返回为 []byte
	HandleError(err)
	var str bytes.Buffer
	_ = json.Indent(&str, body, "", "    ")
	return str.String()
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatalln("关闭文件失败", err.Error())
	}
}
