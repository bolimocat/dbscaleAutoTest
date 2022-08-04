-- dbscale check table—数据库名带中横线，普通表带中横线
--创建一个带中横线的数据库：
sql:create database `test-t-db1`;

--创建一个不带中横线的普通表：
sql:dbscale dynamic add normal_table dataspace `test-t-db1`.`test-t3` datasource = "normal_0";
sql:create table `test-t-db1`.`test-t3` (id int primary key, name varchar(10), time TIMESTAMP);

--执行dbscale check table命令：
sql:dbscale check table `test-t-db1`.`test-t3`;

--clean
sql:drop database `test-t-db1`;
sql:dbscale dynamic remove table `test-t-db1`.`test-t3`;
sql:dbscale dynamic remove schema `test-t-db1`;


