--临时表其他创建方式
sql:create database testdb;
sql:use testdb
sql:CREATE TABLE testdb.tb_base(a int,b varchar(4));
sql:insert into testdb.tb_base values(1,'a');
sql:select * from testdb.tb_base;

--从现有表中select字段：
sql:CREATE TEMPORARY TABLE testdb.tb_select as select b from testdb.tb_base;show full columns from testdb.tb_select;

--like一个现有的表：
sql:CREATE TEMPORARY TABLE testdb.tb_like like testdb.tb_base;show full columns from testdb.tb_like; 

--清理环境：
sql:drop database testdb;