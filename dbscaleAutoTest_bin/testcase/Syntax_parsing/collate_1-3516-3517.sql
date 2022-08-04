--前置条件
sql:dbscale set global "use-partial-parse" = 1;

--create table ，char字段，编码gbk，collate使用gbk_bin   3516 
sql:create database test1;
sql:create table test1.tb_char_gbk_bin (id int primary key auto_increment,c1 char(12) character set gbk  collate gbk_bin);
sql:insert into test1.tb_char_gbk_bin (c1) values ("一");
sql:insert into test1.tb_char_gbk_bin (c1) values ("二");
sql:insert into test1.tb_char_gbk_bin (c1) values ("三");
sql:insert into test1.tb_char_gbk_bin (c1) values ("四");
sql:insert into test1.tb_char_gbk_bin (c1) values ("五");
sql:insert into test1.tb_char_gbk_bin (c1) values ("六");
sql:insert into test1.tb_char_gbk_bin (c1) values ("七");
sql:insert into test1.tb_char_gbk_bin (c1) values ("二");
sql:insert into test1.tb_char_gbk_bin (c1) values ("三");
sql:insert into test1.tb_char_gbk_bin (c1) values ("四");
sql:insert into test1.tb_char_gbk_bin (c1) values ("五");
sql:insert into test1.tb_char_gbk_bin (c1) values ("八");
sql:insert into test1.tb_char_gbk_bin (c1) values ("二");
sql:insert into test1.tb_char_gbk_bin (c1) values ("三");
sql:insert into test1.tb_char_gbk_bin (c1) values ("四");
sql:insert into test1.tb_char_gbk_bin (c1) values ("五");
sql:insert into test1.tb_char_gbk_bin (c1) values ("六");
sql:insert into test1.tb_char_gbk_bin (c1) values ("七");
sql:insert into test1.tb_char_gbk_bin (c1) values ("八");
sql:insert into test1.tb_char_gbk_bin (c1) values ("二");
sql:insert into test1.tb_char_gbk_bin (c1) values ("三");
sql:insert into test1.tb_char_gbk_bin (c1) values ("四");
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_char_gbk_bin";
sql:select distinct c1 from  test1.tb_char_gbk_bin;
sql:select id,c1 from  test1.tb_char_gbk_bin order by c1;
sql:select c1 from  test1.tb_char_gbk_bin group by c1;
sql:select id,c1 from  test1.tb_char_gbk_bin having id < 8;
sql:select id,c1 from  test1.tb_char_gbk_bin where id > 5;
sql:drop database test1;

--前置条件
sql:dbscale set global "use-partial-parse" = 1;

--create table ，char字段，编码gbk，collate使用gbk_chinese_ci 3517
sql:create database test1;
sql:create table test1.tb_char_gbk_chinese_ci (id int primary key auto_increment,c1 char(12) character set gbk  collate gbk_chinese_ci);
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("一");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("六");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("七");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("八");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("六");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("七");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("八");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_char_gbk_chinese_ci (c1) values ("四");
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_char_gbk_chinese_ci";
sql:select distinct c1 from  test1.tb_char_gbk_chinese_ci;
sql:select id,c1 from  test1.tb_char_gbk_chinese_ci order by c1;
sql:select c1 from  test1.tb_char_gbk_chinese_ci group by c1;
sql:select id,c1 from  test1.tb_char_gbk_chinese_ci having id < 8;
sql:select id,c1 from  test1.tb_char_gbk_chinese_ci where id > 5;
sql:drop database test1;