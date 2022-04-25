/*
Navicat MySQL Data Transfer

Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50639
File Encoding         : 65001

Date: 2018-03-18 16:52:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text COMMENT '内容',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `blog_auth` (`id`, `username`, `password`) VALUES ('1', 'test', 'test123');

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';


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
    `type_id` int(10) unsigned DEFAULT '0' COMMENT '分类类型的ID',
    `category_id` int(10) unsigned DEFAULT '0' COMMENT '分类的ID',
    `category_name` varchar(100) DEFAULT '' COMMENT '分类名称',
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

-- ----------------------------
-- Table structure for bill_type
-- 支出和收入,
-- ----------------------------

CREATE TABLE  `bill_category_type` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) DEFAULT '' COMMENT '分类类型',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='一级分类';

