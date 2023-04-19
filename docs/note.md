# Dodo

## 获得群主身份信息

[获取群信息](https://open.imdodo.com/dev/api/island.html#%E8%8E%B7%E5%8F%96%E7%BE%A4%E4%BF%A1%E6%81%AF)-> 响应中取ownerDodoSourceId -> [获取成员信息](https://open.imdodo.com/dev/api/member.html#%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF)

## 通过dodo号获取dodo source id
[获取成员dodo号映射列表](https://open.imdodo.com/dev/api/member.html#%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98dodo%E5%8F%B7%E6%98%A0%E5%B0%84%E5%88%97%E8%A1%A8)


# 数据结构

- SocialToolServer: {[]PushInfo, RainbowUser}
- PushInfo: {Activity, Channel, IdentityGroup, Msg, ColorTheme}


# 数据迁移
anywebUsers -> walletUsers
```sql
CREATE TABLE `wallet_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `wallet` varchar(256) DEFAULT NULL,
  `union_id` varchar(256) DEFAULT NULL,
  `access_token` text,
  `expire` int DEFAULT NULL,
  `refresh_token` text,
  `refresh_expire` int DEFAULT NULL,
  `phone` varchar(256) DEFAULT NULL,
  `address` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_wallet_users_deleted_at` (`deleted_at`),
  KEY `idx_wallet_users_union_id` (`union_id`),
  KEY `idx_wallet_users_phone` (`phone`),
  KEY `idx_wallet_users_address` (`address`),
  KEY `idx_wallet_phone` (`phone`,`wallet`)
)  ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

insert into wallet_users (created_at,updated_at,deleted_at,wallet,union_id,access_token,expire,refresh_token,refresh_expire,phone,address) select created_at,updated_at,deleted_at,1,union_id,access_token,expire,refresh_token,refresh_expire,phone,address from anyweb_users;
```
