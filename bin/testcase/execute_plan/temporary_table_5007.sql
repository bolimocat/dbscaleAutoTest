--临时表列字段列格式：DEFAULT DYNAMIC COMPRESSED REDUNDANT COMPACT
sql:create database testdb;
sql:use testdb
sql:CREATE TABLE testdb.tb_base(a int,b varchar(4));
sql:insert into testdb.tb_base values(1,'a');
sql:select * from testdb.tb_base;

--DEFAULT：
sql:CREATE TEMPORARY TABLE testdb.tb_def(c1 int primary key auto_increment) row_format = default;show create table testdb.tb_def;

--DYNAMIC
sql:CREATE TEMPORARY TABLE testdb.tb_dynamic(c1 int primary key auto_increment) row_format = dynamic;show create table testdb.tb_dynamic;

--COMPRESSED：
sql:CREATE TEMPORARY TABLE testdb.tb_compressed(c1 int primary key auto_increment) row_format = compressed;show create table testdb.tb_compressed;

--REDUNDANT：
sql:CREATE TEMPORARY TABLE testdb.tb_redundant(c1 int primary key auto_increment) row_format = redundant;show create table testdb.tb_redundant;

--COMPACT：
sql:CREATE TEMPORARY TABLE testdb.tb_compact(c1 int primary key auto_increment) row_format = compact;show create table testdb.tb_compact;

--清理环境：
sql:drop database testdb;