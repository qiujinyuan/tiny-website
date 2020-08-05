create database tiny_website;
use tiny_website;
drop table blog_tag;
drop table blog_article;

CREATE TABLE `blog_tag` (
  `id` VARCHAR(36) NOT NULL,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_at` int(10) DEFAULT NULL COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) DEFAULT NULL COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) DEFAULT NULL,
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

CREATE TABLE `blog_article` (
  `id` VARCHAR(36) NOT NULL,
  `tag_id` VARCHAR(36) COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_at` int(11) DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_at` int(10) DEFAULT NULL COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_at` int(10) DEFAULT NULL,
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';

CREATE TABLE `blog_auth` (
  `id` VARCHAR(36) NOT NULL,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `tiny_website`.`blog_auth` (`id`, `username`, `password`) VALUES (uuid(), 'test', 'test123456');

INSERT INTO `tiny_website`.`blog_tag` (`id`, `name`) VALUES (uuid(), 'default');

INSERT INTO `tiny_website`.`blog_article` (`id`, `tag_id`, `title`, `desc`, `content`) VALUES (uuid(), '1eb562e9-d6ea-11ea-a475-0242ac110002', 'default title', 'default desc', 'default content');