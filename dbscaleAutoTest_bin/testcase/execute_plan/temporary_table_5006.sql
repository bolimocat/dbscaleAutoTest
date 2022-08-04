--临时表列字段包含数类型：主键自增列、唯一约束、备注内容
sql:create database testdb;
sql:use testdb
sql:CREATE TABLE testdb.tb_base(a int,b varchar(4));
sql:insert into testdb.tb_base values(1,'a');
sql:select * from testdb.tb_base;

--主键及自增列：
sql:CREATE TEMPORARY TABLE testdb.tb_aticremt(c1 int primary key auto_increment,c2 int);insert into testdb.tb_aticremt(c2)values(10);insert into testdb.tb_aticremt(c2)values(20);insert into testdb.tb_aticremt(c2)values(30);select * from testdb.tb_aticremt;

--唯一约束：
sql:CREATE TEMPORARY TABLE testdb.tb_unique(c1 int unique);insert into testdb.tb_unique values (1);select * from testdb.tb_unique;insert into testdb.tb_unique values (1);select * from testdb.tb_unique;

--备注字符内容：
sql:CREATE TEMPORARY TABLE testdb.tb_comment(c1 int comment "its comment of column");show full columns from testdb.tb_comment;

--清理环境：
sql:drop database testdb;