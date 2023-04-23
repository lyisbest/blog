create database IF NOT EXISTS blog;

use blog;

CREATE TABLE `blog_log_tab` (
                                `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                `blog_id` bigint(20) DEFAULT NULL,
                                `operator` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                `operation_type` smallint(1) DEFAULT NULL,
                                `operation` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                `ctime` int(10) unsigned DEFAULT NULL,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COMMENT='博客日志表';

CREATE TABLE `blog_tab` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `title` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `content` text,
                            `creator` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `ctime` int(10) unsigned DEFAULT NULL,
                            `mtime` int(10) unsigned DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `title_UNIQUE` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COMMENT='博客表';

CREATE TABLE `user_tab` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
                            `username` varchar(10) NOT NULL COMMENT '用户名',
                            `password` varchar(10) NOT NULL COMMENT '密码',
                            `ctime` int(10) unsigned NOT NULL COMMENT '创建时间',
                            `mtime` int(10) unsigned NOT NULL COMMENT '修改时间',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

insert into user_tab(`username`, `password`, `ctime`, `mtime`) value('root', 'root', 0, 0);