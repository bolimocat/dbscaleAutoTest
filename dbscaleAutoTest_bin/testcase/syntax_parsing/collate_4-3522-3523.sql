--前置条件
sql:dbscale set global "use-partial-parse" = 1;

--create table ，varchar字段，编码utf8mb4，collate使用utf8mb4_bin 3522 
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
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("f");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("g");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("e");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("h");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("b");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("c");
sql:insert into test1.tb_varchar_utf8mb4_bin (c1) values ("d");
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_varchar_utf8mb4_bin";
sql:select distinct c1 from  test1.tb_varchar_utf8mb4_bin;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_bin order by c1;
sql:select c1 from  test1.tb_varchar_utf8mb4_bin group by c1;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_bin having id < 8;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_bin where id > 5;
sql:drop database test1;


--前置条件
sql:dbscale set global "use-partial-parse" = 1;

--create table ，varchar字段，编码utf8mb4，collate使用utf8mb4_general_ci 3523 
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
sql:select TABLE_NAME,COLLATION_NAME,DATA_TYPE from information_schema.columns where table_schema = "test1" and table_name = "test1.tb_varchar_utf8mb4_general_ci";
sql:select distinct c1 from  test1.tb_varchar_utf8mb4_general_ci;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_general_ci order by c1;
sql:select c1 from  test1.tb_varchar_utf8mb4_general_ci group by c1;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_general_ci having id < 8;
sql:select id,c1 from  test1.tb_varchar_utf8mb4_general_ci where id > 5;
sql:drop database test1;
