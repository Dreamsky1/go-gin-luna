/*
Navicat MySQL Data Transfer

Source Database       : bill

Target Server Type    : MYSQL
Target Server Version : 50639
File Encoding         : 65001

Date: 2022-05-07 19:52:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for 用户user
-- ----------------------------
CREATE TABLE `bill_user` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `username` varchar(50) DEFAULT '' COMMENT '账号',
 `password` varchar(50) DEFAULT '' COMMENT '密码',
 `avatar`    varchar(255) DEFAULT '' COMMENT '头像',
 `signature` varchar(100) DEFAULT '' COMMENT '个性签名',
 `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
 `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
 `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `integral`  int(10) DEFAULT '0' COMMENT '个人积分',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for 一级分类
-- ----------------------------
CREATE TABLE `bill_type_category` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
     `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
     `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `name` varchar(100) DEFAULT '' COMMENT '一级分类名',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for bill_bill
-- 账单bill
-- ----------------------------
CREATE TABLE `bill_bill` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int(10) unsigned NOT NULL COMMENT '用户的ID',
    `type_id` int(10) unsigned DEFAULT '0' COMMENT '分类类型的ID',
    `category_id` int(10) unsigned DEFAULT '0' COMMENT '分类的ID',
    `category_name` varchar(100) DEFAULT '' COMMENT '分类名称',
    `accounting_date` int(11) DEFAULT '0' COMMENT '记账日期',
    `amount` int(10) unsigned DEFAULT '0' COMMENT '价格',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    `remark` varchar(255) DEFAULT '' COMMENT '备注',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='账单';

-- ----------------------------
-- Table structure for bill_category
-- 二级分类：交通、娱乐....,
-- ----------------------------
CREATE TABLE `bill_category` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '分类名称',
    `type_id` int(10) unsigned NOT NULL COMMENT '分类类型',
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    `image` varchar(225) DEFAULT '' COMMENT '图片名称',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='一级分类';

