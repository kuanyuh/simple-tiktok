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

 Date: 07/06/2022 15:53:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author_id` bigint(20) NULL DEFAULT NULL,
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `favorite_count` bigint(20) NULL DEFAULT NULL,
  `comment_count` bigint(20) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id_author_id`(`author_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, 1, 'https://www.w3schools.com/html/movie.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 5, 0, 'bear', 2);
INSERT INTO `video` VALUES (2, 10, 'https://oss-p56.xpccdn.com/prod/footage/preview/j2MvZd1p5G.mp4', 'https://kuanyuh.github.io/img/favicon.png', 6, 0, 'night', 1);
INSERT INTO `video` VALUES (3, 11, 'http://100.120.62.90:8080/static/v2.mp4', 'https://kuanyuh.github.io/img/favicon.png', 8, 0, 'sky', 3);
INSERT INTO `video` VALUES (4, 13, 'http://121.5.225.209:9000/product/product_1654357481775.mp4', 'http://121.5.225.209:9000/product/product_1654357233945.jpg', 1098, 0, 'happy', 4);
INSERT INTO `video` VALUES (5, 9, 'http://121.5.225.209:9000/product/product_1654358443093.mp4', 'http://121.5.225.209:9000/product/product_1654358548050.png', 10, 0, 'cat', 5);

SET FOREIGN_KEY_CHECKS = 1;
