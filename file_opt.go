// @Title
// @Author  zls  2023/7/29 10:50
package main

import (
	"amazon_stream/common"
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
	"time"
)

// 将字符串写入文件
func writeStr2File4AccessToken(token string) error {
	if token == "" {
		return errors.New("参数不能为空")
	}
	// 读取文件
	file, err := os.OpenFile("./token.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer common.CloseFile(file)

	// 获取当前时间
	timeStr := time.Now().Format(common.TimeTemplate)

	// 开始写入文件
	_, err = file.WriteString("time:" + timeStr)
	if err != nil {
		return err
	}
	_, err = file.WriteString("token:" + token)
	if err != nil {
		return err
	}
	return nil
}

// 从文件中获取token
func getTokenFromFile() (string, error) {
	// 读取文件
	file, err := os.OpenFile("./token.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer common.CloseFile(file)

	var token string
	var timeStr string
	reader := bufio.NewReader(file)
	// 按行读取  ReaderString
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if strings.HasPrefix(line, "token") {
			token = strings.ReplaceAll(line, "token:", "")
		}
		if strings.HasPrefix(line, "time") {
			timeStr = strings.ReplaceAll(line, "time:", "")
		}
	}
	if token == "" {
		return "", nil
	}
	// 如果时间超过1小时, 也返回""
	oldTime, err := time.Parse(common.TimeTemplate, timeStr)
	if err != nil {
		return "", err
	}
	sub := time.Now().Sub(oldTime)
	if sub.Seconds() >= 3600 {
		return "", nil
	}
	return token, nil
}
