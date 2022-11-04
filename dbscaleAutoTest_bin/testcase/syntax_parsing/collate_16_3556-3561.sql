--前置条件
sql:dbscale set global "use-partial-parse" = 1;

--procedure ，输入参数varchar字段，编码gbk，collate使用gbk_bin 3556 - 3561
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
--CREATE PROCEDURE test1.test_procedure (IN p_in varchar(12) character set gbk  collate gbk_bin) BEGIN  select distinct c1 from  test1.tb_varchar_gbk_bin; select id,c1 from  test1.tb_varchar_gbk_bin order by c1; select c1 from  test1.tb_varchar_gbk_bin group by c1; select id,c1 from  test1.tb_varchar_gbk_bin having id < 8; select id,c1 from  test1.tb_varchar_gbk_bin where id > 5;END //
--DELIMITER ;
--call test1.test_procedure("a");
--show CREATE PROCEDURE test1.test_procedure\G;
--drop PROCEDURE test1.test_procedure;
sql:drop database test1;

--procedure ，输入参数text字段，编码gbk，collate使用gbk_chinese_ci 
sql:create database test1;
sql:use test1;
sql:create table test1.tb_text_gbk_chinese_ci (id int primary key auto_increment,c1 text character set gbk  collate gbk_chinese_ci);
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("一");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("六");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("七");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("八");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("四");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("五");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("六");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("七");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("八");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("二");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("三");
sql:insert into test1.tb_text_gbk_chinese_ci (c1) values ("四");
--DELIMITER // CREATE PROCEDURE test1.test_procedure (IN p_in text character set gbk  collate gbk_chinese_ci) BEGIN  select distinct c1 from  test1.tb_text_gbk_chinese_ci; select id,c1 from test1.tb_text_gbk_chinese_ci order by c1; select c1 from  test1.tb_text_gbk_chinese_ci group by c1; select id,c1 from  test1.tb_text_gbk_chinese_ci having id < 8; select id,c1 from  test1.tb_text_gbk_chinese_ci where id > 5; END //
--DELIMITER ;
--call test1.test_procedure("a");
--show CREATE PROCEDURE test1.test_procedure\G;
--drop procedure test1.test_procedure;
sql:drop database test1;

--procedure ，输入参数text字段，编码gbk，collate使用gbk_bin 
sql:create database test1;
sql:use test1;
sql:create table test1.tb_text_gbk_bin (id int primary key auto_increment,c1 text character set gbk  collate gbk_bin);
sql:insert into test1.tb_text_gbk_bin (c1) values ("一");
sql:insert into test1.tb_text_gbk_bin (c1) values ("二");
sql:insert into test1.tb_text_gbk_bin (c1) values ("三");
sql:insert into test1.tb_text_gbk_bin (c1) values ("四");
sql:insert into test1.tb_text_gbk_bin (c1) values ("五");
sql:insert into test1.tb_text_gbk_bin (c1) values ("六");
sql:insert into test1.tb_text_gbk_bin (c1) values ("七");
sql:insert into test1.tb_text_gbk_bin (c1) values ("二");
sql:insert into test1.tb_text_gbk_bin (c1) values ("三");
sql:insert into test1.tb_text_gbk_bin (c1) values ("四");
sql:insert into test1.tb_text_gbk_bin (c1) values ("五");
sql:insert into test1.tb_text_gbk_bin (c1) values ("八");
sql:insert into test1.tb_text_gbk_bin (c1) values ("二");
sql:insert into test1.tb_text_gbk_bin (c1) values ("三");
sql:insert into test1.tb_text_gbk_bin (c1) values ("四");
sql:insert into test1.tb_text_gbk_bin (c1) values ("五");
sql:insert into test1.tb_text_gbk_bin (c1) values ("六");
sql:insert into test1.tb_text_gbk_bin (c1) values ("七");
sql:insert into test1.tb_text_gbk_bin (c1) values ("八");
sql:insert into test1.tb_text_gbk_bin (c1) values ("二");
sql:insert into test1.tb_text_gbk_bin (c1) values ("三");
sql:insert into test1.tb_text_gbk_bin (c1) values ("四");
--DELIMITER // CREATE PROCEDURE test1.test_procedure (IN p_in text character set gbk  collate gbk_bin) BEGIN select distinct c1 from  test1.tb_text_gbk_bin; select id,c1 from  test1.tb_text_gbk_bin order by c1; select c1 from  test1.tb_text_gbk_bin group by c1; select id,c1 from  test1.tb_text_gbk_bin having id < 8; select id,c1 from  test1.tb_text_gbk_bin where id > 5;END //
--DELIMITER ;
--call test1.test_procedure("a");
--show CREATE PROCEDURE test1.test_procedure\G;
--drop PROCEDURE test1.test_procedure;
sql:drop database test1;

--function ，输入参数char字段，编码gbk，collate使用gbk_bin 
sql:create database test1;
sql:use test1;
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
--DELIMITER // CREATE FUNCTION test_function(p_in char(12) character set gbk  collate gbk_bin) RETURNS char(6)  character set gbk  collate gbk_bin BEGIN return "a";END //
--DELIMITER ;
--select test_function("a");
--show create function test_function\G;
--drop function test_function;
sql:drop database test1;

--function ，输入参数char字段，编码gbk，collate使用gbk_chinese_ci
sql:create database test1;
sql:use test1;
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
--DELIMITER //
--CREATE FUNCTION test_function (p_in char(12) character set gbk  collate gbk_chinese_ci)  RETURNS char(12) character set gbk  collate gbk_chinese_ci BEGIN return "0"; END //
--DELIMITER ;
--select test_function("a");
--show CREATE FUNCTION test_function\G;
--drop function test_function;
sql:drop database test1;

--procedure ，输入参数text字段，编码utf8mb4，collate使用utf8mb4_bin
sql:create database test1;
sql:use test1;
sql:create table test1.tb_text_utf8mb4_bin (id int primary key auto_increment,c1 text character set utf8mb4  collate utf8mb4_bin);
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("a");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("f");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("g");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("h");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("f");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("g");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("h");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_text_utf8mb4_bin (c1) values ("d");
--DELIMITER //
--CREATE PROCEDURE test1.test_procedure (IN p_in text character set utf8mb4  collate utf8mb4_bin) BEGIN  select distinct c1 from  test1.tb_text_utf8mb4_bin; select id,c1 from  test1.tb_text_utf8mb4_bin order by c1;select c1 from  test1.tb_text_utf8mb4_bin group by c1;  select id,c1 from  test1.tb_text_utf8mb4_bin having id < 8;  select id,c1 from  test1.tb_text_utf8mb4_bin where id > 5; END //
--DELIMITER ;
--CALL test1.test_procedure("A");
--SHOW CREATE PROCEDURE test1.test_procedure\G;
--drop procedure test1.test_procedure;
sql:drop database test1;
