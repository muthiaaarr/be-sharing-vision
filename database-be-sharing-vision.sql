/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80024
 Source Host           : localhost:3306
 Source Schema         : be-sharing-vision

 Target Server Type    : MySQL
 Target Server Version : 80024
 File Encoding         : 65001

 Date: 06/05/2021 14:59:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (10, 'username3', '$2a$14$71eJcxN2yqYYbcA9.HsWr.EP2l48wj6SYeB5mSsQ4QhRrEe3nc8MG', 'name3');
INSERT INTO `users` VALUES (11, 'username4', '$2a$14$HRC7KOH4ayI9RTZpeq3ceO3/.EzcRDVgkK/x/d/q5lj/I0OVamRnO', 'name4');
INSERT INTO `users` VALUES (12, 'username5', '$2a$14$n3TOaunGTWytvFFMORU9h.nI0QWarn6HoTcOre9Pc4M6WszqzgsGW', 'name5');
INSERT INTO `users` VALUES (13, 'username6', '$2a$14$QCV4ZcprE/XF22.1340PaebHV8EqFuRNajeVEJWi5b1fyR0RQkLke', 'name6');

SET FOREIGN_KEY_CHECKS = 1;
