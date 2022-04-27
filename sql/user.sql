CREATE DATABASE IF NOT EXISTS `db_youcaibi`
DEFAULT CHARACTER SET utf8
DEFAULT COLLATE utf8_chinese_cli;

DROP TABLE IF EXISTS 't_users';
CREATE TABLE IF NOT EXISTS 't_users' (
    `id` UNSIGNED INT PRIMARY KEY AUTO_INCREMENT,
    `login_name` VARCHAR(64) UNIQUE KEY,
    `pwd` TEXT
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS 't_video_info';
CREATE TABLE IF NOT EXISTS 't_video_info' (
    `id` VARCHAR(64) PRIMARY KEY NOT NULL,
    `author_id` UNSIGNED INT,
    `name` TEXT,
    `display_ctime` TEXT,
    `create_time` DATETIME
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS 't_comments';
CREATE TABLE IF NOT EXISTS 't_comments' (
    `id` VARCHAR(64) PRIMARY KEY NOT NULL,
    `author_id` UNSIGNED INT,
    `video_id` VARCHAR(64),
    `content` TEXT,
    `time` DATETIME
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS 't_sessions';
CREATE TABLE IF NOT EXISTS 't_sessions' (
    `session_id` TINYTEXT PRIMARY KEY NOT NULL,
    `TTL` TINYTEXT,
    `login_name` VARCHAR(64)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;