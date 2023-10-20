# Amazon Marketing Stream API 订阅
> 广告小时级别数据订阅
> 
> 订单小时级别数据订阅 

## 介绍
Amazon Marketing Stream 是亚马逊广告业务的信息流.
在没有订阅的时候, 只能通过调用API去更新亚马逊广告数据, 不仅有大量的无效请求, 而且耗时很长, 导致时效性非常差.
本项目主要实现亚马逊店铺的营销数据流订阅功能, 在店铺很多的时候, 可以解放双手, 添加好配置后, 一键完成订阅功能.

## 准备
一、需要提前拿到店铺授权的如下信息: 
1. client_id 
2. client_secret
3. profile_id 站点id
4. refresh_token

二、在AWS账号下完成如下操作
1. 创建SNS的主题和订阅
2. 创建SQS, 配置访问策略!!!! , 必须先配置后才能调用创建订阅接口, 否则要等3天!!!(为什么等3天, 见订阅状态图)
3. SQS订阅SNS

注意: 如果因为配置的问题, 导致SQS一直未收到订阅确认消息, 而又不想等3天的话, 只能重新创建SQS队列, 确保配置没问题再创建订阅. 

## 功能
1. 生成 access token
2. 创建订阅
3. 查询订阅
4. 更新订阅
5. 一键生成SQS访问策略
6. 一键生成角色权限访问策略

## 订阅状态
![状态](https://d3a0d0y2hgofx6.cloudfront.net/en-us/_images/amazon-marketing-stream/state-diagram.png)

## 说明
广告: 目前支持的订阅数据集:
1. 商品推广流量（sp-traffic）
2. 商品推广转化 (sp-conversion)
3. 赞助广告预算使用 (budget-usage)
4. 展示流量（sd-traffic）
5. 展示转化（sd-conversion）
6. 品牌流量（sb-traffic）
7. 品牌转化（sb-conversion）

订单: 目前支持的类型: 
1. 订单项时间变化数据(ITEM_SALES_EVENT_CHANGE)

## Q&A
一、关于如何订阅, 请参考官网文档：

https://advertising.amazon.com/API/docs/en-us/guides/amazon-marketing-stream/overview

二、如何确认订阅
创建订阅后, 订阅状态是待确认, 随后AWS账号下的SQS上会收到待确认消息, 从消息体中拿到确认订阅URL, 放到浏览器上访问一下即可, 这样订阅才算最终成功, 亚马逊才能把广告数据近实时的推送到SQS上.

至于看了官方文档还是无法完成订阅, 可以发issue或邮件`jszls65@qq.com`探讨.
确实, 作者前期在调研的时候也是花了2周的时间, 才最终完成公共开发者账号和私有开发者账号的订阅.

# 最后

如果本项目对你有所启发或帮助, 欢迎📩交流, 欢迎Star