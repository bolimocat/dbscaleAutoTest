--dbscale check table—数据库名，分片表名均带中横线，修改par_0和par_1中部分shard的表结构 (part1-4)
--创建一个带中横线的数据库：
sql:create database `test-t-db2`;

--创建一个不带中横线的分片表：
sql:dbscale dynamic add hash_type partition_scheme "hash_test" partition="part_0" partition="part_1" partition="part_2" is_shard shard_nums 15;
sql:dbscale dynamic add partition_table dataspace `test-t-db2`.`test-t5` partition_key="id" partition_scheme="hash_test";
sql:create table `test-t-db2`.`test-t5` (id int primary key, name varchar(10), time TIMESTAMP);
sql:dbscale show table location `test-t-db2`.`test-t5`;

--执行dbscale check table命令：
sql:dbscale check table `test-t-db2`.`test-t5`;


