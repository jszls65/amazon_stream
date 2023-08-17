// @Title
// @Author  zls  2023/7/28 21:58
package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/objx"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func GetRandomToken(maxLen int) string {
	uuidStr := uuid.New().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	if len(uuidStr) > maxLen {
		return uuidStr[:maxLen]
	}
	return uuidStr

}

func JsonToMap(jsonStr string) (objx.Map, error) {
	fromJSON, err := objx.FromJSON(jsonStr)
	if err != nil {
		return nil, err
	}
	return fromJSON, nil
}

// 结构体转json字符串
func ToJsonStr(st interface{}) string {
	marshal, err := json.Marshal(st)
	if err != nil {
		log.Fatalln("结构体转json字符串 异常,", err)
	}
	return string(marshal)

}

func GetId(pre string) string {
	return pre + strconv.Itoa(int(time.Now().Unix()))
}
