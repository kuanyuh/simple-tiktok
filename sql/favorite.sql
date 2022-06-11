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

 Date: 07/06/2022 15:52:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id',
  `video_id` bigint(20) NULL DEFAULT NULL COMMENT '用户已点赞的视频id',
  `is_favorite` tinyint(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id_user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite
-- ----------------------------
INSERT INTO `favorite` VALUES (14, 16, 5, 2);
INSERT INTO `favorite` VALUES (15, 16, 4, 2);
INSERT INTO `favorite` VALUES (16, 16, 3, 1);
INSERT INTO `favorite` VALUES (17, 16, 1, 2);
INSERT INTO `favorite` VALUES (18, 16, 2, 1);
INSERT INTO `favorite` VALUES (19, 15, 5, 2);
INSERT INTO `favorite` VALUES (20, 15, 4, 1);
INSERT INTO `favorite` VALUES (21, 15, 3, 1);
INSERT INTO `favorite` VALUES (22, 15, 1, 1);
INSERT INTO `favorite` VALUES (23, 15, 2, 1);
INSERT INTO `favorite` VALUES (24, 1, 3, 1);

SET FOREIGN_KEY_CHECKS = 1;
