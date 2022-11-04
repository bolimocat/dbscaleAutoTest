--临时表列字段包含数类型：Numeric、DataTime、String
sql:create database testdb;
sql:use testdb
sql:CREATE TABLE testdb.tb_base(a int,b varchar(4));
sql:insert into testdb.tb_base values(1,'a');
sql:select * from testdb.tb_base;

--numeric数据类型：
sql:CREATE TEMPORARY TABLE testdb.tb_numeric(c1 tinyint,c2 smallint,c3 mediumint,c4 int,c5 bigint,c6 float,c7 double,c8 decimal);show columns from testdb.tb_numeric;insert into testdb.tb_numeric values (1,2,3,4,5,6.1,7.2,8);select * from testdb.tb_numeric;

--DataTime数据类型：
sql:CREATE TEMPORARY TABLE testdb.tb_datetime(c1 date,c2 time,c3 year,c4 datetime,c5 timestamp);show columns from testdb.tb_datetime;
--引号错误insert into testdb.tb_datetime values('2021-3-22','17:34:00','2021','2021-3-22 17:44:00','2021-3-22 17:44:00');select * from testdb.tb_datetime;

--String数据类型：
sql:CREATE TEMPORARY TABLE testdb.tb_string(c1 char(2),c2 varchar(4),c3 tinyblob,c4 tinytext,c5 blob,c6 text,c7 mediumblob,c8 mediumtext,c9 longblob,c10 longtext);show columns from testdb.tb_string;

--清理环境：
sql:drop database testdb;