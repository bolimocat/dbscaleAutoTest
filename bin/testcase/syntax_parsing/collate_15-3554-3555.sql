--前置条件
sql:dbscale set global "use-partial-parse" = 1;
--procedure ，输入参数char字段，编码gbk，collate使用gbk_chinese_ci 3554
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
--DELIMITER // CREATE PROCEDURE test1.test_procedure (IN p_in char(12) character set gbk  collate gbk_chinese_ci) BEGIN  select distinct c1 from  test1.tb_char_gbk_chinese_ci; select id,c1 from test1.tb_char_gbk_chinese_ci order by c1; select c1 from  test1.tb_char_gbk_chinese_ci group by c1; select id,c1 from  test1.tb_char_gbk_chinese_ci having id < 8; select id,c1 from  test1.tb_char_gbk_chinese_ci where id > 5;END //
--DELIMITER ;
--call test1.test_procedure("a");
--show CREATE PROCEDURE test1.test_procedure\G;
--drop procedure test1.test_procedure;
sql:drop database test1;

--procedure ，输入参数varchar字段，编码gbk，collate使用gbk_chinese_ci  3555
sql:create database test1;
sql:use test1
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
--CREATE PROCEDURE test1.test_procedure (IN p_in varchar(12) character set gbk  collate gbk_chinese_ci) BEGIN  select distinct c1 from  test1.tb_varchar_gbk_chinese_ci; select id,c1 from test1.tb_varchar_gbk_chinese_ci order by c1; select c1 from  test1.tb_varchar_gbk_chinese_ci group by c1; select id,c1 from  test1.tb_varchar_gbk_chinese_ci having id < 8; select id,c1 from test1.tb_varchar_gbk_chinese_ci where id > 5; END //
--DELIMITER ;
--call test1.test_procedure("a");
--show CREATE PROCEDURE test1.test_procedure\G;
--drop procedure test1.test_procedure;
sql:drop database test1;
