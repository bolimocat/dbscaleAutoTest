--前置条件
sql:dbscale set global "use-partial-parse" = 1;

--alter table ，char字段，collate从无到有gbk_bin 3530 
sql:create database test1;
sql:use test1;
sql:create table test1.tb_alter (id int primary key auto_increment,c1 char(12) );
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_alter";
sql:alter table  test1.tb_alter  change c1  ca char(12) character set gbk collate gbk_bin;
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_alter";
sql:drop database test1;

--alter table ，char字段，collate从无到有gbk_chinese_ci 3531
sql:create database test1;
sql:use test1;
sql:create table test1.tb_alter (id int primary key auto_increment,c1 char(12) );
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_alter";
sql:alter table  test1.tb_alter change column c1 ca char(6) character set gbk collate gbk_chinese_ci;
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_alter";
sql:drop database test1;