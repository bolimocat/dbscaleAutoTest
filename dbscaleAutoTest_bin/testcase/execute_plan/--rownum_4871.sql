--  rownum—其他不支持rownum语法测试（单机单张普通表）
--前置条件:
sql:create database rownum_test;
sql:use rownum_test;
sql:dbscale dynamic add normal_table dataspace rownum_test.tb_1 datasource = "normal_0";
sql:dbscale dynamic add normal_table dataspace rownum_test.tb_2 datasource = "normal_0";
sql:create table rownum_test.tb_1 (tb_1_id int primary key ,tb_1_name char(10),tb_1_part char(10),achievement int) DEFAULT CHARSET=gbk; 
sql:create table rownum_test.tb_2 (tb_2_id int primary key ,tb_2_name char(10),tb_2_part char(10),achievement int) DEFAULT CHARSET=gbk;
sql:insert into rownum_test.tb_1 values (1,'老张','A CLASS',80);
sql:insert into rownum_test.tb_1 values (2,'老王','B CLASS',100);
sql:insert into rownum_test.tb_1 values (3,'老李','C CLASS',90);
sql:insert into rownum_test.tb_1 values (4,'老吴','A CLASS',70);
sql:insert into rownum_test.tb_1 values (5,'老郑','B CLASS',89);
sql:insert into rownum_test.tb_1 values (6,'老何','C CLASS',95);
sql:insert into rownum_test.tb_1 values (7,'老张','D CLASS',66);
sql:insert into rownum_test.tb_1 values (8,'老吴','E CLASS',75);
sql:insert into rownum_test.tb_2 values (1,'老王','A CLASS',90);
sql:insert into rownum_test.tb_2 values (2,'老刘','B CLASS',89);
sql:insert into rownum_test.tb_2 values (3,'老张','C CLASS',66);
sql:insert into rownum_test.tb_2 values (4,'老吴','D CLASS',78);
sql:insert into rownum_test.tb_2 values (5,'老李','E CLASS',85);
sql:insert into rownum_test.tb_2 values (6,'老王','F CLASS',93);
sql:insert into rownum_test.tb_2 values (7,'老孟','I CLASS',80);
sql:insert into rownum_test.tb_2 values (8,'老田','J CLASS',79);

--rownum在select的非where位置出现
sql:select tb_1_name,rownum from rownum_test.tb_1 where rownum <= 6;

--rownum和limit共用1
sql:select * from rownum_test.tb_1 where rownum <= 6 limit 4;

--rownum和in共用1
sql:select * from rownum_test.tb_1 where rownum in (1,2);

--rownum和in共用2
sql:select * from rownum_test.tb_1 where rownum<=6 and tb_1_id in (1,2);

--rownum和in共用3
sql:select * from rownum_test.tb_1 where tb_1_id in (select tb_2_id from rownum_test.tb_2 where rownum <=4);

--rownum和all共用1
sql:select * from rownum_test.tb_1 where rownum <=6 and achievement > all(select achievement from rownum_test.tb_2);

--rownum和all共用2
sql:select * from rownum_test.tb_1 where achievement > all(select achievement from rownum_test.tb_2 where rownum <=6);

--rownum和any共用1
sql:select * from rownum_test.tb_1 where rownum <=6 and achievement > any(select achievement from rownum_test.tb_2);

--rownum和any共用2
sql:select * from rownum_test.tb_1 where  achievement > any(select achievement from rownum_test.tb_2 where rownum <=6);

--rownum和any共用3
sql:select * from rownum_test.tb_1 where achievement > any(select achievement from rownum_test.tb_2) and rownum <=6;

--rownum和some共用1
sql:select * from rownum_test.tb_1 where rownum <=6 and achievement > some(select achievement from rownum_test.tb_2);

--rownum和some共用2
sql:select * from rownum_test.tb_1 where  achievement > some(select achievement from rownum_test.tb_2 where rownum <=6);

--clean
sql:drop table rownum_test.tb_1;
sql:drop table rownum_test.tb_2;
sql:drop database rownum_test;
sql:dbscale dynamic remove table rownum_test.tb_1;
sql:dbscale dynamic remove table rownum_test.tb_2;
sql:dbscale dynamic remove schema rownum_test;
