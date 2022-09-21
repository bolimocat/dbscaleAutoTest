//驱动用例活动
package execute

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	_ "github.com/go-sql-driver/mysql"
	 "database/sql"
	remote "dbscaleAutoTest/remote"
	transmit "dbscaleAutoTest/transmit" 
	load "dbscaleAutoTest/load"
	local "dbscaleAutoTest/local"
	test "dbscaleAutoTest/test"
	base "dbscaleAutoTest/database"
)


func FunctionCase(casetype string,casename string,user string,host string,dbpass string,dbport string,caseexp string,sshpass string,sceneid string,caseid string,testroot string,testpath string){
	fmt.Println("按行执行function用例")
	caselist := load.Loadcase("./testcase/"+casetype+"/"+casename)
				for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
					if strings.Split(value, ":")[0] == "sql"{
						//record the content of execution in fact
						fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
						remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+casetype+"/"+caseexp+".result"  )
						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+casetype+"/"+casename)
						remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+dbpass+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+casetype+"/"+caseexp+".result" +" 2>&1"  )
						remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+casetype+"/"+caseexp+".result"  )
						transmit.Download(user,host,testroot+testpath+"/testresult/"+casetype+"/"+caseexp+".result","./testcurrent/"+casetype);
					}
					if strings.Split(value, ":")[0] == "sys"{
						fmt.Println("执行系统操作，包括重启，停机等相关...")
					}
				}
				remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testresult/"+casetype+"/"+caseexp+".result")
				
			//执行测试
			var expectfile []string = make([]string,0)
			var currentfile []string = make([]string,0)

			expectfile = load.Loadexpect("./testexpect/"+casetype+"/"+caseexp)
			currentfile = load.Loadcurrent("./testcurrent/"+casetype+"/"+caseexp+".result")
			testresult := test.Testcase(expectfile, currentfile)
			if len(testresult) != 0{
				local.Recordresult("scene "+sceneid)
				local.Recordresult("zento caseid : "+caseid)
				local.Recordresult("test on dbscale : "+host)
				local.Recordresult("test type : function")
				local.Recordresult("test case sql : ./testcase/"+casetype+"/"+casename)
				local.Recordresult("test expect file : ./testexpect/"+casetype+"/"+caseexp)
				local.Recordresult("test current file : ./testcurrent/"+casetype+"/"+caseexp+".result")
				local.Recordresult("test err info :")
				for _,value := range testresult{
					fmt.Println("err info : "+value)
					local.Recordresult(value)
				}
				fmt.Println("There is sth wrong in your test!")
			}
			if len(testresult) == 0{
				local.Recordresult("scene "+sceneid)
				local.Recordresult("zento caseid : "+caseid)
				local.Recordresult("test on dbscale : "+host)
				local.Recordresult("test type : function")
				local.Recordresult("test case sql : ./testcase/"+casetype+"/"+casename)
				local.Recordresult("test expect file : ./testexpect/"+casetype+"/"+caseexp)
				local.Recordresult("test current file : ./testcurrent/"+casetype+"/"+caseexp+".result")
				local.Recordresult("The test is no wrong!")
				fmt.Println("The test is no wrong!")
			}
			local.Recordresult("\n")			
}

func InterfaseCase(host string,user string,password string,dbport string,tbnum string,jdbctype string,interfacetype string,sceneid string,caseid string){
					fmt.Println("按行执行interface用例")
				if interfacetype == "jdbc" {
					fmt.Println("执行jdbc测试  :")
					local.Interfacejdbc("jre/bin/java", " -jar" , "interface/testMysqlJdbc.jar "+host+" "+user+" "+password+" "+dbport+" test "+tbnum+" "+jdbctype+" > testcurrent/jdbc/"+jdbctype+".out.result")
					
				//执行测试
			var expectfile []string = make([]string,0)
			var currentfile []string = make([]string,0)

			expectfile = load.Loadexpect("./testexpect/jdbc/"+jdbctype+".out")
			currentfile = load.Loadcurrent("./testcurrent/jdbc/"+jdbctype+".out.result")
			testresult := test.Testcase(expectfile, currentfile)
			if len(testresult) != 0{
				local.Recordresult("scene "+sceneid)
				local.Recordresult("zento caseid : "+caseid)
				local.Recordresult("test on dbscale : "+host)
				local.Recordresult("test type : interface-jdbc")
				local.Recordresult("test case sql : ./testcase/jdbc/"+jdbctype)
				local.Recordresult("test expect file : ./testexpect/jdbc/"+jdbctype+".out")
				local.Recordresult("test current file : ./testcurrent/jdbc/"+jdbctype+".out.result")
				local.Recordresult("test err info :")
				for _,value := range testresult{
					fmt.Println("err info : "+value)
					local.Recordresult(value)
				}
				fmt.Println("There is sth wrong in your test!")
			}
			if len(testresult) == 0{
				local.Recordresult("scene "+sceneid)
				local.Recordresult("zento caseid : "+caseid)
				local.Recordresult("test on dbscale : "+host)
				local.Recordresult("test type : interface-jdbc")
				local.Recordresult("test case sql : ./testcase/jdbc/"+jdbctype)
				local.Recordresult("test expect file : ./testexpect/jdbc/"+jdbctype+".out")
				local.Recordresult("test current file : ./testcurrent/jdbc/"+jdbctype+".out.result")
				local.Recordresult("The test is no wrong!")
				fmt.Println("The test is no wrong!")
			}
			local.Recordresult("\n")	
				
				
				}	
}


func HaifCase(){
	
}




func Hakill15Case(host string,user string,cuser string,password string,sshpass string,dbport string,caseexp string,casename string,testroot string,testpath string,loopnum int){
		fmt.Println("按行执行ha-kill15用例")
				fmt.Println("执行主机 : "+host)
				conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",user, password, "tcp", host, 16310, "dbscale")	
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
//				var loopnum int
				for i := 0 ;i < loopnum; i++{
					if base.FetchDataserversInfo(db){
						remote.Nodemission(user, sshpass, host, 22, "echo '第 "+strconv.Itoa(i)+" '轮测试  >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/ha/"+casename)
						
						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "master"{
							killedhost = base.Killdataserversmaster(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
							fmt.Println("触发，杀死"+killedhost+"上，端口"+strings.Split(strings.Split(casename, ".")[0], "-")[1]+"的master实例。")
							remote.Nodemission(cuser, sshpass, killedhost, 22, "kill_pid.sh  15 "+strings.Split(strings.Split(casename, ".")[0], "-")[1])
							
							caselist := load.Loadcase("./testcase/ha/"+casename)
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
									remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/ha/"+casename)
									remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+password+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result" +" 2>&1"  )
									remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
									//get the execute result from remote host by current build
									transmit.Download(user,host,testroot+testpath+"/testresult/ha/"+caseexp+".result","./testcurrent/ha");
								}
							}
						}
						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "slave"{
							killedhost = base.Killdataserverslastslave(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
							fmt.Println("触发，杀死"+killedhost+"上，端口"+strings.Split(strings.Split(casename, ".")[0], "-")[1]+"的最后一个slave实例。")
							
							remote.Nodemission(cuser, sshpass, killedhost, 22, "kill_pid.sh  15 "+strings.Split(strings.Split(casename, ".")[0], "-")[1] )
//							remote.Nodemission(user, password, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
						
						caselist := load.Loadcase("./testcase/ha/"+casename)
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
									remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/ha/"+casename)
									remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+password+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result" +" 2>&1"  )
									remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
									//get the execute result from remote host by current build
									transmit.Download(user,host,testroot+testpath+"/testresult/ha/"+caseexp+".result","./testcurrent/ha");
								}
							}
						
						}
						
						
					}else{
						fmt.Println("等待中，不杀实例... ...")
						errnum = errnum + 1
					}
					if errnum > 5{
						fmt.Println("异常时间太久，检查节点状态是否恢复。")
						errnum = 0
						remote.Nodemission(user, sshpass, host, 22, "echo ' 异常时间太久，检查节点状态是否恢复。'  >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/ha/"+casename)
						remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
					}else{
						fmt.Println("状态基本正常")
					}
					
					time.Sleep(5 * time.Second)
				}
}


func HaDisableCase(dbuser string,dbpass string,dbhost string,servername string,rolename string,sshuser string,sshpass,loopnum int){
	fmt.Println("按行执行ha-disable用例")
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
	
//		disablelist := base.DisableServerInfo(db, "global_0")
//		for _,value := range base.DisableServerInfo(db, servername){
//			fmt.Println(" -- : "+value)
//		
//		
//				fmt.Println("执行主机 : "+strings.Split(value, ",")[1])
////				
//				var disableinfo []string
//				var errnum int
//				var loopnum int
//				for i := 0 ;i < loopnum; i++{
//					if base.FetchDataserversInfo(db){
////						remote.Nodemission(user, sshpass, host, 22, "echo '第 "+strconv.Itoa(i)+" '轮测试  >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
////						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/ha/"+casename)
//						
//						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "master"{
////							killedhost = base.Killdataserversmaster(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
//							disableinfo = base.DisableServerInfo(db, servername, "master")
//							fmt.Println("触发，disable "+strings.Split(disableinfo[0], ",")[0]+" on host "+strings.Split(disableinfo[0], ",")[1])
////							remote.Nodemission(cuser, sshpass, strings.Split(disableinfo[0], 22, "mysql -u"+dbuser+" -p"+password+" -h"+dbhost+" -P"+dbpass+" -e 'dbscale disable dataserver  "+strings.Split(disableinfo[0], ",")[0]+"'")
//							
//							caselist := load.Loadcase("./testcase/ha/"+casename)
//							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
//								if strings.Split(value, ":")[0] == "sql"{
//									//record the content of execution in fact
//									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
//									remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
//									remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/ha/"+casename)
//									remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+password+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result" +" 2>&1"  )
//									remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
//									//get the execute result from remote host by current build
//									transmit.Download(user,host,testroot+testpath+"/testresult/ha/"+caseexp+".result","./testcurrent/ha");
//								}
//							}
//						}
//						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "slave"{
//							killedhost = base.Killdataserverslastslave(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
//							fmt.Println("触发，杀死"+killedhost+"上，端口"+strings.Split(strings.Split(casename, ".")[0], "-")[1]+"的最后一个slave实例。")
//							
//							remote.Nodemission(cuser, sshpass, killedhost, 22, "kill_pid.sh  15 "+strings.Split(strings.Split(casename, ".")[0], "-")[1] )
////							remote.Nodemission(user, password, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
//						
//						caselist := load.Loadcase("./testcase/ha/"+casename)
//							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
//								if strings.Split(value, ":")[0] == "sql"{
//									//record the content of execution in fact
//									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
//									remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
//									remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/ha/"+casename)
//									remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+password+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result" +" 2>&1"  )
//									remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
//									//get the execute result from remote host by current build
//									transmit.Download(user,host,testroot+testpath+"/testresult/ha/"+caseexp+".result","./testcurrent/ha");
//								}
//							}
//						
//						}
//						
//						
//					}else{
//						fmt.Println("等待中，不杀实例... ...")
//						errnum = errnum + 1
//					}
//					if errnum > 5{
//						fmt.Println("异常时间太久，检查节点状态是否恢复。")
//						errnum = 0
//						remote.Nodemission(user, sshpass, host, 22, "echo ' 异常时间太久，检查节点状态是否恢复。'  >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
//						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/ha/"+casename)
//						remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
//					}else{
//						fmt.Println("状态基本正常")
//					}
//					
//					time.Sleep(5 * time.Second)
//		}
}









//-----------------FUNCTION-------------
//				fmt.Println("按行执行function用例")
//				caselist := load.Loadcase("./testcase/"+exe_info[7]+"/"+exe_info[8])
//				for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
//					if strings.Split(value, ":")[0] == "sql"{
//						//record the content of execution in fact
//						fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
//						remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
//						remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
//						remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -uroot -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result" +" 2>&1"  )
//						remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
//					//get the execute result from remote host by current build
//					transmit.Download(exe_info[2],exe_info[4],testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result","./testcurrent/"+exe_info[7]);
//					}
////					fmt.Println("test sql stmt : "+value)
////					remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '"+value+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
////					remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
////					remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -uroot -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e '"+value+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result" +" 2>&1"  )
////					remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
////					//get the execute result from remote host by current build
////					transmit.Download(exe_info[2],exe_info[4],testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result","./testcurrent/"+exe_info[7]);
//				
//					if strings.Split(value, ":")[0] == "sys"{
//						fmt.Println("执行系统操作，包括重启，停机等相关...")
//					}
//				}
//				
//				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result")
//				
//			//执行测试
//			var expectfile []string = make([]string,0)
//			var currentfile []string = make([]string,0)
//
//			expectfile = load.Loadexpect("./testexpect/"+exe_info[7]+"/"+exe_info[9])
//			currentfile = load.Loadcurrent("./testcurrent/"+exe_info[7]+"/"+exe_info[9]+".result")
////			fmt.Println("currentfile length : "+strconv.Itoa(len(currentfile)))
//			testresult := test.Testcase(expectfile, currentfile)
//			if len(testresult) != 0{
//				local.Recordresult("scene "+exe_info[0])
//				local.Recordresult("zento caseid : "+exe_info[1])
//				local.Recordresult("test on dbscale : "+exe_info[4])
//				local.Recordresult("test type : "+exe_info[11])
//				local.Recordresult("test case sql : ./testcase/"+exe_info[7]+"/"+exe_info[8])
//				local.Recordresult("test expect file : ./testexpect/"+exe_info[7]+"/"+exe_info[9])
//				local.Recordresult("test current file : ./testcurrent/"+exe_info[7]+"/"+exe_info[9]+".result")
//				local.Recordresult("test err info :")
//				for _,value := range testresult{
//					fmt.Println("err info : "+value)
//					local.Recordresult(value)
//				}
//				fmt.Println("There is sth wrong in your test!")
//			}
//			if len(testresult) == 0{
//				local.Recordresult("scene "+exe_info[0])
//				local.Recordresult("zento caseid : "+exe_info[1])
//				local.Recordresult("test on dbscale : "+exe_info[4])
//				local.Recordresult("test type : "+exe_info[11])
//				local.Recordresult("test case sql : ./testcase/"+exe_info[7]+"/"+exe_info[8])
//				local.Recordresult("test expect file : ./testexpect/"+exe_info[7]+"/"+exe_info[9])
//				local.Recordresult("test current file : ./testcurrent/"+exe_info[7]+"/"+exe_info[9]+".result")
//				local.Recordresult("The test is no wrong!")
//				fmt.Println("The test is no wrong!")
//			}
//			local.Recordresult("\n")
//-----------------FUNCTION-------------