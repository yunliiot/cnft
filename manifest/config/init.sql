create database nft_entity;
use nft_entity;
CREATE TABLE `nft_entity` (
          `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
          `entity_id` varchar(50) NOT NULL COMMENT '实体id##',
          `entity_type` varchar(32) NOT NULL COMMENT '实体类型，1:装备##',
          `token` varchar(128) NOT NULL COMMENT '区块链token##',
          `created_at` bigint(13) NOT NULL DEFAULT '0' COMMENT '创建时间##',
          `updated_at` bigint(13) NOT NULL DEFAULT '0' COMMENT '创建时间##',
          `deleted_at` bigint(13) NOT NULL DEFAULT '0' COMMENT '创建时间##',
          PRIMARY KEY (`id`),
          UNIQUE KEY `unq_key_index` (`entity_id`,`entity_type`),
          KEY `idx_token` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='nft-实体映射表##';