--建库
sql:create database mytest111;
--建表
sql:create table mytest111.tb(id int primary key auto_increment,value varchar(6));
--查表
sql:select * from mytest111.tb;
--插入
sql:insert into mytest111.tb(value) values ("my");
sql:insert into mytest111.tb(value) values ("your");
--查表
sql:select * from mytest111.tb;
--更新
sql:update mytest111.tb set value = "Tana" where id = 2;
--删除
sql:delete from mytest111.tb where id = 1;
--查表
sql:select * from mytest111.tb;
--恢复环境
sql:drop database mytest111;