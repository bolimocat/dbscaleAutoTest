--建库
sql:create database mytest;
--建表
sql:create table mytest.tb(id int primary key auto_increment,value varchar(6));
--查表
sql:select * from mytest.tb;
--插入
sql:insert into mytest.tb(value) values ("my");
sql:insert into mytest.tb(value) values ("your");
--查表
sql:select * from mytest.tb;
--更新
sql:update mytest.tb set value = "Tana" where id = 2;
--删除
sql:delete from mytest.tb where id = 1;
--查表
sql:select * from mytest.tb;
--恢复环境
sql:drop database mytest;
