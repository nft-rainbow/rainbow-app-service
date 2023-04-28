use rainbow_dev_apps;
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
)  ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

insert into wallet_users (created_at,updated_at,deleted_at,wallet,union_id,access_token,expire,refresh_token,refresh_expire,phone,address) select created_at,updated_at,deleted_at,1,union_id,access_token,expire,refresh_token,refresh_expire,phone,address from anyweb_users;

-- activities
alter table white_list_infos DROP foreign key fk_activities_white_list_infos;
alter table white_list_infos DROP index fk_activities_white_list_infos;

alter table nft_configs DROP foreign key fk_activities_nft_configs;
alter table nft_configs DROP INDEX fk_activities_nft_configs;

drop table if exists activities;

CREATE TABLE `activities` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `app_id` bigint unsigned DEFAULT NULL,
  `amount` int DEFAULT NULL,
  `name` longtext,
  `description` longtext,
  `app_name` longtext,
  `activity_type` bigint unsigned DEFAULT NULL,
  `command` longtext,
  `is_phone_white_list_opened` tinyint(1) DEFAULT '0',
  `ended_time` int DEFAULT NULL,
  `started_time` int DEFAULT NULL,
  `max_mint_count` varchar(256) DEFAULT NULL,
  `metadata_uri` longtext,
  `activity_picture_url` longtext,
  `contract_raw_id` varchar(32) DEFAULT NULL,
  `activity_code` varchar(191) DEFAULT NULL,
  `rainbow_user_id` int DEFAULT NULL,
  `activity_poster_url` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_activities_deleted_at` (`deleted_at`),
  KEY `idx_activities_app_id` (`app_id`),
  KEY `idx_activities_activity_code` (`activity_code`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

insert into activities (id,created_at,updated_at,deleted_at,app_id,`amount`,`name`,`description`,`app_name`,`activity_type`,`command`,`is_phone_white_list_opened`,`ended_time`,`started_time`,`max_mint_count`,`metadata_uri`,`activity_picture_url`,`contract_raw_id`,`activity_code`,`rainbow_user_id`,`activity_poster_url`) select id,created_at,updated_at,deleted_at,app_id,`amount`,`name`,`description`,`app_name`,`activity_type`,`command`,`is_phone_white_list_opened`,`ended_time`,`started_time`,`max_mint_count`,`metadata_uri`,`activity_picture_url`,`contract_id`,`activity_id`,`rainbow_user_id`,`activity_poster_url` FROM poap_activity_configs;

update activities set contract_raw_id=NULL where contract_raw_id=0;


-- contracts
CREATE TABLE `contracts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `contract_raw_id` int DEFAULT NULL,
  `contract_address` longtext,
  `contract_type` int DEFAULT NULL,
  `chain_id` int DEFAULT NULL,
  `chain_type` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_contracts_contract_raw_id` (`contract_raw_id`),
  KEY `idx_contracts_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

insert into contracts (created_at,updated_at,deleted_at, contract_raw_id, contract_address, contract_type, chain_id, chain_type)  select created_at,updated_at,deleted_at,contract_id, contract_address, contract_type, chain_id, chain_type from poap_activity_configs where id in (SELECT min(ID) FROM poap_activity_configs where contract_id>0 GROUP BY(contract_id));

-- drop foreign keys
alter table white_list_infos DROP foreign key fk_poap_activity_configs_white_list_infos;
alter table white_list_infos DROP index fk_poap_activity_configs_white_list_infos;

alter table nft_configs DROP foreign key fk_poap_activity_configs_nft_configs;
alter table nft_configs DROP INDEX fk_poap_activity_configs_nft_configs;

-- drop tables
drop table if exists bind_cfxes;
drop table if exists poap_activity_configs;

-- rename columns
alter table white_list_infos drop column activity_id;
alter table white_list_infos rename column poap_activity_config_id to activity_id;
alter table nft_configs drop column activity_id;
alter table nft_configs rename column poap_activity_config_id to activity_id;
