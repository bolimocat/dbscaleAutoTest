--临时表字符CHARACTER和COLLATE
sql:create database testdb;
sql:use testdb
sql:CREATE TABLE testdb.tb_base(a int,b varchar(4));
sql:insert into testdb.tb_base values(1,'a');
sql:select * from testdb.tb_base;

--CHARACTER：
sql:CREATE TEMPORARY TABLE testdb.tb_charset(c1 int primary key,c2 varchar(2)) character set "gb2312";show full columns from testdb.tb_charset;

--COLLATE：
sql:CREATE TEMPORARY TABLE testdb.tb_collate(c1 int primary key,c2 varchar(2)) character set "utf8mb4" collate  "utf8mb4_unicode_ci"; show full columns from testdb.tb_collate;

--清理环境：
sql:drop database testdb;