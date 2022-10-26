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

//sql功能测试
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

//执行接口测试
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

//执行高可用断网测试
func HaifCase(dbhost string,dbuser string,dbpass string,dbport string,password string,rootpass string,testroot string,testpath string,caseclass string,casename string,expectname string,killnum int,healthnum int){
		fmt.Println("按行执行断网用例")
				fmt.Println("执行主机 : "+dbhost)
				//连接集群库，查看操作过程中的节点状态信息。
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
				
				var killedhost string  //实际要被执行断网的ip地址
				var errnum int	//异常状态计数
				var mos string //master or slavae
				var ifdownPort string //要断网的端口
//				var killnum int 
				
				
				for i:=0;i<killnum;i++{
//					if base.FetchDataserversInfo(db){//judge the datasource status is working
					if base.FetchCurrentWorkingNum(db) == healthnum {
						//在每个实际执行了测试的远程节点上开始写实际测试结果
						remote.Nodemission("root", rootpass, dbhost, 22, "echo '第 "+strconv.Itoa(i)+" '轮测试  >> "+testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result"  )
						remote.Nodemission("root", rootpass, dbhost, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+caseclass+"/"+casename)
					
						//根据用例信息判断去master还是slave上执行断网
						mos = strings.Split(strings.Split(casename, ".")[0], "-")[2] //从用例名的'.'切，第0个元素，再切第0个元素，第3个元素
						if mos == "master"{//if the node is master,so ifdown the interface on master host
							//判断当前应该断哪个IP的网
							ifdownPort = strings.Split(strings.Split(casename, ".")[0], "-")[1] //casename切2遍，第一遍过滤掉扩展名，第二遍得到端口。
							killedhost = base.IfdownHost(db, ifdownPort,"master")
							fmt.Println("全部实例状态恢复，断开端口"+ifdownPort+"的master节点所在的网络。")
							//用root到远程主机的/usr/bin下执行断网操作
							fmt.Println("killedhost : "+killedhost)
							fmt.Println(" -- root", password, killedhost, 22, "ifcontrol.sh > /dev/null 2>&1 & ")
							remote.Nodemission("root", rootpass, killedhost, 22, "ifcontrol.sh > /dev/null 2>&1 & " )
							time.Sleep(15 * time.Second)
							fmt.Println("ifdown over")
							//读取测试工具本地的用例分类-用例名，将用例sql文本的每一行做为切片内容
							caselist := load.Loadcase("./testcase/"+caseclass+"/"+casename)
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{//每行的内容用“：”分割，第0个元素如果是“sql”，则第1个元素作为执行的sql语句，使用-e的方式在远端执行
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission("root", rootpass, dbhost, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result"  )
									remote.Nodemission("root", rootpass, dbhost, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+caseclass+"/"+casename)
									remote.Nodemission("root", rootpass, dbhost, 22, "mysql -uroot -p"+dbpass+" -h"+dbhost+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result" +" 2>&1"  )
									remote.Nodemission("root", rootpass, dbhost, 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result"  )
									//get the execute result from remote host by current build
									//将本次在远端执行的结果文本，回传到本地工具的current相关目录下
									transmit.Download("root",dbhost,testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result","./testcurrent/"+caseclass);
								}
							}
						}
						if mos == "slave"{
							killedhost = base.IfdownHost(db, ifdownPort,"slave")
							fmt.Println("全部实例状态恢复，断开端口"+ifdownPort+"的last slave节点所在的网络。")
							remote.Nodemission("root", rootpass, killedhost, 22, "kill_pid.sh  15 "+strings.Split(strings.Split(casename, ".")[0], "-")[1] )
							
						
						caselist := load.Loadcase("./testcase/"+caseclass+"/"+casename)
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission("root", rootpass, dbhost, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result"  )
									remote.Nodemission("root", rootpass, dbhost, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+caseclass+"/"+casename)
									remote.Nodemission("root", rootpass, dbhost, 22, "mysql -uroot -p"+dbpass+" -h"+dbhost+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result" +" 2>&1"  )
									remote.Nodemission("root", rootpass, dbhost, 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result"  )
									//get the execute result from remote host by current build
									//将本次在远端执行的结果文本，回传到本地工具的current相关目录下
									transmit.Download("root",dbhost,testroot+testpath+"/testresult/"+caseclass+"/"+expectname+".result","./testcurrent/"+caseclass);
								}
							}
						
						}
					}else{ //当本次判断时，datasource的状态不是Working，则记录一次状态未恢复或异常
						fmt.Println("状态异常，本轮不杀实例进程。") //只有当所有节点的状态恢复之后，才重新断网或杀实例，以确保出现无法恢复的状态是由单独原因造成的
						errnum = errnum + 1
					}
				}
}



//执行高可用kill-15测试
func Hakill15Case(host string,user string,cuser string,password string,sshpass string,dbport string,caseexp string,casename string,testroot string,testpath string,loopnum int,healthnum int){
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
//					if base.FetchDataserversInfo(db){
					fmt.Println("目标状态数 : "+strconv.Itoa(healthnum)+" 当前状态数 : "+strconv.Itoa(base.FetchCurrentWorkingNum(db)))
					if base.FetchCurrentWorkingNum(db) == healthnum {
						remote.Nodemission(user, sshpass, host, 22, "echo '第 "+strconv.Itoa(i)+" '轮测试  >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result"  )
						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/kill15/"+casename)
						
						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "master"{
							killedhost = base.Killdataserversmaster(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
							fmt.Println(strconv.Itoa(i+1)+"/"+strconv.Itoa(loopnum)+" 触发 kill -15 ，杀死"+killedhost+"上，端口"+strings.Split(strings.Split(casename, ".")[0], "-")[1]+"的master实例。")
							remote.Nodemission(cuser, sshpass, killedhost, 22, "kill_pid.sh  15 "+strings.Split(strings.Split(casename, ".")[0], "-")[1])
							
							caselist := load.Loadcase("./testcase/kill15/"+casename)
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result"  )
									remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/kill15/"+casename)
									remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+password+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result" +" 2>&1"  )
									remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result"  )
									//get the execute result from remote host by current build
									transmit.Download(user,host,testroot+testpath+"/testresult/kill15/"+caseexp+".result","./testcurrent/kill15");
								}
							}
						}
						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "slave"{
							killedhost = base.Killdataserverslastslave(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
							fmt.Println(strconv.Itoa(i)+"/"+strconv.Itoa(loopnum)+" 触发 kill -15 ，杀死"+killedhost+"上，端口"+strings.Split(strings.Split(casename, ".")[0], "-")[1]+"的最后一个slave实例。")
							
							remote.Nodemission(cuser, sshpass, killedhost, 22, "kill_pid.sh  15 "+strings.Split(strings.Split(casename, ".")[0], "-")[1] )
//							remote.Nodemission(user, password, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
						
						caselist := load.Loadcase("./testcase/kill15/"+casename)
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result"  )
									remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/kill15/"+casename)
									remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+password+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result" +" 2>&1"  )
									remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result"  )
									//get the execute result from remote host by current build
									transmit.Download(user,host,testroot+testpath+"/testresult/kill15/"+caseexp+".result","./testcurrent/kill15");
								}
							}
						
						}
						
						
					}else{
						fmt.Println("等待中，不杀实例... ...")
						errnum = errnum + 1
						fmt.Println("异常程度： "+strconv.Itoa(errnum)+"/10")
						i = i - 1
					}
					if errnum >= 10{
						fmt.Println("异常时间太久，检查节点状态是否恢复。")
						errnum = 0
						remote.Nodemission(user, sshpass, host, 22, "echo ' 异常时间太久，检查节点状态是否恢复。'  >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result"  )
						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/kill15/"+casename)
						remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/kill15/"+caseexp+".result"  )
						fmt.Println("集群节点太长时间不能恢复正常，跳出本次循环 ... ...")
						break
					}else{
						fmt.Println("")
					}
					
					time.Sleep(5 * time.Second)
				}
}


//执行高可用kill-9测试
func Hakill9Case(host string,user string,cuser string,password string,sshpass string,dbport string,caseexp string,casename string,testroot string,testpath string,loopnum int,healthnum int){
		fmt.Println("按行执行ha-kill9用例")
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
//					if base.FetchDataserversInfo(db){
					fmt.Println("目标状态数 : "+strconv.Itoa(healthnum)+" 当前状态数 : "+strconv.Itoa(base.FetchCurrentWorkingNum(db)))
					if base.FetchCurrentWorkingNum(db) == healthnum {
						remote.Nodemission(user, sshpass, host, 22, "echo '第 "+strconv.Itoa(i)+" '轮测试  >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result"  )
						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/kill9/"+casename)
						
						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "master"{
							killedhost = base.Killdataserversmaster(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
							fmt.Println(strconv.Itoa(i+1)+"/"+strconv.Itoa(loopnum)+" 触发 kill -9 ，杀死"+killedhost+"上，端口"+strings.Split(strings.Split(casename, ".")[0], "-")[1]+"的master实例。")
							remote.Nodemission(cuser, sshpass, killedhost, 22, "kill_pid.sh  9 "+strings.Split(strings.Split(casename, ".")[0], "-")[1])
							
							caselist := load.Loadcase("./testcase/kill9/"+casename)
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result"  )
									remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/kill9/"+casename)
									remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+password+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result" +" 2>&1"  )
									remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result"  )
									//get the execute result from remote host by current build
									transmit.Download(user,host,testroot+testpath+"/testresult/kill9/"+caseexp+".result","./testcurrent/kill9");
								}
							}
						}
						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "slave"{
							killedhost = base.Killdataserverslastslave(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
							fmt.Println(strconv.Itoa(i)+"/"+strconv.Itoa(loopnum)+" 触发 kill -9 ，杀死"+killedhost+"上，端口"+strings.Split(strings.Split(casename, ".")[0], "-")[1]+"的最后一个slave实例。")
							
							remote.Nodemission(cuser, sshpass, killedhost, 22, "kill_pid.sh  15 "+strings.Split(strings.Split(casename, ".")[0], "-")[1] )
//							remote.Nodemission(user, password, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/ha/"+caseexp+".result"  )
						
						caselist := load.Loadcase("./testcase/kill9/"+casename)
							for _,value := range caselist {	//load all testcase from the file,then split the string in every line for judge the first posation is sql or sys.
								if strings.Split(value, ":")[0] == "sql"{
									//record the content of execution in fact
									fmt.Println("test sql stmt : "+strings.Split(value, ":")[1])
									remote.Nodemission(user, sshpass, host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result"  )
									remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/kill9/"+casename)
									remote.Nodemission(user, sshpass, host, 22, "mysql -uroot -p"+password+" -h"+host+" -P"+dbport+" -e '"+strings.Split(value, ":")[1]+"' >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result" +" 2>&1"  )
									remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result"  )
									//get the execute result from remote host by current build
									transmit.Download(user,host,testroot+testpath+"/testresult/kill9/"+caseexp+".result","./testcurrent/kill9");
								}
							}
						
						}
						
						
					}else{
						fmt.Println("等待中，不杀实例... ...")
						errnum = errnum + 1
						fmt.Println("异常程度： "+strconv.Itoa(errnum)+"/10")
						i = i - 1
					}
					if errnum >= 10{
						fmt.Println("异常时间太久，检查节点状态是否恢复。")
						errnum = 0
						remote.Nodemission(user, sshpass, host, 22, "echo ' 异常时间太久，检查节点状态是否恢复。'  >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result"  )
						remote.Nodemission(user, sshpass, host, 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/kill9/"+casename)
						remote.Nodemission(user, sshpass, host, 22, "echo '\n' >> "+testroot+testpath+"/testresult/kill9/"+caseexp+".result"  )
						fmt.Println("集群节点太长时间不能恢复正常，跳出本次循环 ... ...")
						break
					}else{
						fmt.Println("")
					}
					
					time.Sleep(5 * time.Second)
				}
}

//host string,user string,cuser string,password string,sshpass string,dbport string,caseexp string,casename string,testroot string,testpath string,loopnum int,healthnum int
func HaDisableCase(dbuser string,dbpass string,dbhost string,servername string,rolename string,loopnum int,user string,password string,dbport string,caseexp string,casename string,testroot string,testpath string){
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

//		for _,value := range base.DisableServerInfo(db, servername,rolename){
//				fmt.Println("执行主机 : "+strings.Split(value, ",")[1])
//				var disableinfo []string
//				var errnum int
//				fmt.Println(" value "+value)
				for i := 0 ;i < loopnum; i++{
//					if base.FetchDataserversInfo(db){
						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "master"{
							value:=base.DisableServerInfo(db, servername,rolename)[0]
							remote.Nodemission(user, password, strings.Split(value,",")[1], 22, "echo '第 "+strconv.Itoa(i)+" '轮测试  >> "+testroot+testpath+"/testresult/disable/"+caseexp+".result"  )
							remote.Nodemission(user, password, strings.Split(value,",")[1], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/disable/"+casename)
//							killedhost := base.Killdataserversmaster(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
//							disableinfo = base.DisableServerInfo(db, servername, "master")
							fmt.Println("触发，disable master "+strings.Split(value,",")[1]+" on host "+strings.Split(value,",")[0])
							remote.Nodemission(user, password, strings.Split(value,",")[1],22, "mysql -u"+dbuser+" -p"+dbpass+" -h"+dbhost+" -P"+dbport+" -e 'dbscale disable dataserver  "+strings.Split(value,",")[0]+"'")
							time.Sleep(10 * time.Second)
							remote.Nodemission(user, password, strings.Split(value,",")[1],22, "mysql -u"+dbuser+" -p"+dbpass+" -h"+dbhost+" -P"+dbport+" -e 'dbscale enable dataserver  "+strings.Split(value,",")[0]+"'")

						}
						if strings.Split(strings.Split(casename, ".")[0], "-")[2] == "slave"{
							value:=base.DisableServerInfo(db, servername,rolename)[0]
							remote.Nodemission(user, password, strings.Split(value,",")[1], 22, "echo '第 "+strconv.Itoa(i)+" '轮测试  >> "+testroot+testpath+"/testresult/disable/"+caseexp+".result"  )
							remote.Nodemission(user, password, strings.Split(value,",")[1], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/disable/"+casename)
//							killedhost := base.Killdataserversmaster(db, strings.Split(strings.Split(casename, ".")[0], "-")[1])
//							disableinfo = base.DisableServerInfo(db, servername, "slave")
							fmt.Println("触发，disable slave "+strings.Split(value,",")[1]+" on host "+strings.Split(value,",")[0])
							remote.Nodemission(user, password,strings.Split(value,",")[1], 22, "mysql -u"+dbuser+" -p"+dbpass+" -h"+dbhost+" -P"+dbport+" -e 'dbscale disable dataserver  "+strings.Split(value,",")[0]+"'")
							time.Sleep(10 * time.Second)
							remote.Nodemission(user, password,strings.Split(value,",")[1], 22, "mysql -u"+dbuser+" -p"+dbpass+" -h"+dbhost+" -P"+dbport+" -e 'dbscale enable dataserver  "+strings.Split(value,",")[0]+"'")						
			
						}
						
						
//					}else{
//						fmt.Println("等待中，不disable dataserver... ...")
//						errnum = errnum + 1
//					}
//					if errnum > 10{
//						fmt.Println("异常时间太久，检查节点状态是否恢复。")
//						errnum = 0
//						remote.Nodemission(user, password, strings.Split(disableinfo[0],",")[0], 22, "echo ' 异常时间太久，检查节点状态是否恢复。'  >> "+testroot+testpath+"/testresult/disable/"+caseexp+".result"  )
//						remote.Nodemission(user, password, strings.Split(disableinfo[0],",")[0], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/disable/"+casename)
//						remote.Nodemission(user, password, strings.Split(disableinfo[0],",")[0], 22, "echo '\n' >> "+testroot+testpath+"/testresult/disable/"+caseexp+".result"  )
//					}else{
//						fmt.Println("状态基本正常")
//					}
					
					time.Sleep(5 * time.Second)
		}
//}
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