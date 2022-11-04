--dbscale check table—数据库名，分片表名均带中横线，修改par_0和par_1中部分shard的表结构 (part3-4)
--登录part_0 master，对shard（test-t-db2_dbscale_1_4）中表结构进行修改：mysql -uroot -p123456 -P16315 -h192.168.2.21
sql:alter table `test-t-db2_dbscale_1_4`.`test-t5` change  COLUMN name name_b varchar(20);
sql:show create table `test-t-db2_dbscale_1_4`.`test-t5`;




