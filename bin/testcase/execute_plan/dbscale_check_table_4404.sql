-- dbscale check table—数据库名不带中横线，普通表带中横线 
--创建一个带中横线的数据库：
sql:create database db1;

--创建一个不带中横线的分片表：
sql:dbscale dynamic add normal_table dataspace db1.`test-t1` datasource = "normal_0";
sql:create table db1.`test-t1` (id int primary key, name varchar(10), time TIMESTAMP);


--执行dbscale check table命令：
sql:dbscale check table db1.`test-t1`;

--clean
sql:drop database db1;
sql:dbscale dynamic remove table db1.`test-t1`;
sql:dbscale dynamic remove schema db1;
