-- dbscale check table—数据库名、分片表名带*连接符
--创建一个带中横线的数据库：
sql:create database `test*db1`;

--创建一个不带中横线的分片表：
sql:dbscale dynamic add hash_type partition_scheme "hash_test" partition="part_0" partition="part_1" partition="part_2" is_shard shard_nums 15;
sql:dbscale dynamic add partition_table dataspace `test*db1`.`test*t1` partition_key="id" partition_scheme="hash_test";
sql:create table `test*db1`.`test*t1` (id int primary key, name varchar(10), time TIMESTAMP);
sql:dbscale show table location `test*db1`.`test*t1`;

--执行dbscale check table命令：
sql:dbscale check table `test*db1`.`test*t1`;

--clean
sql:drop table  `test*db1`.`test*t1`;
sql:drop database `test*db1`;
sql:dbscale dynamic remove table `test*db1`.`test*t1`;
sql:dbscale dynamic remove partition_scheme hash_test;
sql:dbscale dynamic remove schema `test*db1`;


