/*
 Navicat Premium Data Transfer

 Source Server         : cyb
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : 120.25.179.82:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 07/06/2022 15:53:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for relation
-- ----------------------------
DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '粉丝',
  `to_user_id` bigint(20) NULL DEFAULT NULL COMMENT '作者',
  `is_follow` tinyint(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id_user_id`(`user_id`) USING BTREE,
  INDEX `id_to_user_id`(`to_user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of relation
-- ----------------------------
INSERT INTO `relation` VALUES (1, 1, 10, 1);
INSERT INTO `relation` VALUES (2, 12, 1, 1);
INSERT INTO `relation` VALUES (6, 14, 9, 1);
INSERT INTO `relation` VALUES (7, 14, 1, 1);
INSERT INTO `relation` VALUES (9, 14, 11, 1);
INSERT INTO `relation` VALUES (10, 14, 10, 1);
INSERT INTO `relation` VALUES (11, 14, 13, 1);
INSERT INTO `relation` VALUES (13, 10, 9, 1);
INSERT INTO `relation` VALUES (14, 1, 9, 1);

SET FOREIGN_KEY_CHECKS = 1;
