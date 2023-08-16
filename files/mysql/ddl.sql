-- by zls at 2023-08-07
CREATE TABLE `t_amz_stream_subscribe` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `profile_id` bigint DEFAULT NULL COMMENT '站点id',
  `shop_name` varchar(100) DEFAULT NULL COMMENT '店铺名称',
  `seller_id` varchar(30) NOT NULL COMMENT '卖家Sellerid',
  `marketplace_id` varchar(15) NOT NULL COMMENT '市场id',
  `config_id` int NOT NULL DEFAULT '0' COMMENT 'nacos配置中的id',
  `sqs_arn` varchar(50) DEFAULT NULL COMMENT 'SQS arn',
  `client_id` varchar(100) DEFAULT NULL COMMENT '开发者账号 client_id',
  `client_secret` varchar(100) DEFAULT NULL COMMENT '开发者账号 client_secret',
  `access_token` varchar(500) DEFAULT NULL COMMENT 'access token 3600s失效',
  `access_token_ttl` int DEFAULT '3600' COMMENT 'access token 3600s失效',
  `access_token_time` datetime DEFAULT NULL COMMENT 'access token创建时间',
  `account_type` tinyint DEFAULT '0' COMMENT '开发者账号类型: 0-私有账号, 1-公共账号',
  `refresh_token` varchar(500) DEFAULT NULL COMMENT '开发者账号的 refresh token',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_shopname` (`shop_name`)
) COMMENT='amazon marketing stream流订阅信息';

-- 2023-08-11 添加主题和iam root arn by zls
alter table t_amz_stream_subscribe add column `topic_arn` varchar(100) null comment '主题arn' after sqs_arn;
alter table t_amz_stream_subscribe add column `iam_root` varchar(100) null comment 'aws账号iam root' after sqs_arn;

-- 2023-08-14 添加 destinations 的id和name by zls
alter table t_amz_stream_subscribe add column `destinationId` varchar(100) null comment 'sqs的destination的id, 亚马逊返回' after refresh_token;
alter table t_amz_stream_subscribe add column `destinationName` varchar(100) null comment 'sqs的destination的name, 开发者定义' after destinationId;