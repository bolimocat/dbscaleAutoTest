--临时表列字段包含数类型：非空值、默认字符或表达式
sql:create database testdb;
sql:use testdb
sql:CREATE TABLE testdb.tb_base(a int,b varchar(4));
sql:insert into testdb.tb_base values(1,'a');
sql:select * from testdb.tb_base;

--非空约束：
sql:CREATE TEMPORARY TABLE testdb.tb_notnull(c1 int,c2 int not null);insert into testdb.tb_notnull (c1) values (1);insert into testdb.tb_notnull (c2) values (2);select * from testdb.tb_notnull;

--字段默认值：
sql:CREATE TEMPORARY TABLE testdb.tb_default(c1 int default 1,c2 int);insert into testdb.tb_default(c2)values(2);select * from testdb.tb_default;

--清理环境：
sql:drop database testdb;