--前置条件
sql:dbscale set global "use-partial-parse" = 1;
--create table ，varchar字段，编码gbk，collate使用gbk_bin 3520 
sql:create database test1;
sql:use test1;
sql:create table test1.tb_varchar_gbk_bin (id int primary key auto_increment,c1 varchar(12) character set gbk  collate gbk_bin);
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("a");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("b");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("c");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("d");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("e");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("f");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("g");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("b");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("c");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("d");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("e");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("f");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("b");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("c");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("d");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("e");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("f");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("g");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("h");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("b");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("c");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("d");
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_varchar_gbk_bin";
sql:select distinct c1 from  test1.tb_varchar_gbk_bin;
sql:select id,c1 from  test1.tb_varchar_gbk_bin order by c1;
sql:select c1 from  test1.tb_varchar_gbk_bin group by c1;
sql:select id,c1 from  test1.tb_varchar_gbk_bin having id < 8;
sql:select id,c1 from  test1.tb_varchar_gbk_bin where id > 5;
sql:drop database test1;

--create table ，varchar字段，编码gbk，collate使用gbk_chinese_ci 3521
sql:create database test1;
sql:use test1;
sql:create table test1.tb_varchar_gbk_chinese_ci (id int primary key auto_increment,c1 varchar(12) character set gbk  collate gbk_chinese_ci);
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("一");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("六");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("七");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("八");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("六");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("七");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("八");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_varchar_gbk_chinese_ci (c1) values ("四");
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_varchar_gbk_chinese_ci";
sql:select distinct c1 from  test1.tb_varchar_gbk_chinese_ci;
sql:select id,c1 from  test1.tb_varchar_gbk_chinese_ci order by c1;
sql:select c1 from  test1.tb_varchar_gbk_chinese_ci group by c1;
sql:select id,c1 from  test1.tb_varchar_gbk_chinese_ci having id < 8;
sql:select id,c1 from  test1.tb_varchar_gbk_chinese_ci where id > 5;
sql:drop database test1; 