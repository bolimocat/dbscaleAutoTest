--创建测试库
sql:create database test;
sql:use test

--查询强制使用 整型 单字段主键
sql:create table test.tb_intkey(id int primary key,name varchar(6));
sql:insert into test.tb_intkey (id,name) values (1,"Tana");
sql:insert into test.tb_intkey (id,name) values (2,"Nicuhe");
sql:insert into test.tb_intkey (id,name) values (3,"Tom");
sql:insert into test.tb_intkey (id,name) values (4,"Jerry");
sql:insert into test.tb_intkey (id,name) values (5,"Gofe");
sql:select id,name from test.tb_intkey;
sql:select id,name from test.tb_intkey force index(PRI);
sql:select count(*) from ( select id,name from test.tb_intkey FORCE INDEX( pri ) order by id DESC limit 5) tb_intkey where (id = 3);
sql:drop table test.tb_intkey;

--查询强制使用 浮点 单字段主键
sql:create table test.tb_floatkey(id float(2,1) primary key,name varchar(6));
sql:insert into test.tb_floatkey (id,name) values (1.1,"Tana");
sql:insert into test.tb_floatkey (id,name) values (2.1,"Nicuhe");
sql:insert into test.tb_floatkey (id,name) values (3.1,"Tom");
sql:insert into test.tb_floatkey (id,name) values (4.1,"Jerry");
sql:insert into test.tb_floatkey (id,name) values (5.1,"Gofe");
sql:select id,name from test.tb_floatkey;
sql:select id,name from test.tb_floatkey force index(PRI);
sql:select count(*) from ( select id,name from test.tb_floatkey FORCE INDEX( pri ) order by id DESC limit 4) tb_floatkey where (id = 3.1);
sql:drop table test.tb_floatkey;

--查询强制使用 字符 单字段主键
sql:create table test.tb_charkey(id int,name varchar(6)  primary key);
sql:insert into test.tb_charkey (id,name) values (1,"Tana");
sql:insert into test.tb_charkey (id,name) values (2,"Nicuhe");
sql:insert into test.tb_charkey (id,name) values (3,"Tom");
sql:insert into test.tb_charkey (id,name) values (4,"Jerry");
sql:insert into test.tb_charkey (id,name) values (5,"Gofe");
sql:select id,name from test.tb_charkey;
sql:select id,name from test.tb_charkey force index(PRI);
sql:select count(*) from ( select name from test.tb_charkey FORCE INDEX( pri ) order by id DESC limit 4) tb_charkey where (name = "Nicuhe");
sql:drop table test.tb_charkey;

--查询强制使用 整型，浮点，字符 多字段联合主键
sql:create table test.tb_unionkey(id int,weight float(2,1) ,name varchar(6),content text , primary key(id,weight,name));
sql:insert into test.tb_unionkey (id,weight,name,content) values (1,2.2,"Tana","daughter");
sql:insert into test.tb_unionkey (id,weight,name,content) values (2,1.0,"Nicuhe","son");
sql:insert into test.tb_unionkey (id,weight,name,content) values (3,0.2,"Tom","my cat");
sql:insert into test.tb_unionkey (id,weight,name,content) values (4,0.1,"Jerry","a rat");
sql:insert into test.tb_unionkey (id,weight,name,content) values (5,9.9,"Gofe","puppy");
sql:select id,name from test.tb_unionkey;
sql:select id,name from test.tb_unionkey force index(PRI);
sql:select count(*) from ( select id,weight,name,content from test.tb_unionkey FORCE INDEX( pri ) order by id DESC limit 4) tb_unionkey where (content = "son");
sql:drop table test.tb_unionkey;

--clean database
sql:drop database test;

