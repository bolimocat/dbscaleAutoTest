--在客户端使用loadfile语句产生的warning提示，支持limit限制显示数量。
--创建测试库和测试表
sql:create database testdb;
sql:CREATE TABLE testdb.tb3 (id int(11) NOT NULL,name varchar(6) DEFAULT NULL,age int(11) DEFAULT NULL,level varchar(8) NOT NULL,PRIMARY KEY (id),UNIQUE KEY name (name)) ENGINE=InnoDB DEFAULT CHARSET=latin1;
sql:desc testdb.tb3;

sql:set sql_mode = "";LOAD DATA LOCAL INFILE "/home/greatdb//dbscaleAutoTest/file/syntax_parsing/loadfile" INTO TABLE testdb.tb3 (id, name, age);show warnings;show warnings limit 3;show warnings limit 2,4;show warnings limit -2;show warnings limit a;show warnings limit 2,1;show warnings limit a,b;

sql:drop database testdb;