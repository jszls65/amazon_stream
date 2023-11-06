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

var mainStr string

// 公共
var gonggong string
var spTraffic string
var spConversion string
var sdTraffic string
var sdConversion string
var budgetUsage string
var sbTraffic string
var sbConversion string
var sbClickstream string
var sponsoredAdsCampaignDiagnosticsRecommendations string
var campaigns string
var adgroups string
var ads string
var targets string
var queueRoot string
var sns string

// 亚马逊对sqs的操作
var amazonOpt string

// 订单小时级别数据
var orderOpt string
var rolePolicy string
var nacosConfig string

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

	initOriginJson()

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
	mapList = append(mapList, replaceCon(amazonOpt))

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
  // sbTaffice
  mapList = append(mapList, replaceCon(sbTraffic))
  // sbConversion
  mapList = append(mapList, replaceCon(sbConversion))
  // sbClickstream
  mapList = append(mapList, replaceCon(sbClickstream))

	// 订单小时级别数据配置
	// mapList = append(mapList, replaceCon(orderOpt))

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
		"a_nacos_config": objx.MustFromJSON(nacosConfig),
		"b_role_policy":  objx.MustFromJSON(rolePolicy),
		"c_sqs_policy":   mainMap,
	}
	return resultMap
}

func initOriginJson() {

	mainStr = `{
  "Version": "2012-10-17",
  "Id": "{id}"
}`

	// 公共
	gonggong = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::926844853897:role/ReviewerRole"
      },
      "Action": "SQS:GetQueueAttributes",
      "Resource": "*"
    }`

	spTraffic = `
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

	spConversion = ` {
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

	sdTraffic = `{
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

	sdConversion = `{
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

    sbTraffic = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:709476672186:*"
        }
      }
    }`

    sbConversion = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:154357381721:*"
        }
      }
    }`

    sbClickstream = ` {
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "Service": "sns.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "{queue}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "arn:aws:sns:us-east-1:091028706140:*"
        }
      }
    }`
	budgetUsage = `{
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

	sponsoredAdsCampaignDiagnosticsRecommendations = `{
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

	campaigns = `{
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

	adgroups = ` {
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

	ads = `{
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

	targets = `{
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

	queueRoot = `{
      "Sid": "{id}",
      "Effect": "Allow",
      "Principal": {
        "AWS": "{queue-root}" 
      },
      "Action": "SQS:*",
      "Resource": "{queue}"
    }`
	sns = ` {
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

	// 亚马逊对sqs的操作, 订单小时级别数据配置
	amazonOpt = `{
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
	orderOpt = `{
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

	rolePolicy = `{"list":[{
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

	nacosConfig = `{"nacos配置 id: {config-id}":"sqsStandPath: {queue-url}"}`

}

// 替换内容
func replaceCon(str string) objx.Map {
	str = strings.ReplaceAll(str, "{id}", common.GetId("Item"))
	str = strings.ReplaceAll(str, "{queue}", queue)
	str = strings.ReplaceAll(str, "{topic}", topic)
	str = strings.ReplaceAll(str, "{queue-root}", queueRootArn)
	return objx.MustFromJSON(str)
}
