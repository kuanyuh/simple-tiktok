/*
 Navicat Premium Data Transfer

 Source Server         : ESC
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : 120.25.179.82:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 24/05/2022 18:24:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `follow_count` bigint(20) NULL DEFAULT NULL,
  `follower_count` bigint(20) NULL DEFAULT NULL,
  `is_follow` tinyint(4) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
