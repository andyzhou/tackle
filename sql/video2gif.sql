-- base on sqlite3
-- db name: video2gif.db

-- users table
DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
userId integer PRIMARY KEY AUTOINCREMENT,
name varchar(64),
password varchar(32),
data integer default 0,
ip varchar(32),
status integer default 0,
activeAt integer default 0,
createAt integer default 0
);


-- users index
CREATE UNIQUE INDEX users_id_index ON users(userId);
CREATE INDEX users_name_index ON users(name);
CREATE INDEX users_createAt_index ON users(createAt);


-- files table
DROP TABLE IF EXISTS `files`;

CREATE TABLE `files` (
fileId integer PRIMARY KEY AUTOINCREMENT,
shortUrl varchar(64),
userId integer default 0,
md5 varchar(32),
snap varchar(64),
gif varchar(64),
tags varchar(128),
score integer default 0,
likes integer default 0,
downloads integer default 0,
createAt integer default 0
);

-- files index
CREATE UNIQUE INDEX files_id_index ON files(fileId);
CREATE UNIQUE INDEX files_us_index ON files(userId, shortUrl);
CREATE INDEX files_md5_index ON files(md5);
CREATE INDEX files_score_index ON files(score);
CREATE INDEX files_createAt_index ON files(createAt);


-- file opt log table
DROP TABLE IF EXISTS `fileOptLogs`;

CREATE TABLE `fileOptLogs` (
ip varchar(32),
shortUrl varchar(64),
liked integer default 0,
download integer default 0
);

-- file opt index
CREATE UNIQUE INDEX fileOptLogs_index ON fileOptLogs(ip);