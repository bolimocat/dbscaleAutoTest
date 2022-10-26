--前置条件
sql:dbscale set global "use-partial-parse" = 1;

--function ，输入参数varchar字段，编码gbk，collate使用gbk_bin 3568 - 3575  
sql:create database test1;
sql:use test1;
sql:create table test1.tb_varchar_gbk_bin (id int primary key auto_increment,c1 varchar(12) character set gbk  collate gbk_bin);
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("一");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("二");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("三");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("四");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("五");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("六");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("七");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("二");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("三");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("四");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("五");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("八");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("二");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("三");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("四");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("五");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("六");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("七");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("八");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("二");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("三");
sql:insert into test1.tb_varchar_gbk_bin (c1) values ("四");
--DELIMITER //
--CREATE FUNCTION test_function(p_in varchar(12) character set gbk  collate gbk_bin) RETURNS  varchar(12) character set gbk  collate gbk_bin BEGIN  return "0";END //
--DELIMITER ;
--select test_function("a");
--show CREATE FUNCTION test_function \G;
--drop function test_function;
sql:drop database test1;

--create table ，多char字段，同编码gbk，collate使用gbk_bin
sql:create database test1;
sql:use test1;
sql:create table test1.tb_char_gbk_bin (id int primary key auto_increment,c1 char(12) character set gbk  collate gbk_bin,c2 char(12) character set gbk  collate gbk_bin);
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("一","8");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("二","7");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("三","6");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("四","5");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("五","4");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("六","3");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("七","2");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("二","1");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("三","6");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("四","5");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("五","4");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("八","8");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("二","3");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("三","2");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("四","1");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("五","6");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("六","5");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("七","4");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("八","2");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("二","1");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("三","4");
sql:insert into test1.tb_char_gbk_bin (c1,c2) values ("四","5");
sql:select distinct c1 from  test1.tb_char_gbk_bin;
sql:select id,c1 from  test1.tb_char_gbk_bin order by c1;
sql:select c1 from  test1.tb_char_gbk_bin group by c1;
sql:select id,c1 from  test1.tb_char_gbk_bin having id < 8;
sql:select id,c1 from  test1.tb_char_gbk_bin where id > 5
sql:drop database test1;

--create table ，多varchar字段，同编码utf8mb4，collate使用utf8mb4_bin
sql:create database test1;
sql:use test1;
sql:create table test1.tb_varchar_utf8mb4_bin (id int primary key auto_increment,c1 char(12) character set utf8mb4  collate utf8mb4_bin,c2 char(12) character set utf8mb4  collate utf8mb4_bin);
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("一","8");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("二","7");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("三","6");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("四","5");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("五","4");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("六","3");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("七","2");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("二","1");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("三","6");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("四","5");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("五","4");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("八","8");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("二","3");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("三","2");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("四","1");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("五","6");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("六","5");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("七","4");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("八","2");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("二","1");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("三","4");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("四","5");
sql:select distinct c1 from  test1.tb_varchar_utf8mb4_bin;
sql:select id,c1,c2 from  test1.tb_varchar_utf8mb4_bin order by c2;
sql:select c1 from  test1.tb_varchar_utf8mb4_bin group by c1;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_bin having id < 8;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_bin where id > 5;
sql:drop database test1;

--create table ，多text字段，不同编码gbk，utf8mb4，collate使用gbk_chinese_ci，utf8mb4_general_ci 
sql:create database test1;
sql:use test1;
sql:create table test1.tb_varchar_utf8mb4_bin (id int primary key auto_increment,c1 char(12) character set gbk  collate gbk_chinese_ci,c2 char(12) character set utf8mb4  collate utf8mb4_general_ci);
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("一","8");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("二","7");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("三","6");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("四","5");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("五","4");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("六","3");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("七","2");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("二","1");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("三","6");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("四","5");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("五","4");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("八","8");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("二","3");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("三","2");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("四","1");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("五","6");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("六","5");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("七","4");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("八","2");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("二","1");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("三","4");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1,c2) values ("四","5");
sql:select distinct c1 from  test1.tb_varchar_utf8mb4_bin;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_bin order by c1;
sql:select c1 from  test1.tb_varchar_utf8mb4_bin group by c1;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_bin having id < 8;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_bin where id > 5;
sql:drop database test1;

--alter table ，添加新char字段，collate使用gbk_bin
sql:create database test1;
sql:use test1;
sql:create table test1.tb_alter (id int primary key auto_increment,c1 char(12) );
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "tb_alter";
sql:alter table test1.tb_alter add c2 char(6)  character set gbk collate gbk_bin;
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "tb_alter";
sql:drop database test1;

--alter table ，varchar字段变text字段，collate从gbk_bin到gbk_chinese_ci
sql:create database test1;
sql:use test1;
sql:create table test1.tb_alter (id int primary key auto_increment,c1 varchar(12) character set gbk  collate gbk_bin);
sql:insert into test1.tb_alter (c1) values ("一");
sql:insert into test1.tb_alter (c1) values ("二");
sql:insert into test1.tb_alter (c1) values ("三");
sql:insert into test1.tb_alter (c1) values ("四");
sql:insert into test1.tb_alter (c1) values ("五");
sql:insert into test1.tb_alter (c1) values ("六");
sql:insert into test1.tb_alter (c1) values ("七");
sql:insert into test1.tb_alter (c1) values ("二");
sql:insert into test1.tb_alter (c1) values ("三");
sql:insert into test1.tb_alter (c1) values ("四");
sql:insert into test1.tb_alter (c1) values ("五");
sql:insert into test1.tb_alter (c1) values ("八");
sql:insert into test1.tb_alter (c1) values ("二");
sql:insert into test1.tb_alter (c1) values ("三");
sql:insert into test1.tb_alter (c1) values ("四");
sql:insert into test1.tb_alter (c1) values ("五");
sql:insert into test1.tb_alter (c1) values ("六");
sql:insert into test1.tb_alter (c1) values ("七");
sql:insert into test1.tb_alter (c1) values ("八");
sql:insert into test1.tb_alter (c1) values ("二");
sql:insert into test1.tb_alter (c1) values ("三");
sql:insert into test1.tb_alter (c1) values ("四");
sql:select distinct c1 from  test1.tb_alter;
sql:select id,c1 from  test1.tb_alter order by c1;
sql:select c1 from  test1.tb_alter group by c1;
sql:select id,c1 from  test1.tb_alter having id < 8;
sql:select id,c1 from  test1.tb_alter where id > 5;
sql:alter table  test1.tb_alter modify  c1 text character set utf8mb4 collate utf8mb4_general_ci;
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "tb_alter" ;
sql:drop database test1;