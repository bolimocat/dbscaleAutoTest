package main

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	 "database/sql"
	remote "dbscaleAutoTest/remote"		
	transmit "dbscaleAutoTest/transmit" 
	load "dbscaleAutoTest/load"
	local "dbscaleAutoTest/local"
//	test "dbscaleAutoTest/test"
	base "dbscaleAutoTest/database"
	execute "dbscaleAutoTest/execute"
	tool "dbscaleAutoTest/tool"

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
	var healthnode int //集群正常时节点数量
	var washfile string
//	var scenenum int //场景数
//	scenenum = len(scene_info)
	var currentscene int
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
			fmt.Println("在 "+value+" dbscale的节点上创建需要的目录")
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
			//从已有的预期文件分片列表，在每个dbscale所在节点上创建测试结果文件目录
			for _,v := range expressfilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testresult/"+strings.Split(v, ":")[0])
			}
			//从已有的预期文件分片列表，在每个dbscale所在节点上创建附件文件目录
			for _,v := range expressfilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/file/"+strings.Split(v, ":")[0])
			}
			
			//向每个目标节点上发送本次场景的测试用例sql文件
			for _,casename := range casefilelist {
				transmit.Upload(user, password, value, 22,"./testcase/"+strings.Split(casename, ":")[0]+"/"+strings.Split(casename, ":")[1],testroot+testpath+"/testcase/"+strings.Split(casename, ":")[0],"/"+strings.Split(casename, ":")[1],"向 "+value+" 发送casefile")
				
			}
			//向每个目标节点上发送本次场景的测试预期out文件
			for _,expressname := range expressfilelist{
				transmit.Upload(user, password, value, 22,"./testexpect/"+strings.Split(expressname, ":")[0]+"/"+strings.Split(expressname, ":")[1],testroot+testpath+"/testexpect/"+strings.Split(expressname, ":")[0],"/"+strings.Split(expressname, ":")[1],"向 "+value+" 发送expressfile")
			}
			//向每个目标节点上发送本次场景中测试需要的外部文件
			for _,attachname := range attachment_info{
				transmit.Upload(user, password, value, 22, "./file/"+strings.Split(attachname, ",")[7]+"/"+strings.Split(attachname, ",")[12], testroot+testpath+"/file/"+strings.Split(attachname, ",")[7],strings.Split(attachname, ",")[12], "向 "+value+" 发送attache")
				
			}
		}
	

	
	//读取场景信息，按行读取，切分每个元素的内容。
//	将sceneinfo的每行赋值给exe_info，在用","分割：
//	0：场景编号
//	1：用例编号，可以对应禅道上的用例编号
//	2：登录集群或数据库节点的用户名
//	3：登录集群或数据库节点的用户密码
//	4：登录的集群或数据库IP
//	5：登录的集群或数据库的端口
//	6：暂时没有实际使用
//	7：用例类型分类，结合bin目录下的用例目录、预期目录、本次实际结果目录等中的内容
//	8：用例文件名，测试时执行的sql文件
//	9：预期文件名，名称位置和8一样，扩展名改为out
//	10：当前行场景状态，TRUE为生效，FALSE为不生效
//	11：当前行场景的测试类型，function为功能sql测试，interface为接口测试，kill15和disable为对应的高可用测试
//	12：对应当前测试场景中如果有上传文件，这里写file中的文件名
	currentscene = 0
	for _,value := range scene_info{
		if value != "" {
			
			exe_info := strings.Split(value, ",")
			sceneid := exe_info[0]
			caseid := exe_info[1]
			dbuser := exe_info[2]
			dbpass := exe_info[3]
			dbhost := exe_info[4]
			dbport := exe_info[5]
			caseclass := exe_info[7]	//用例分类
			casename := exe_info[8]
			expectname := exe_info[9]
//			active := exe_info[10]
			casetype := exe_info[11]
//			attachment := exe_info[12]
			fmt.Println("执行场景：  ",sceneid,caseid,casename,casetype)
			
			conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",dbuser, dbpass, "tcp", dbhost, 16310, "dbscale")	
				db, err := sql.Open("mysql", conn)
			   defer db.Close()                                            //关闭数据库
			    err = db.Ping()                                            //连接数据库
			    if err != nil {
		        fmt.Println("数据库连接失败")                             //连接失败
		        return
			    } else {
			        fmt.Println("数据库连接成功")                             //连接成功
			     }
			 
			 
			
			 	if currentscene == 0{ //当进入第1个场景时，开始创建测试表和数据规则。
			 	//创建测试用库
					base.CreateTestDb(db)
					//创建表规则和分片规则
					base.CreateRule(db)
				 }  
			 	 

			 
			 
			
			
			 //获取集群正常时Working状态数量   
			healthnode = base.FetchHealthWorkingNum(db)
			
			//读取场景中的配置状态
			if casetype == "sysbench" {
//				init-oltp_read_write-50000-32.sql
				sysbenchPhase := strings.Split(strings.Split(casename, ".")[0], "-")[0]
				scriptfile := strings.Split(strings.Split(casename, ".")[0], "-")[1]
				tbsize := strings.Split(strings.Split(casename, ".")[0], "-")[2]
				threads := strings.Split(strings.Split(casename, ".")[0], "-")[3]
				recordfile := "./sysbench/record/"+sceneid+"_"+sysbenchPhase+"_"+scriptfile+"_"+tbsize+"_"+threads+".out"
				
				if sysbenchPhase == "init" {
					fmt.Println("init sysbench table : ")
					local.SysbenchInit(scriptfile, dbhost, dbuser, dbpass, tbsize, threads, recordfile)
				}
				if sysbenchPhase == "run" {
					fmt.Println("run sysbench test : ")
					runtime := strings.Split(strings.Split(casename, ".")[0], "-")[4]
					time.Sleep(10 * time.Second)
					washfile = recordfile
					go local.SysbenchRun(scriptfile, dbhost, dbuser, dbpass, tbsize, threads, runtime, recordfile)
					
					
					
				}
				if sysbenchPhase == "stop" {
					fmt.Println("stop sysbench test : ")
					local.SysbenchStop()
					for i:=0;i<360;i++ {
						tool.WashSysbench(load.Loadsysbench(washfile, 12+i),washfile+".result")
					}
				}
			}
			
			
			if casetype == "function" {
//				execute.FunctionCase(exe_info[7], exe_info[8], exe_info[2], exe_info[4], exe_info[3], exe_info[4], exe_info[9], password, exe_info[0], exe_info[1],testroot,testpath)
				execute.FunctionCase(caseclass, casename, dbuser, dbhost, dbpass,dbhost, expectname, password, sceneid, caseid,testroot,testpath)
			
			}
			if casetype == "interface" {
//				execute.InterfaseCase(exe_info[4], exe_info[2], exe_info[3], exe_info[5], "5", exe_info[8], exe_info[7],exe_info[0],exe_info[1])
				execute.InterfaseCase(dbhost, dbuser, dbpass, dbport, "5", casename, caseclass,sceneid,caseid)
				
			}
			if casetype == "stable" {
				fmt.Println("按行执行stable用例")
			}
			if casetype == "stress" {
				fmt.Println("按行执行stress用例")
				base.CreateStressDb(db)
				base.CreateStressRule(db)
		
				sysbenchPhase := strings.Split(strings.Split(casename, ".")[0], "-")[0]
				scriptfile := strings.Split(strings.Split(casename, ".")[0], "-")[1]
				tbsize := strings.Split(strings.Split(casename, ".")[0], "-")[2]
				threads := strings.Split(strings.Split(casename, ".")[0], "-")[3]
				recordfile := "./sysbench/record/"+sceneid+"_"+sysbenchPhase+"_"+scriptfile+"_"+tbsize+"_"+threads+".out"
				runtime := strings.Split(strings.Split(casename, ".")[0], "-")[4]
				local.StressRun(scriptfile, dbhost, dbuser, dbpass, tbsize, threads, runtime, recordfile)
				
				base.DropStressDb(db)
				base.DropStressRule(db)
				
			}
			
			if casetype == "ha"{
				if caseclass == "ifcontrol" {
					loopnum,_ :=  strconv.Atoi(strings.Split(strings.Split(casename, ".")[0], "-")[3])   //要被执行断网的次数
					execute.HaifCase(dbhost, dbuser, dbpass, dbport, password, rootpass, testroot, testpath, caseclass, casename, expectname,loopnum,healthnode)
				}
			
				if caseclass == "kill15" {
					loopnum,_ := strconv.Atoi(strings.Split(strings.Split(casename, ".")[0], "-")[3]) 
					execute.Hakill15Case(dbhost, dbuser,user,dbpass, password, dbport, expectname,casename,testroot,testpath,loopnum,healthnode)
				}
			
				if caseclass == "kill9" {
					loopnum,_ := strconv.Atoi(strings.Split(strings.Split(casename, ".")[0], "-")[3]) 
					execute.Hakill9Case(dbhost, dbuser,user,dbpass, password, dbport, expectname,casename,testroot,testpath,loopnum,healthnode)
				}
			
				if caseclass == "disable" {
				loopnum,_ := strconv.Atoi(strings.Split(strings.Split(casename, ".")[0], "-")[3]) 
				servername := strings.Split(strings.Split(casename,".")[0], "-")[1]
				rolename := strings.Split(strings.Split(casename,".")[0], "-")[2]
				execute.HaDisableCase(dbuser,dbpass,dbhost,servername,rolename,loopnum,user,password,dbport,expectname,casename,testroot,testpath)
				}
			}
			
			
			
			currentscene = currentscene + 1
		if currentscene == len(scene_info)-1 {
				//在最后一个场景结束后，删除测试表和规则。
				base.DropTestDb(db)
				base.DropRule(db)
		}
		}

	}

	
}

