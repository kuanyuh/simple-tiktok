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

 Date: 07/06/2022 15:52:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for login
-- ----------------------------
DROP TABLE IF EXISTS `login`;
CREATE TABLE `login`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of login
-- ----------------------------
INSERT INTO `login` VALUES (1, '1119526092@qq.com', '123456');
INSERT INTO `login` VALUES (9, '234552@qq.com', '222222');
INSERT INTO `login` VALUES (10, '23776548@qq.com', '888888');
INSERT INTO `login` VALUES (11, '12343552@qq.com', 'aaaaaa');
INSERT INTO `login` VALUES (12, '2396083753@qq.com', '12341233');
INSERT INTO `login` VALUES (13, '1001@qq.com', 'qwerqwer');
INSERT INTO `login` VALUES (14, 'chenyb@163.com', '123456');
INSERT INTO `login` VALUES (15, 'zhuxk', '20010213');
INSERT INTO `login` VALUES (16, 'cyb', '123456');
INSERT INTO `login` VALUES (17, '12', '123456');

SET FOREIGN_KEY_CHECKS = 1;
