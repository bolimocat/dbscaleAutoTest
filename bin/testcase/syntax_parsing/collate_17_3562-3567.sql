--前置条件
sql:dbscale set global "use-partial-parse" = 1;

--procedure ，输入参数text字段，编码utf8mb4，collate使用utf8mb4_general_ci 3562 - 3567
sql:create database test1;
sql:use test1;
sql:create table test1.tb_text_utf8mb4_general_ci (id int primary key auto_increment,c1 text character set utf8mb4  collate utf8mb4_general_ci);
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("A");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("B");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("C");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("D");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("E");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("F");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("G");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("B");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("C");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("D");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("E");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("H");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("B");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("C");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("D");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("E");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("F");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("G");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("H");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("B");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("C");
sql:insert into test1.tb_text_utf8mb4_general_ci (c1) values ("D");
--DELIMITER //
--CREATE PROCEDURE test_procedure (IN p_in text character set utf8mb4  collate utf8mb4_general_ci) BEGIN  select distinct c1 from  test1.tb_text_utf8mb4_general_ci; select id,c1 from test1.tb_text_utf8mb4_general_ci order by c1; select c1 from  test1.tb_text_utf8mb4_general_ci group by c1; select id,c1 from  test1.tb_text_utf8mb4_general_ci having id < 8; select id,c1 from  test1.tb_text_utf8mb4_general_ci where id > 5; END //
--DELIMITER ;
--CALL test_procedure("A");
--SHOW CREATE PROCEDURE test_procedure\G;
--drop procedure test_procedure;
sql:drop database test1;

--procedure ，输入参数varchar字段，编码utf8mb4，collate使用utf8mb4_bin 
sql:create database test1;
sql:use test1;
sql:create table test1.tb_varchar_utf8mb4_bin (id int primary key auto_increment,c1 varchar(12) character set utf8mb4  collate utf8mb4_bin);
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("a");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("f");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("g");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("h");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("f");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("g");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("h");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("d");
--DELIMITER //
--CREATE PROCEDURE test_procedure (IN p_in varchar(12) character set utf8mb4  collate utf8mb4_bin) BEGIN select distinct c1 from  test1.tb_varchar_utf8mb4_bin; select id,c1 from  test1.tb_varchar_utf8mb4_bin order by c1;  select c1 from  test1.tb_varchar_utf8mb4_bin group by c1; select id,c1 from  test1.tb_varchar_utf8mb4_bin having id < 8; select id,c1 from  test1.tb_varchar_utf8mb4_bin where id > 5;END //
--DELIMITER ;
--CALL test_procedure("A");
--SHOW CREATE PROCEDURE test_procedure\G;
--drop procedure test_procedure;
sql:drop database test1;

--procedure ，输入参数varchar字段，编码utf8mb4，collate使用utf8mb4_general_ci 
sql:create database test1;
sql:use test1;
sql:create table test1.tb_varchar_utf8mb4_general_ci (id int primary key auto_increment,c1 varchar(12) character set utf8mb4  collate utf8mb4_general_ci);
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("a");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("d");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("e");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("f");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("g");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("d");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("e");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("h");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("d");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("e");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("f");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("g");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("h");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_general_ci (c1) values ("d");
--DELIMITER //
--CREATE PROCEDURE test_procedure (IN p_in varchar(12) character set utf8mb4  collate utf8mb4_general_ci) BEGIN select distinct c1 from  test1.tb_varchar_utf8mb4_general_ci; select id,c1 from  test1.tb_varchar_utf8mb4_general_ci order by c1; select c1 from  test1.tb_varchar_utf8mb4_general_ci group by c1; select id,c1 from  test1.tb_varchar_utf8mb4_general_ci having id < 8; select id,c1 from  test1.tb_varchar_utf8mb4_general_ci where id > 5; END //
--DELIMITER ;
--call test_procedure("a");
--show CREATE PROCEDURE test_procedure\G;
--drop procedure test_procedure;
sql:drop database test1;

--procedure ，输入参数char字段，编码utf8mb4，collate使用utf8mb4_general_ci 
sql:create database test1;
sql:use test1;
sql:create table test1.tb_char_utf8mb4_general_ci (id int primary key auto_increment,c1 char(12) character set utf8mb4  collate utf8mb4_general_ci);
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("a");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("b");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("c");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("d");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("e");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("f");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("g");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("b");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("c");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("d");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("e");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("h");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("b");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("c");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("d");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("e");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("f");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("g");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("h");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("b");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("c");
sql:insert into test1.tb_char_utf8mb4_general_ci (c1) values ("d");
--DELIMITER //
--CREATE PROCEDURE test_procedure (IN p_in char(12) character set utf8mb4  collate utf8mb4_general_ci) BEGIN  select distinct c1 from  test1.tb_char_utf8mb4_general_ci; select id,c1 from  test1.tb_char_utf8mb4_general_ci order by c1; select c1 from  test1.tb_char_utf8mb4_general_ci group by c1; select id,c1 from  test1.tb_char_utf8mb4_general_ci having id < 8; select id,c1 from  test1.tb_char_utf8mb4_general_ci where id > 5;END //
--DELIMITER ;
--call test_procedure("a");
--show CREATE PROCEDURE test_procedure\G;
--drop procedure test_procedure;
sql:drop database test1;

--procedure ，输入参数char字段，编码utf8mb4，collate使用utf8mb4_bin
sql:create database test1;
sql:use test1;
sql:create table test1.tb_char_utf8mb4_bin (id int primary key auto_increment,c1 char(12) character set utf8mb4  collate utf8mb4_bin);
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("a");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("f");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("g");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("h");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("d");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("f");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("g");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("h");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_char_utf8mb4_bin (c1) values ("d");
--DELIMITER //
--CREATE PROCEDURE test_procedure (IN p_in char(12) character set utf8mb4  collate utf8mb4_bin) BEGIN  select distinct c1 from  test1.tb_char_utf8mb4_bin; select id,c1 from  test1.tb_char_utf8mb4_bin order by c1; select c1 from  test1.tb_char_utf8mb4_bin group by c1; select id,c1 from  test1.tb_char_utf8mb4_bin having id < 8; select id,c1 from  test1.tb_char_utf8mb4_bin where id > 5;END //
--DELIMITER ;
--call test_procedure("a");
--show CREATE PROCEDURE test_procedure;
--drop procedure test_procedure;
sql:drop database test1;

--function ，输入参数varchar字段，编码gbk，collate使用gbk_chinese_ci 
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
--DELIMITER //
--CREATE FUNCTION test_function (p_in varchar(12) character set gbk  collate gbk_chinese_ci)  RETURNS varchar(12) character set gbk  collate gbk_chinese_ci BEGIN return "a"; END //
--DELIMITER ;
--select test_function("a");
--show CREATE FUNCTION test_function \G;
--drop function test_function;
sql:drop database test1;
