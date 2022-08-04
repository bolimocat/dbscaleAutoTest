-- dbscale check table—数据库名带中横线，分片表不带中横线
--创建一个带中横线的数据库：
sql:create database `test-t-db1`;

--创建一个不带中横线的分片表：
sql:dbscale dynamic add hash_type partition_scheme "hash_test" partition="part_0" partition="part_1" partition="part_2" is_shard shard_nums 15;
sql:dbscale dynamic add partition_table dataspace `test-t-db1`.table2 partition_key="id" partition_scheme="hash_test";
sql:create table `test-t-db1`.table2 (id int primary key, name varchar(10), time TIMESTAMP);
sql:dbscale show table location `test-t-db1`.table2;

--执行dbscale check table命令：
sql:dbscale check table `test-t-db1`.table2;

--clean
sql:drop table  `test-t-db1`.table2;
sql:drop database `test-t-db1`;
sql:dbscale dynamic remove table `test-t-db1`.table2;
sql:dbscale dynamic remove partition_scheme hash_test;
sql:dbscale dynamic remove schema `test-t-db1`;


