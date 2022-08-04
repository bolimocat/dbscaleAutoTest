--前置条件
sql:dbscale set global "use-partial-parse" = 1;
--alter table ，varchar字段，collate从无到有utf8mb4_bin 3538
sql:create database test1;
sql:use test1;
sql:create table test1.tb_alter (id int primary key auto_increment,c1 varchar(12) );
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_alter";
sql:alter table  test1.tb_alter change column c1 ca varchar(6) character set utf8mb4 collate utf8mb4_bin;
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_alter";
sql:drop database test1;

--alter table ，varchar字段，collate从无到有utf8mb4_general_ci 3539
sql:create database test1;
sql:use test1;
sql:create table test1.tb_alter(id int primary key auto_increment,c1 varchar(12) );
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_alter";
sql:alter table  test1.tb_alter change column c1 ca  varchar(6) character set utf8mb4 collate utf8mb4_general_ci;
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_alter";
sql:drop database test1;
