// @Title
// @Author  zls  2023/8/8 14:07
package subfunc

import (
	"amazon_stream/common"
	"fmt"
	"github.com/stretchr/objx"
	"strconv"
	"strings"
	"time"
)

var queue string // sqs arn
var queueRootArn string
var topic string // topic arn

// GenSqsPolicy 生成sqs的访问策略
func GenSqsPolicy(shopName string) map[string]interface{} {
	shopData := common.GetShopDataMap(shopName)
	if shopData.SqsArn == "" || shopData.IamRoot == "" || shopData.TopicArn == "" {
		return objx.Map{
			"错误":       "表中缺少数据",
			"SqsArn":   shopData.SqsArn,
			"IamRoot":  shopData.IamRoot,
			"TopicArn": shopData.TopicArn,
		}
	}
	queue = shopData.SqsArn
	queueRootArn = shopData.IamRoot
	topic = shopData.TopicArn

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

	// 主内容
	mainMap := objx.MustFromJSON(mainStr)
	mainMap["Statement"] = mapList
	mainMap["Id"] = getId("main")
	str := common.ToJsonStr(mainMap)
	fmt.Println("sqs访问策略")
	fmt.Println(str)
	rolePolicy = strings.ReplaceAll(rolePolicy, "{queue}", queue)
	rolePolicy = strings.ReplaceAll(rolePolicy, "{topic}", topic)
	fmt.Println("角色访问策略")
	fmt.Println(rolePolicy)

	// 返回值
	resultMap := objx.Map{
		"sqsPolicy":  mainMap,
		"rolePolicy": objx.MustFromJSON(rolePolicy),
	}
	return resultMap
}

// 替换内容
func replaceCon(str string) objx.Map {
	str = strings.ReplaceAll(str, "{id}", getId("Item"))
	str = strings.ReplaceAll(str, "{queue}", queue)
	str = strings.ReplaceAll(str, "{topic}", topic)
	str = strings.ReplaceAll(str, "{queue-root}", queueRootArn)
	return objx.MustFromJSON(str)
}

func getId(pre string) string {
	return pre + "-" + strconv.Itoa(int(time.Now().Unix()))
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
