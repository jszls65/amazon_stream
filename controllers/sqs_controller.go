// 操作sqs controller
package controllers

import (
	"amazon_stream/conf"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/gin-gonic/gin"
)

type SqsController struct{}

// 拉取消息
func (sc SqsController) PullMsg(context *gin.Context) {
	// 校验参数
	sqsName, ok := context.GetQuery("sqsName")
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": "参数sqsName不能为空"})
		return
	}
	selection := conf.GetConfigByName(sqsName)
	region := selection.Key("region").String()
	if "" == region {
		context.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": "sqsName不存在"})
		return
	}
	accessKeyID := selection.Key("accessKeyID").String()
	secretAccessKey := selection.Key("secretAccessKey").String()
	queueUrl := selection.Key("queueURL").String()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		fmt.Println("创建session异常:", err.Error())
		return
	}

	sqsClient := sqs.New(sess)

	params := &sqs.ReceiveMessageInput{
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: aws.Int64(1),
		WaitTimeSeconds:     aws.Int64(20),
	}
	go sc.doReceiveMessage(sqsClient, params)

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "任务已提交, 请稍后",
	})
}

func (sc SqsController) doReceiveMessage(sqsClient *sqs.SQS, params *sqs.ReceiveMessageInput) {
	for {
		resp, err := sqsClient.ReceiveMessage(params)
		if err != nil {
			fmt.Println("sqs接收消息异常:", err)
			continue
		}

		for _, msg := range resp.Messages {
			fmt.Println("Received message:", *msg.Body)

			deleteParams := &sqs.DeleteMessageInput{
				QueueUrl:      params.QueueUrl,
				ReceiptHandle: msg.ReceiptHandle,
			}
			// 手动删除
			_, err := sqsClient.DeleteMessage(deleteParams)
			if err != nil {
				fmt.Println("删除消息失败:", err, ", 消息内容:", *msg.Body)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
