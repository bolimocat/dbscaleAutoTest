--在客户端使用insert语句产生的warning提示，支持limit限制显示数量。
--创建测试库和测试表
sql:create database testdb;
sql:CREATE TABLE testdb.tb2 (id int(11) NOT NULL,name varchar(6) DEFAULT NULL,age int(11) DEFAULT NULL,level varchar(8) NOT NULL,PRIMARY KEY (id),UNIQUE KEY name (name)) ENGINE=InnoDB DEFAULT CHARSET=latin1;
sql:desc testdb.tb2;

sql:set sql_mode = "";source /tmp/import.sql;select * from testdb.tb2;

sql:drop database testdb;