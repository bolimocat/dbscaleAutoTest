--创建测试库
create database test;
use test

--查询强制使用 整型 单字段主键
create table test.tb_intkey(id int primary key,name varchar(6));
insert into test.tb_intkey (id,name) values (1,"Tana");
insert into test.tb_intkey (id,name) values (2,"Nicuhe");
insert into test.tb_intkey (id,name) values (3,"Tom");
insert into test.tb_intkey (id,name) values (4,"Jerry");
insert into test.tb_intkey (id,name) values (5,"Gofe");
select id,name from test.tb_intkey;
select id,name from test.tb_intkey force index(PRI);
select count(*) from ( select id,name from test.tb_intkey FORCE INDEX( pri ) order by id DESC limit 5) test.tb_intkey where (id = 3);
drop table test.tb_intkey;

--查询强制使用 浮点 单字段主键
create table test.tb_floatkey(id float(2,1) primary key,name varchar(6));
insert into test.tb_floatkey (id,name) values (1.1,"Tana");
insert into test.tb_floatkey (id,name) values (2.1,"Nicuhe");
insert into test.tb_floatkey (id,name) values (3.1,"Tom");
insert into test.tb_floatkey (id,name) values (4.1,"Jerry");
insert into test.tb_floatkey (id,name) values (5.1,"Gofe");
select id,name from test.tb_floatkey;
select id,name from test.tb_floatkey force index(PRI);
select count(*) from ( select name from test.tb_floatkey FORCE INDEX( pri ) order by id DESC limit 4) test.tb_floatkey where (id = 3.1);
drop table test.tb_floatkey;

--查询强制使用 字符 单字段主键
create table test.tb_charkey(id int,name varchar(6)  primary key);
insert into test.tb_charkey (id,name) values (1,"Tana");
insert into test.tb_charkey (id,name) values (2,"Nicuhe");
insert into test.tb_charkey (id,name) values (3,"Tom");
insert into test.tb_charkey (id,name) values (4,"Jerry");
insert into test.tb_charkey (id,name) values (5,"Gofe");
select id,name from test.tb_charkey;
select id,name from test.tb_charkey force index(PRI);
select count(*) from ( select name from test.tb_charkey FORCE INDEX( pri ) order by id DESC limit 4) test.tb_charkey where (name = "Nicuhe");
drop table test.tb_charkey;

--查询强制使用 整型，浮点，字符 多字段联合主键
create table test.tb_unionkey(id int,weight float(2,1) ,name varchar(6),content text , primary key(id,weight,name));
insert into test.tb_unionkey (id,weight,name,content) values (1,2.2,"Tana","daughter");
insert into test.tb_unionkey (id,weight,name,content) values (2,1.0,"Nicuhe","son");
insert into test.tb_unionkey (id,weight,name,content) values (3,0.2,"Tom","my cat");
insert into test.tb_unionkey (id,weight,name,content) values (4,0.1,"Jerry","a rat");
insert into test.tb_unionkey (id,weight,name,content) values (5,9.9,"Gofe","puppy");
select id,name from test.tb_unionkey;
select id,name from test.tb_unionkey force index(PRI);
select count(*) from ( select id,weight,name,content from test.tb_unionkey FORCE INDEX( pri ) order by id DESC limit 4) test.tb_unionkey where (content = "son");
drop table test.tb_unionkey;

--clean database
drop database test;

