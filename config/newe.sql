/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50644
 Source Host           : localhost:3306
 Source Schema         : newe

 Target Server Type    : MySQL
 Target Server Version : 50644
 File Encoding         : 65001

 Date: 29/04/2022 15:09:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_depart
-- ----------------------------
DROP TABLE IF EXISTS `sys_depart`;
CREATE TABLE `sys_depart`  (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'uuid',
  `pid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父级ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分组名称（机构名称）',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '分组编码',
  `type` int(11) NULL DEFAULT NULL COMMENT '类型（1集团，2公司，3部门，4服务门店）',
  `telephone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系手机',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址',
  `sort_no` int(11) NULL DEFAULT NULL COMMENT '排序',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_depart
-- ----------------------------
INSERT INTO `sys_depart` VALUES ('c1a4ad97b18a404cbbe83a31d1ee02d6', '', '系统管理', 'A01', 1, '', '', '', 0, 1651113401, 1651113401);

-- ----------------------------
-- Table structure for sys_depart_rules
-- ----------------------------
DROP TABLE IF EXISTS `sys_depart_rules`;
CREATE TABLE `sys_depart_rules`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `depart_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织结构ID',
  `menu_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `org_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织结构编码',
  `create_time` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_depart_rules
-- ----------------------------
INSERT INTO `sys_depart_rules` VALUES ('0e2b01d324f341d6bcc1348cbdd12c65', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '90ce42219a504ba387c10f022dc51c2d', '', 1651138655);
INSERT INTO `sys_depart_rules` VALUES ('2882f2d51a7640f1a583514bd9f5cbc1', 'c1a4ad97b18a404cbbe83a31d1ee02d6', 'eeac301ece904d6e86f47e3c702aa134', '', 1651135189);
INSERT INTO `sys_depart_rules` VALUES ('2da6dcc2e6274b248960a4272c130c67', 'c1a4ad97b18a404cbbe83a31d1ee02d6', 'fc6e241275ab4b12b5e1d72835c5bbd7', '', 1651137183);
INSERT INTO `sys_depart_rules` VALUES ('406d42ed8ebc452c82cdba4cc30a8610', 'c1a4ad97b18a404cbbe83a31d1ee02d6', 'f4b8fe6baa844cfd890fa37f7092b831', '', 1651135189);
INSERT INTO `sys_depart_rules` VALUES ('51a72f1a78dc4275af88344287c71645', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '19092509384c483494921ab287b38f09', '', 1651113553);
INSERT INTO `sys_depart_rules` VALUES ('5a6257ce6e774565af45f74013f0af6d', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '41afa5e45e7e488c84a0966e89df73cb', '', 1651138655);
INSERT INTO `sys_depart_rules` VALUES ('5b8cc45eedae40dbb50acd8643ee8c61', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '226dfde007e44920a2dd18d416f63ac6', '', 1651138655);
INSERT INTO `sys_depart_rules` VALUES ('6a58d2b293f944e9a433a91c864b0ac8', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '05fab54d3b08457393b0abbb05691e01', '', 1651138655);
INSERT INTO `sys_depart_rules` VALUES ('6b5adb79da034e19b24290645b77b110', 'c1a4ad97b18a404cbbe83a31d1ee02d6', 'b27c1551aaa94a6bb64940d6c13f1694', '', 1651138655);
INSERT INTO `sys_depart_rules` VALUES ('951d1f7b98ab43418ad3b73f1bde1ece', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '3440cf7efb804fa78fee862f048a588b', '', 1651135189);
INSERT INTO `sys_depart_rules` VALUES ('aa2f8a76c5224bf89c63d62c3329131c', 'c1a4ad97b18a404cbbe83a31d1ee02d6', 'yingdaoye', '', 1651137183);
INSERT INTO `sys_depart_rules` VALUES ('b65ddaadbe564c4ea33aa9c8d2934f18', 'c1a4ad97b18a404cbbe83a31d1ee02d6', 'bb259e48d4294881aa0853e5683efc8e', '', 1651138655);
INSERT INTO `sys_depart_rules` VALUES ('b90a49f4d1dd493581e6f604e781d5a1', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '84cb2653afa7499e829f00bc0c8367f7', '', 1651113553);
INSERT INTO `sys_depart_rules` VALUES ('c33d1c786210467db10e2c058de483dc', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '93db86b5d7c54bf8b27a25710e5e4eff', '', 1651113553);
INSERT INTO `sys_depart_rules` VALUES ('c7320be6916d4188a63b940b3ad5f4be', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '4bb7482f10724915b40dd4b033937f44', '', 1651138655);
INSERT INTO `sys_depart_rules` VALUES ('da9aa7a7c3214d8aad537d0669957f6b', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '8083ccbd7359433a81e82f7c80f246d1', '', 1651113553);
INSERT INTO `sys_depart_rules` VALUES ('e57283730dd3417995cd98f96df5aa6d', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '98f3646ebee84e58b58f06265c2ea95f', '', 1651138655);
INSERT INTO `sys_depart_rules` VALUES ('fcfedd4ff6d34eabbd9620105ef74316', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '3baf203ee0dd45e3accaaae6a6412aa6', '', 1651138655);

-- ----------------------------
-- Table structure for sys_member
-- ----------------------------
DROP TABLE IF EXISTS `sys_member`;
CREATE TABLE `sys_member`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `depart_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织结构ID',
  `uid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '会员ID',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '登陆账号',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `realname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `headimgurl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `mp` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `idcard` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '身份证号码',
  `sex` int(1) NULL DEFAULT NULL COMMENT '性别 1男2女',
  `status` int(1) NULL DEFAULT NULL COMMENT '1正常，2禁用',
  `org_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组织结构编码',
  `create_time` bigint(20) NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` bigint(20) NULL DEFAULT NULL COMMENT '更新时间',
  `files` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附件',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_member
-- ----------------------------
INSERT INTO `sys_member` VALUES ('72bf2cad7f6f4a3f8acd8d4808885293', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '', 'admin', '21232f297a57a5a743894a0e4a801fc3', '管理员', '管理员', '', '13333333333', '', 1, 0, '', 1651113987, 1651113987, '');
INSERT INTO `sys_member` VALUES ('cbf4182e44eb40e285b3fe9b26eed6bb', 'c1a4ad97b18a404cbbe83a31d1ee02d6', '', 'test', '21232f297a57a5a743894a0e4a801fc3', '昵称', '测试', '', '13333333331', '', 1, 0, '', 1651135876, 1651141289, '');

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `pid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组件地址',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片',
  `is_ext` int(1) NULL DEFAULT NULL COMMENT '是否外链',
  `keepalive` int(1) NULL DEFAULT NULL COMMENT '是否缓存',
  `show` int(1) NULL DEFAULT NULL COMMENT '是否显示',
  `type` int(1) NULL DEFAULT NULL COMMENT '类型 1目录2菜单 3按钮',
  `sort_no` int(11) NULL DEFAULT NULL COMMENT '排序',
  `route_path` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由地址',
  `permission` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限编码',
  `status` int(1) NULL DEFAULT NULL COMMENT '是否启用 0启用1禁用',
  `create_time` bigint(20) NULL DEFAULT NULL,
  `is_iframe` int(1) NULL DEFAULT NULL COMMENT '是否嵌套',
  `route_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES ('05fab54d3b08457393b0abbb05691e01', 'yingdaoye', '文件下载', '/demo/feat/download/index', 'ant-design:arrow-down-outlined', 2, 1, 1, 2, 6, 'download', '', 1, 1651138028, 2, 'download');
INSERT INTO `sys_menus` VALUES ('19092509384c483494921ab287b38f09', '', '系统管理', '', 'ant-design:setting-outlined', 0, 0, 1, 1, 1, '/system', '', 1, 1649857779, 0, 'newesys');
INSERT INTO `sys_menus` VALUES ('226dfde007e44920a2dd18d416f63ac6', 'yingdaoye', '标签', '/demo/feat/tabs/index', 'ant-design:check-square-outlined', 2, 1, 1, 2, 4, 'tabs', '', 1, 1651137896, 2, 'tabs');
INSERT INTO `sys_menus` VALUES ('3440cf7efb804fa78fee862f048a588b', '', 'Dashboard', '/dashboard/analysis/index', 'ant-design:home-outlined', 0, 0, 1, 2, 0, '/dashboard', '', 1, 1649868128, 0, 'dashboard');
INSERT INTO `sys_menus` VALUES ('3baf203ee0dd45e3accaaae6a6412aa6', 'yingdaoye', 'websocket测试', '/demo/feat/ws/index', 'ant-design:branches-outlined', 2, 2, 1, 2, 2, 'wstest', '', 1, 1651140042, 2, 'wstest');
INSERT INTO `sys_menus` VALUES ('41afa5e45e7e488c84a0966e89df73cb', 'yingdaoye', '消息提示', '/demo/feat/msg/index', 'ant-design:aliwangwang-outlined', 2, 1, 1, 2, 9, 'msg', '', 1, 1651138248, 2, 'msg');
INSERT INTO `sys_menus` VALUES ('4bb7482f10724915b40dd4b033937f44', 'yingdaoye', '打印', '/demo/feat/print/index', 'ant-design:check-square-twotone', 2, 1, 1, 2, 3, 'parent', '', 1, 1651137811, 2, 'parent');
INSERT INTO `sys_menus` VALUES ('8083ccbd7359433a81e82f7c80f246d1', '19092509384c483494921ab287b38f09', '账号管理', '/system/member/index', 'ant-design:usergroup-add-outlined', 2, 1, 1, 2, 0, 'SystemMember', '', 1, 1651052047, 2, 'member');
INSERT INTO `sys_menus` VALUES ('84cb2653afa7499e829f00bc0c8367f7', '19092509384c483494921ab287b38f09', '菜单管理', '/system/menu/index', 'ant-design:menu-unfold-outlined', 2, 2, 1, 2, 1, 'menus', '', 1, 1650025127, 2, 'NeweMenus');
INSERT INTO `sys_menus` VALUES ('90ce42219a504ba387c10f022dc51c2d', 'yingdaoye', '图片预览', '/demo/feat/img-preview/index', 'ant-design:camera-outlined', 2, 1, 1, 2, 7, 'img-preview', '', 1, 1651138137, 2, 'img-preview');
INSERT INTO `sys_menus` VALUES ('93db86b5d7c54bf8b27a25710e5e4eff', '19092509384c483494921ab287b38f09', '组织结构', '/system/depart/index', 'ant-design:apartment-outlined', 2, 2, 1, 2, 2, 'depart', '', 1, 1650025115, 2, 'SystemDepart');
INSERT INTO `sys_menus` VALUES ('98f3646ebee84e58b58f06265c2ea95f', 'yingdaoye', '右键菜单', '/demo/feat/context-menu/index', 'ant-design:check-circle-twotone', 2, 1, 1, 2, 5, 'context-menu', '', 1, 1651137945, 2, 'context-menu');
INSERT INTO `sys_menus` VALUES ('b27c1551aaa94a6bb64940d6c13f1694', 'yingdaoye', '剪切板', '/demo/feat/copy/index', 'ant-design:ci-circle-filled', 2, 1, 1, 2, 8, 'copy', '', 1, 1651138183, 2, 'copy');
INSERT INTO `sys_menus` VALUES ('bb259e48d4294881aa0853e5683efc8e', 'yingdaoye', '图标', '/demo/feat/icon/index', 'ant-design:alert-outlined', 2, 1, 1, 2, 1, 'icon', '', 1, 1651137717, 2, 'icon');
INSERT INTO `sys_menus` VALUES ('eeac301ece904d6e86f47e3c702aa134', '3440cf7efb804fa78fee862f048a588b', 'analysis', '/dashboard/analysis/index', 'ant-design:home-outlined', 2, 2, 1, 2, 1, 'analysis', '', 1, 1650025137, 2, 'analysis');
INSERT INTO `sys_menus` VALUES ('f4b8fe6baa844cfd890fa37f7092b831', '3440cf7efb804fa78fee862f048a588b', 'workbench', '/demo/tree/index', 'ant-design:appstore-add-outlined', 0, 0, 1, 2, 2, 'workbench', '', 1, 1649868280, 0, 'workbench');
INSERT INTO `sys_menus` VALUES ('fc6e241275ab4b12b5e1d72835c5bbd7', '', 'demo', '', 'ant-design:appstore-filled', 2, 0, 1, 1, 100, '/demo', '', 1, 1651136209, 2, 'Demo');
INSERT INTO `sys_menus` VALUES ('yingdaoye', 'fc6e241275ab4b12b5e1d72835c5bbd7', '功能', '', 'ant-design:appstore-add-outlined', 2, 0, 1, 1, 100, 'feal', '', 1, 1651137679, 2, 'FeatDemo');

-- ----------------------------
-- Table structure for sys_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_records`;
CREATE TABLE `sys_records`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `state` bigint(20) NULL DEFAULT NULL COMMENT '请求状态',
  `api` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求地址',
  `create_time` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
