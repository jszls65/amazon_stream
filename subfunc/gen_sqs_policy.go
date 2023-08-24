// @Title
// @Author  zls  2023/8/8 14:07
package subfunc

import (
	"amazon_stream/common"
	"github.com/stretchr/objx"
	"strconv"
	"strings"
)

var queue string // sqs arn
var queueUrl string
var queueRootArn string
var topic string   // topic arn
var configId int32 // topic arn

// GenSqsPolicy 生成sqs的访问策略
func GenSqsPolicy(shopName string) map[string]interface{} {
	shopData := common.GetShopDataMap(shopName)
	if shopData.SqsArn == "" || shopData.IamRoot == "" || shopData.TopicArn == "" || shopData.SqsURL == "" {
		return objx.Map{
			"错误":       "表中缺少数据",
			"SqsArn":   shopData.SqsArn,
			"SqsUrl":   shopData.SqsURL,
			"IamRoot":  shopData.IamRoot,
			"TopicArn": shopData.TopicArn,
		}
	}
	queue = shopData.SqsArn
	queueUrl = shopData.SqsURL
	queueRootArn = shopData.IamRoot
	topic = shopData.TopicArn
	configId = shopData.ConfigID

	mapList := make([]objx.Map, 0)
	// queueRoot
	mapList = append(mapList, replaceCon(queueRoot))
	// sns
	mapList = append(mapList, replaceCon(sns))
	// amazonOpt
	//mapList = append(mapList, replaceCon(amazonOpt))

	// 公共
	mapList = append(mapList, replaceCon(gonggong))
	// spTraffic
	mapList = append(mapList, replaceCon(spTraffic))
	// spConversion
	mapList = append(mapList, replaceCon(spConversion))
	// sdTraffic
	mapList = append(mapList, replaceCon(sdTraffic))
	// sdConversion
	mapList = append(mapList, replaceCon(sdConversion))
	// budgetUsage
	mapList = append(mapList, replaceCon(budgetUsage))
	// sponsoredAdsCampaignDiagnosticsRecommendations
	mapList = append(mapList, replaceCon(sponsoredAdsCampaignDiagnosticsRecommendations))
	// campaigns
	mapList = append(mapList, replaceCon(campaigns))
	// adgroups
	mapList = append(mapList, replaceCon(adgroups))
	// ads
	mapList = append(mapList, replaceCon(ads))
	// targets
	mapList = append(mapList, replaceCon(targets))

	// 订单小时级别数据配置
	//mapList = append(mapList, replaceCon(orderOpt))

	// 主内容
	mainMap := objx.MustFromJSON(mainStr)
	mainMap["Statement"] = mapList
	mainMap["Id"] = common.GetId("main")
	// 角色策略
	rolePolicy = strings.ReplaceAll(rolePolicy, "{queue}", queue)
	rolePolicy = strings.ReplaceAll(rolePolicy, "{topic}", topic)
	// nacos配置
	nacosConfig = strings.ReplaceAll(nacosConfig, "{config-id}", strconv.Itoa(int(configId)))
	nacosConfig = strings.ReplaceAll(nacosConfig, "{queue-url}", queueUrl)
	// 返回值
	resultMap := objx.Map{
		"sqs策略":   mainMap,
		"角色策略":    objx.MustFromJSON(rolePolicy),
		"nacos配置": objx.MustFromJSON(nacosConfig),
	}
	return resultMap
}

// 替换内容
func replaceCon(str string) objx.Map {
	str = strings.ReplaceAll(str, "{id}", common.GetId("Item"))
	str = strings.ReplaceAll(str, "{queue}", queue)
	str = strings.ReplaceAll(str, "{topic}", topic)
	str = strings.ReplaceAll(str, "{queue-root}", queueRootArn)
	return objx.MustFromJSON(str)
}

var mainStr = `{
  "Version": "2012-10-17",
  "Id": "{id}"
}`

// 公共
var gonggong = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::926844853897:role/ReviewerRole"
      },
      "Action": "SQS:GetQueueAttributes",
      "Resource": "*"
    }`

var spTraffic = `
	{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:906013806264:*"
        }
      }
    }`

var spConversion = ` {
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:802324068763:*"
        }
      }
    }`

var sdTraffic = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:370941301809:*"
        }
      }
    }`

var sdConversion = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:877712924581:*"
        }
      }
    }`

var budgetUsage = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:055588217351:*"
        }
      }
    }`

var sponsoredAdsCampaignDiagnosticsRecommendations = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:084590724871:*"
        }
      }
    }`

var campaigns = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "*",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:570159413969:*"
        }
      }
    }`

var adgroups = ` {
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "*",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:118846437111:*"
        }
      }
    }`

var ads = `{
	  "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "*",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:305370293182:*"
        }
      }
    }`

var targets = `{
	  "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "*",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:644124924521:*"
        }
      }
    }`

var queueRoot = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "AWS": "{queue-root}" 
      },
      "Action": "SQS:*",
      "Resource": "{queue}"
    }`
var sns = ` {
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "sqs:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "{topic}"
        }
      }
    }`

// 亚马逊对sqs的操作
var amazonOpt = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::437568002678:root"
      },
      "Action": [
        "sqs:GetQueueAttributes",
        "sqs:SendMessage"
      ],
      "Resource": "{queue}"
    }`

// 订单小时级别数据
var orderOpt = `{
      "Sid": "{id}",
      "Action": [
        "sqs:GetQueueAttributes",
        "sqs:SendMessage"
      ],
      "Effect": "Allow",
      "Resource": "{queue}",
      "Principal": {
        "AWS": [
          "437568002678"
        ]
      }
    }`

var rolePolicy = `{"list":[{
    "Effect": "Allow",
    "Action": "sns:Publish",
    "Resource": "{topic}"
},
{
    "Effect": "Allow",
    "Action": [
        "sqs:ReceiveMessage",
        "sqs:DeleteMessage"
    ],
    "Resource": [
        "{queue}"
    ]
}]}`

var nacosConfig = `{"nacos配置id={config-id}":"sqsStandPath: {queue-url}"}`
