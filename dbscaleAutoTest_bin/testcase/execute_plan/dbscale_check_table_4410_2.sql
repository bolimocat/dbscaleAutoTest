--dbscale check table—数据库名，分片表名均带中横线，修改par_0和par_1中部分shard的表结构 (part2-4)
--登录part_0 master，对shard（test-t-db2、test-t-db2_dbscale_1_5）中表结构进行修改：mysql -uroot -p123456 -P16314 -h192.168.2.21
sql:alter table `test-t-db2`.`test-t5` change  COLUMN name name_a varchar(20);
sql:alter table `test-t-db2_dbscale_1_5`.`test-t5` change  COLUMN name name_a varchar(20);
sql:show create table `test-t-db2`.`test-t5`;
sql:show create table `test-t-db2_dbscale_1_5`.`test-t5`;
