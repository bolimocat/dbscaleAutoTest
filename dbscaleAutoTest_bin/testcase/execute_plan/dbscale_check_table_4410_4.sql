--dbscale check table—数据库名，分片表名均带中横线，修改par_0和par_1中部分shard的表结构 (part4-4)
--登录dbscale，执行dbscale check table：
sql:dbscale check table `test-t-db2`.`test-t5`;

--clean
sql:drop table `test-t-db2`.`test-t5`;
sql:drop database `test-t-db2`;
sql:dbscale dynamic remove table `test-t-db2`.`test-t5`;
sql:dbscale dynamic remove partition_scheme hash_test;
sql:dbscale dynamic remove schema `test-t-db2`;
