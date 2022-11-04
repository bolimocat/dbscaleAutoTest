--在客户端使用insert语句产生的warning提示，支持limit限制显示数量。
--创建测试库和测试表
sql:create database testdb;
sql:CREATE TABLE testdb.tb1 (id int(11) NOT NULL,name varchar(6) DEFAULT NULL,age int(11) DEFAULT NULL,level varchar(8) NOT NULL,PRIMARY KEY (id),UNIQUE KEY name (name)) ENGINE=InnoDB DEFAULT CHARSET=latin1;
sql:desc testdb.tb1;

sql:set sql_mode = "";insert into testdb.tb1 (id,name,age) values (1,"Tana111","5"),(2,"Tana222","5"),(3,"Tana333","5"),(4,"Tana444","5");show warnings;show warnings limit 3;show warnings limit 2,4;show warnings limit -2;show warnings limit a;show warnings limit 2,1;show warnings limit a,b;

sql:drop database testdb;