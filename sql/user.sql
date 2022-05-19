/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.8.59-MySQL
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 19/05/2022 10:52:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `follow_count` bigint(0) NULL DEFAULT NULL,
  `follower_count` bigint(0) NULL DEFAULT NULL,
  `is_follow` tinyint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_name`(`name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'zs', 'zs', 'zs', 0, 0, 0);
INSERT INTO `user` VALUES (2, 'ls', 'ls', '李四', 0, 0, 0);
INSERT INTO `user` VALUES (3, 'ww', 'ww', '王五', 0, 0, 0);
INSERT INTO `user` VALUES (4, 'zl', 'zl', '赵六', 0, 0, 0);
INSERT INTO `user` VALUES (5, 'sq', 'sq', '孙七', 0, 0, 0);
INSERT INTO `user` VALUES (6, 'zb', 'zb', '周八', 0, 0, 0);
INSERT INTO `user` VALUES (7, 'wj', 'wj', '吴九', 0, 0, 0);
INSERT INTO `user` VALUES (8, 'zs', 'zs', '郑十', 0, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
