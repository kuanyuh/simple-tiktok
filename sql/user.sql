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

 Date: 07/06/2022 15:53:08
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
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'yuuki', 2, 2, 0);
INSERT INTO `user` VALUES (9, '二哈', 0, 3, 1);
INSERT INTO `user` VALUES (10, 'tuuol3', 1, 1, 1);
INSERT INTO `user` VALUES (11, '张三丰', 1, 2, 0);
INSERT INTO `user` VALUES (12, 'tahammi6', 1, 0, 0);
INSERT INTO `user` VALUES (13, 'hello_pig', 0, 1, 0);
INSERT INTO `user` VALUES (14, 'o6tvpkll', 5, 0, 0);
INSERT INTO `user` VALUES (15, '8h63pnfy', 0, 0, 0);
INSERT INTO `user` VALUES (16, '8topwnbu', 0, 0, 0);
INSERT INTO `user` VALUES (17, 'lengujky', 0, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
