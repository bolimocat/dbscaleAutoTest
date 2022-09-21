package main

import (
	"fmt"
//	"time"
	"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	 "database/sql"
	remote "dbscaleAutoTest/remote"		
	transmit "dbscaleAutoTest/transmit" 
	load "dbscaleAutoTest/load"
//	local "dbscaleAutoTest/local"
//	test "dbscaleAutoTest/test"
	base "dbscaleAutoTest/database"
	execute "dbscaleAutoTest/execute"
//	tool "dbscaleAutoTest/tool"

)

func main(){
	fmt.Println("Cluster 4.0 Auto Test suit")
	var scene_info []string = load.Loadscene("./testscene/sceneinfo.csv") //从sceneinfo.csv按行读取场景信息
	var attachment_info []string = load.Loadattachment("./testscene/sceneinfo.csv") //从sceneinfo.csv按行读取场景中需要传送附件的附件名称列表
	var properties_info []string = load.Loadproperties("./file/properties") //读取环境配置文件（主要是连接测试服务器的ssh相关信息）
	var dbscale_host string
	var all_host string
	var user string
	var password string
	var testroot string 
	var testpath string
	var rootpass string
//	var casefolder string
//	var caseinfo string
	var uif []string = make([]string,0)
	var casefilelist []string = make([]string,0) //以用例文件（scene.csv[8]）的分片集合
	var expressfilelist []string = make([]string,0) //以预期文件（scene.csv[9]）的分片集合	
	var dbpropertieslist []string = make([]string,0) //以测试分类（scene.csv[11]）的分片集合

	//按行读取file/properties文件中的内容，过滤获得被测服务器主机上的登录信息和目录位置信息。
	for index,value := range properties_info{
		if value != "" {
			if index == 3 {
				user = strings.Split(value, "=")[1]
				uif = append(uif,user)
			}
			if index == 4 {
				password = strings.Split(value, "=")[1]
				uif = append(uif,password)
			}
			if index == 5 {
				rootpass = strings.Split(value, "=")[1]
				uif = append(uif,rootpass)
			}
			if index == 6 {
				dbscale_host = strings.Split(value, "=")[1]
				uif = append(uif,dbscale_host)
			}
			if index == 7 {
				all_host = strings.Split(value, "=")[1]
				uif = append(uif,all_host)
			}
			if index == 8 {
				testroot = strings.Split(value, "=")[1]
				uif = append(uif,testroot)
			}
			if index == 9 {
				testpath = strings.Split(value, "=")[1]
				uif = append(uif,testpath)
			}
			
		}
	}
	
	//读取sceneinfo.csv，加载3个分片，分别为用例文件名，预期文件名，测试类型。
	for _,value := range scene_info{
		if value != "" {
			casefile_info := strings.Split(value, ",")
			casefilelist = append(casefilelist,casefile_info[7]+":"+casefile_info[8])
			expressfilelist = append(expressfilelist,casefile_info[7]+":"+casefile_info[9])
			dbpropertieslist = append(dbpropertieslist,casefile_info[11])
		}
		
	}
	
	//从properties中获得当前集群的所有节点，并在所有节点上完成准备工作。
		host := strings.Split(all_host, ",")
		for _,value := range host {
			//在每个dbscale的节点上创建测试目录，存放测试用例
			fmt.Println("dbscale的节点上创建需要的目录")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath)
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testcase")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testresult")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testexpect")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/interface")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testscript")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/file")
			//向每个dbscale所在的目录发送高可用测试需要的脚本
			transmit.Upload(user, password, value, 22, "./testscript/kill_pid.sh", testroot+testpath+"/testscript/","kill_pid.sh", "向 "+value+" 发送kill_pid脚本")
			transmit.Upload(user, password, value, 22, "./testscript/ifcontrol.sh", testroot+testpath+"/testscript/","ifcontrol.sh", "向 "+value+" 发送ifcontrol脚本")
			//修改脚本文件为可执行
			remote.Nodemission(user, password, value, 22, "chmod +x "+testroot+testpath+"/testscript/kill_pid.sh")
			remote.Nodemission(user, password, value, 22, "chmod +x "+testroot+testpath+"/testscript/ifcontrol.sh")
			//修改带有重启实例的脚本属主为普通用户
			remote.Nodemission(user, password, value, 22, "chown -R greatdb:greatdb "+testroot+testpath+"/testscript/kill_pid.sh")
			//将脚本放到目标节点的/usr/bin下，执行时不用设置环境变量
			remote.Nodemission("root", password, value, 22, "mv "+testroot+testpath+"/testscript/kill_pid.sh /usr/bin")
			remote.Nodemission("root", password, value, 22, "mv "+testroot+testpath+"/testscript/ifcontrol.sh /usr/bin")
			
			//从已有的用例分片列表，在每个dbscale所在节点上创建用例目录
			for _,v := range casefilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testcase/"+strings.Split(v, ":")[0])
			}
			//从已有的预期文件分片列表，在每个dbscale所在节点上创建预期文件目录
			for _,v := range expressfilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testexpect/"+strings.Split(v, ":")[0])
			}
			
			for _,v := range expressfilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testresult/"+strings.Split(v, ":")[0])
			}
			
			for _,v := range expressfilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/file/"+strings.Split(v, ":")[0])
			}
			
//			chown -R greatdb:greatdb
			for _,casename := range casefilelist {
				transmit.Upload(user, password, value, 22,"./testcase/"+strings.Split(casename, ":")[0]+"/"+strings.Split(casename, ":")[1],testroot+testpath+"/testcase/"+strings.Split(casename, ":")[0],"/"+strings.Split(casename, ":")[1],"向 "+value+" 发送casefile")
				
			}
			
			for _,expressname := range expressfilelist{
				transmit.Upload(user, password, value, 22,"./testexpect/"+strings.Split(expressname, ":")[0]+"/"+strings.Split(expressname, ":")[1],testroot+testpath+"/testexpect/"+strings.Split(expressname, ":")[0],"/"+strings.Split(expressname, ":")[1],"向 "+value+" 发送expressfile")
			}
			
			for _,attachname := range attachment_info{
				transmit.Upload(user, password, value, 22, "./file/"+strings.Split(attachname, ",")[7]+"/"+strings.Split(attachname, ",")[12], testroot+testpath+"/file/"+strings.Split(attachname, ",")[7],strings.Split(attachname, ",")[12], "向 "+value+" 发送attache")
				
			}
		}
	

	
	//读取场景信息
	for _,value := range scene_info{
		if value != "" {
			fmt.Println("执行场景：  ",value)
			exe_info := strings.Split(value, ",")
//			fmt.Println("exe_info[4] ="+exe_info[4])	//exe ip
//			fmt.Println("exe_info[2] ="+exe_info[2])	//user
//			fmt.Println("exe_info[3] ="+exe_info[3])	//pass
//			fmt.Println("exe_info[5] ="+exe_info[5])	//dbscale port
//			fmt.Println("exe_info[6] ="+exe_info[6])	//local path
//			fmt.Println("exe_info[7] ="+exe_info[7])	//case filt
//			fmt.Println("exe_info[8] ="+exe_info[8])	//case name
//			fmt.Println("exe_info[9] ="+exe_info[9])	//express name
//			fmt.Println("exe_info[10] ="+exe_info[10])	//active
//			fmt.Println("exe_info[11] ="+exe_info[11])	//properties

			//读取场景中的配置状态
			if exe_info[11] == "function" {
				execute.FunctionCase(exe_info[7], exe_info[8], exe_info[2], exe_info[4], exe_info[3], exe_info[4], exe_info[9], password, exe_info[0], exe_info[1],testroot,testpath)

			
			}
			if exe_info[11] == "interface" {
//				默认创建5个表，以此未为基础执行操作.
				execute.InterfaseCase(exe_info[4], exe_info[2], exe_info[3], exe_info[5], "5", exe_info[8], exe_info[7],exe_info[0],exe_info[1])
			}
			if exe_info[11] == "stable" {
				fmt.Println("按行执行stable用例")
			}
			if exe_info[11] == "stress" {
				fmt.Println("按行执行stress用例")
			}
			
			if exe_info[11] == "ifcontrol" {
				fmt.Println("按行执行断网用例")
				fmt.Println("执行主机 : "+exe_info[4])
//				ifcontrolIP := strings.Split(strings.Split(exe_info[8], ".")[3], "-")[0]+"."+strings.Split(exe_info[8], ".")[1]+"."+strings.Split(exe_info[8], ".")[2]+"."+strings.Split(exe_info[8], ".")[3]
				fmt.Println("查询主机 : "+exe_info[4])
//				conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",exe_info[2], "123456", "tcp", exe_info[4], 16310, "dbscale")	
				conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",exe_info[2], "123456", "tcp", exe_info[4], 16310, "dbscale")	
				db, err := sql.Open("mysql", conn)
			   defer db.Close()                                            //关闭数据库
			    err = db.Ping()                                            //连接数据库
			    if err != nil {
		        fmt.Println("数据库连接失败")                             //连接失败
		        return
			    } else {
			        fmt.Println("数据库连接成功")                             //连接成功
			     }
				
				var killedhost string
				var errnum int
				for i:=0;i<5;i++{
					if base.FetchDataserversInfo(db){//judge the dataserver is correct
						remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '第 "+strconv.Itoa(i)+" '轮测试  >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
						remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
//						remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
						
						if strings.Split(strings.Split(exe_info[8], ".")[0], "-")[2] == "master"{//if the node is master,so ifdown the interface on master host
							killedhost = base.Killdataserversmaster(db, strings.Split(strings.Split(exe_info[8], ".")[0], "-")[1])
							fmt.Println("全部实例状态恢复，断开端口"+strings.Split(strings.Split(exe_info[8], ".")[0], "-")[1]+"的master节点所在的网络。")
							remote.Nodemission("root", password, killedhost, 22, "ifcontrol.sh > /dev/null 2>&1 & " )
							
							caselist := load.Loadcase("./testcase/"+exe_info[7]+"/"+exe_info[8])
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
									remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
									remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -uroot -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result" +" 2>&1"  )
									remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
									//get the execute result from remote host by current build
									transmit.Download(exe_info[2],exe_info[4],testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result","./testcurrent/"+exe_info[7]);
								}
							}
						}
						if strings.Split(strings.Split(exe_info[8], ".")[0], "-")[2] == "slave"{
							killedhost = base.Killdataserverslastslave(db, strings.Split(strings.Split(exe_info[8], ".")[0], "-")[1])
							fmt.Println("全部实例状态恢复，杀死"+killedhost+"上，端口"+strings.Split(strings.Split(exe_info[8], ".")[0], "-")[1]+"的最后一个slave实例。")
							remote.Nodemission(user, password, killedhost, 22, "kill_pid.sh  15 "+strings.Split(strings.Split(exe_info[8], ".")[0], "-")[1] )
							remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
						
						caselist := load.Loadcase("./testcase/"+exe_info[7]+"/"+exe_info[8])
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
									remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
									remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -uroot -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result" +" 2>&1"  )
									remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
									//get the execute result from remote host by current build
									transmit.Download(exe_info[2],exe_info[4],testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result","./testcurrent/"+exe_info[7]);
								}
							}
						
						}
						
						
					}else{
						fmt.Println("状态异常，本轮不杀实例进程。")
						errnum = errnum + 1
					}
				}
			}
			
			if exe_info[11] == "kill15" {
				loopnum,_ := strconv.Atoi(strings.Split(strings.Split(exe_info[8], ".")[0], "-")[3]) 
				execute.Hakill15Case(exe_info[4], exe_info[2],user,exe_info[3], password, exe_info[5], exe_info[9], exe_info[8],testroot,testpath,loopnum)
			}
			
			if exe_info[11] == "disable" {
//			loopnum,_ := strconv.Atoi(strings.Split(strings.Split(exe_info[8], ".")[0], "-")[3]) 
//			execute.HaDisableCase(exe_info[2],exe_info[3] ,exe_info[4],strings.Split(strings.Split(exe_info[8], ",")[0], "-")[1],strings.Split(strings.Split(exe_info[8], ",")[0], "-")[2],exe_info[2],rootpass,loopnum)
				
			}
		}
		
	}

	
}

