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

 Date: 07/06/2022 15:52:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `video_id` bigint(20) NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_date` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `deleted_at` bigint(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id_video_id`(`video_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (10, 5, 17, '1', '06-06', 1654519975);
INSERT INTO `comment` VALUES (11, 4, 17, '2', '06-06', 1654519964);
INSERT INTO `comment` VALUES (12, 5, 17, '真不错', '06-06', 1654522060);
INSERT INTO `comment` VALUES (13, 5, 17, '1223', '06-06', 1654522074);
INSERT INTO `comment` VALUES (14, 5, 17, '2', '06-06', 1654522137);
INSERT INTO `comment` VALUES (15, 5, 17, '1', '06-06', 1654522140);

SET FOREIGN_KEY_CHECKS = 1;
