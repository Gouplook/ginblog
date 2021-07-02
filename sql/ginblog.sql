/*
 Navicat Premium Data Transfer

 Source Server         : Mac
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : localhost:3306
 Source Schema         : ginblog

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 02/07/2021 13:18:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文章标题 ',
  `cid` bigint(20) NOT NULL COMMENT '''文章Cid''',
  `decs` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '''文章描述''',
  `content` longtext COLLATE utf8mb4_unicode_ci COMMENT '''文章内容''',
  `read_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '''文章阅读量''',
  `comment_count` bigint(20) NOT NULL DEFAULT '0' COMMENT '''文章评论量''',
  `img` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '''文章相册''',
  `category_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '''文章分类''',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 0-未删除，1-已删除',
  PRIMARY KEY (`id`),
  KEY `idx_article_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` VALUES (1, '2021-07-01 11:56:35.017', '2021-07-01 11:56:35.017', NULL, '比利时VS意大利', 1, '比利时VS意大利', '比利时VS意大利，于7月3日 0:00 开赛 ', 0, 0, '001', 1, 0);
INSERT INTO `article` VALUES (2, '2021-07-01 13:07:18.452', '2021-07-01 13:07:18.452', NULL, '乌克兰VS英格兰', 1, '乌克兰VS英格兰', '乌克兰VS英格兰，于7月3日 03:00 开赛 ', 0, 0, '001', 1, 0);
INSERT INTO `article` VALUES (3, '2021-07-01 13:09:23.969', '2021-07-01 13:09:23.969', NULL, '捷克VS丹麦', 1, '捷克VS丹麦', '捷克VS丹麦，于7月4日 03:00 开赛 ', 0, 0, '001', 1, 0);
INSERT INTO `article` VALUES (4, '2021-07-01 13:09:56.887', '2021-07-01 13:41:46.284', NULL, '瑞士～VS西班牙～', 1, '瑞士VS西班牙', '瑞士VS西班牙，于7月4日 00:00 开赛 ', 3, 0, '001', 1, 0);
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_del` tinyint(4) NOT NULL COMMENT '是否删除分类 0-未删除，1-删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1, '体育新闻', 0);
INSERT INTO `category` VALUES (2, '时政新闻', 0);
INSERT INTO `category` VALUES (3, '互联网新闻', 0);
INSERT INTO `category` VALUES (5, '历史要闻', 0);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `role` bigint(20) DEFAULT NULL,
  `is_del` tinyint(4) DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, '2021-06-30 15:57:32.703', '2021-07-01 09:41:02.285', '2021-07-01 09:51:31.322', 'admin1', 'l+23QUVIZIrEwH4Iekn8eQeCWlOa7z1LxK70IswqvqQ=', 0, 0);
INSERT INTO `user` VALUES (2, '2021-06-30 15:57:41.292', '2021-06-30 17:46:06.905', NULL, 'admin2', 'l+23QUVIZIrEwH4Iekn8eQeCWlOa7z1LxK70IswqvqQ=', 0, 0);
INSERT INTO `user` VALUES (3, '2021-06-30 15:58:22.503', '2021-06-30 15:58:22.503', NULL, 'admin3', 'l+23QUVIZIrEwH4Iekn8eQeCWlOa7z1LxK70IswqvqQ=', 0, 0);
INSERT INTO `user` VALUES (4, '2021-06-30 15:59:45.643', '2021-06-30 15:59:45.643', NULL, 'admin4', 'l+23QUVIZIrEwH4Iekn8eQeCWlOa7z1LxK70IswqvqQ=', 0, 0);
INSERT INTO `user` VALUES (5, '2021-06-30 16:28:07.225', '2021-06-30 17:55:54.558', NULL, 'admin5', 'l+23QUVIZIrEwH4Iekn8eQeCWlOa7z1LxK70IswqvqQ=', 0, 0);
INSERT INTO `user` VALUES (6, '2021-06-30 16:28:11.298', '2021-06-30 16:28:11.298', NULL, 'admin6', 'l+23QUVIZIrEwH4Iekn8eQeCWlOa7z1LxK70IswqvqQ=', 0, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
