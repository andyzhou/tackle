-- base on sqlite3
-- db name: video2gif.db

-- files table
DROP TABLE IF EXISTS `files`;

CREATE TABLE `files` (
shortUrl varchar(64),
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
CREATE UNIQUE INDEX files_shortUrl_index ON files(shortUrl);
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