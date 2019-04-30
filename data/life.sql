-- --------------------------------------------------------
-- 主机:                           118.25.90.130
-- 服务器版本:                        5.7.23-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Linux
-- HeidiSQL 版本:                  9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 life 的数据库结构
CREATE DATABASE IF NOT EXISTS `life` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;
USE `life`;

-- 导出  表 life.asset_bill 结构
CREATE TABLE IF NOT EXISTS `asset_bill` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `io` char(1) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '收支',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `tid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类型ID',
  `tag` varchar(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '标签',
  `money` decimal(12,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '金额',
  `remark` varchar(500) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '备注',
  `date_time` datetime NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '发生时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `user_delete` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除用户ID',
  `user_create` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建用户ID',
  `user_update` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新用户ID',
  `time_delete` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `time_create` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `time_update` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='收支表';

-- 正在导出表  life.asset_bill 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `asset_bill` DISABLE KEYS */;
/*!40000 ALTER TABLE `asset_bill` ENABLE KEYS */;

-- 导出  表 life.asset_conf 结构
CREATE TABLE IF NOT EXISTS `asset_conf` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级ID',
  `sort` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `path` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '路径',
  `name` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '名称',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `user_delete` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除用户ID',
  `user_create` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建用户ID',
  `user_update` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新用户ID',
  `time_delete` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `time_create` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `time_update` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='配置表';

-- 正在导出表  life.asset_conf 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `asset_conf` DISABLE KEYS */;
REPLACE INTO `asset_conf` (`id`, `uid`, `pid`, `sort`, `path`, `name`, `status`, `user_delete`, `user_create`, `user_update`, `time_delete`, `time_create`, `time_update`) VALUES
	(64, 1, 0, 0, '0', '笔记', 0, 0, 1, 0, 0, 1542713715, 0);
/*!40000 ALTER TABLE `asset_conf` ENABLE KEYS */;

-- 导出  表 life.asset_note 结构
CREATE TABLE IF NOT EXISTS `asset_note` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `tid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类型ID',
  `tag` varchar(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '标签',
  `title` varchar(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '标题',
  `content` mediumtext COLLATE utf8mb4_bin NOT NULL COMMENT '内容',
  `date_time` datetime NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT '记录时间',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `user_delete` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除用户ID',
  `user_create` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建用户ID',
  `user_update` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新用户ID',
  `time_delete` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `time_create` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `time_update` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='笔记表';

-- 正在导出表  life.asset_note 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `asset_note` DISABLE KEYS */;
REPLACE INTO `asset_note` (`id`, `uid`, `tid`, `tag`, `title`, `content`, `date_time`, `status`, `user_delete`, `user_create`, `user_update`, `time_delete`, `time_create`, `time_update`) VALUES
	(6, 1, 64, '测试', '测试', '<pre class="ql-syntax" spellcheck="false"><span class="hljs-meta">&lt;?php</span>\n<span class="hljs-keyword">echo</span> <span class="hljs-string">\'你好\'</span>﻿\n</pre><p><br></p><p><br></p><p><img src="http://118.25.90.130/static/upload/image/20181120/ee16db0af7868fb0ae07da8a38c613f7.jpg" width="118"></p>', '2018-11-20 19:26:00', 0, 0, 1, 1, 0, 1542713788, 1542715828),
	(7, 1, 64, '哈哈', '哈哈', '<p><br></p><p><img src="http://127.0.0.1:3000/static/upload/image/20181120/2ae6ed4b502760d1babd4bdd059d9e01.jpg" width="201" style="cursor: nwse-resize;"></p>', '2018-11-20 20:20:00', 0, 0, 1, 1, 0, 1542716509, 1542717675);
/*!40000 ALTER TABLE `asset_note` ENABLE KEYS */;

-- 导出  表 life.asset_user 结构
CREATE TABLE IF NOT EXISTS `asset_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `username` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '账号',
  `nickname` varchar(50) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(200) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `user_delete` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除用户ID',
  `user_create` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建用户ID',
  `user_update` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新用户ID',
  `time_delete` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  `time_create` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `time_update` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

-- 正在导出表  life.asset_user 的数据：~3 rows (大约)
/*!40000 ALTER TABLE `asset_user` DISABLE KEYS */;
REPLACE INTO `asset_user` (`id`, `username`, `nickname`, `password`, `status`, `user_delete`, `user_create`, `user_update`, `time_delete`, `time_create`, `time_update`) VALUES
	(1, 'admin', '管理员', '$2a$10$LQQPJKXZpLDTqtTdtWYjM.zvcXARohwGz/u5MZCrQH8eYUaRJr6pu', 0, 0, 0, 0, 0, 1536309554, 0),
	(2, 'testuser', '测试号', '$2a$10$k5af2DpYqpcAPS/T1dY0ROzWrEXNqUAgEOn6OfNGzyIx6gc5dlwhq', 0, 0, 0, 0, 0, 1541553204, 0),
	(3, 'wnewstar', '汪新星', '$2a$10$Z5by99552xV0lEoYds3OQ.CTCqce10cEOd0DCutB4ODHakmhAq2Uq', 0, 0, 0, 0, 0, 1535800573, 0);
/*!40000 ALTER TABLE `asset_user` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
