
CREATE DATABASE IF NOT EXISTS jixiang default charset utf8 COLLATE utf8_general_ci;
use jixiang;


CREATE TABLE `auth` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) DEFAULT '' COMMENT '账号',
    `password` varchar(50) DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `auth` (`id`, `username`, `password`) VALUES (null, 'jixiang', 'jixiang2021');