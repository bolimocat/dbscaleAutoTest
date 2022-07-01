package main

import (
	"fmt"
	"strings"
	load "dbscaleAutoTest/load"
//	test "dbscaleAutoTest/test"
//	"strconv"
	remote "dbscaleAutoTest/remote"
	transmit "dbscaleAutoTest/transmit" 
//	tool "dbscaleAutoTest/tool"
	local "dbscaleAutoTest/local"
	test "dbscaleAutoTest/test"

)

func main(){
	fmt.Println("Cluster 4.0 Auto Test suit")
	var scene_info []string = load.Loadscene("./testscene/sceneinfo.csv")
	var properties_info []string = load.Loadproperties("./file/properties")
	var dbscale_host string
	var user string
	var password string
	var testroot string 
	var testpath string
	var rootpass string
//	var casefolder string
//	var caseinfo string
	var uif []string = make([]string,0)
	var casefilelist []string = make([]string,0)
	var expressfilelist []string = make([]string,0)
	var dbpropertieslist []string = make([]string,0)

	
	for index,value := range properties_info{
		if value != "" {
//				fmt.Println("index : "+strconv.Itoa(index))
//				fmt.Println(" :-: "+properties_info[index])
			if index == 3 {
				user = strings.Split(value, "=")[1]
//				fmt.Println("user : "+user )
				uif = append(uif,user)
			}
			if index == 4 {
				password = strings.Split(value, "=")[1]
//				fmt.Println("password : "+password)
				uif = append(uif,password)
			}
			if index == 5 {
				rootpass = strings.Split(value, "=")[1]
//				fmt.Println("rootpass : "+rootpass)
				uif = append(uif,rootpass)
			}
			if index == 6 {
				dbscale_host = strings.Split(value, "=")[1]
//				fmt.Println("dbscale_host : "+dbscale_host)
				uif = append(uif,dbscale_host)
			}
			if index == 7 {
				testroot = strings.Split(value, "=")[1]
//				fmt.Println("testroot : "+testroot)
				uif = append(uif,testroot)
			}
			if index == 8 {
				testpath = strings.Split(value, "=")[1]
//				fmt.Println("testpath : "+testpath)
				uif = append(uif,testpath)
			}
			
		}
	}
	
	//filt all testcase and express file from scene
	for _,value := range scene_info{
		if value != "" {
			casefile_info := strings.Split(value, ",")
			
//			fmt.Println("casefile_info[7] ="+casefile_info[7])//case filt
			casefilelist = append(casefilelist,casefile_info[7]+":"+casefile_info[8])
			expressfilelist = append(expressfilelist,casefile_info[7]+":"+casefile_info[9])
			dbpropertieslist = append(dbpropertieslist,casefile_info[11])
		}
		
	}
	
	//过滤所有dbscale节点
		host := strings.Split(dbscale_host, ",")
		for _,value := range host {
			fmt.Println("host : " + value)
			
			//在每个dbscale的节点上创建测试目录，存放测试用例
			fmt.Println("dbscale的节点上创建需要的目录")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath)
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testcase")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testresult")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testexpect")
			remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/interface")
			
			for _,v := range casefilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testcase/"+strings.Split(v, ":")[0])
			}
			
			for _,v := range expressfilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testexpect/"+strings.Split(v, ":")[0])
			}
			
			for _,v := range expressfilelist {
				remote.Nodemission(user, password, value, 22, "mkdir -p "+testroot+testpath+"/testresult/"+strings.Split(v, ":")[0])
			}
			
			for _,casename := range casefilelist {
				transmit.Upload(user, password, value, 22,"./testcase/"+strings.Split(casename, ":")[0]+"/"+strings.Split(casename, ":")[1],testroot+testpath+"/testcase/"+strings.Split(casename, ":")[0],"/"+strings.Split(casename, ":")[1],"向 "+value+" 发送casefile")
				
			}
			
			for _,expressname := range expressfilelist{
				transmit.Upload(user, password, value, 22,"./testexpect/"+strings.Split(expressname, ":")[0]+"/"+strings.Split(expressname, ":")[1],testroot+testpath+"/testexpect/"+strings.Split(expressname, ":")[0],"/"+strings.Split(expressname, ":")[1],"向 "+value+" 发送expressfile")
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
//			fmt.Println("--"+exe_info[11])
			if exe_info[11] == "function" {
				fmt.Println("按行执行function用例")
			
				caselist := load.Loadcase("./testcase/"+exe_info[7]+"/"+exe_info[8])
				for _,value := range caselist {
					remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '"+value+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
					remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
					remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -uroot -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e '"+value+"' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result" +" 2>&1"  )
					remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n' >> "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result"  )
					//get the execute result from remote host by current build
					transmit.Download(exe_info[2],exe_info[4],testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result","./testcurrent/"+exe_info[7]);


				}
				
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+".result")
				
			}
			if exe_info[11] == "interface" {
				fmt.Println("按行执行interface用例")
//					fmt.Println(" ---- "+exe_info[2], password, exe_info[4], 22, "java -jar "+testroot+testpath+"/interface/dbscaleJdbc.jar "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8]);
//					local.Interfacejdbc("./interface", "java -jar ./dbscaleJdbc.jar ", "./testcase/jdbc/jdbc_createtb")
					
			}
			if exe_info[11] == "stable" {
				fmt.Println("按行执行stable用例")
			}
			if exe_info[11] == "stress" {
				fmt.Println("按行执行stress用例")
			}
			
			
			
			//执行测试
			var expectfile []string = make([]string,0)
			var currentfile []string = make([]string,0)

			expectfile = load.Loadexpect("./testexpect/"+exe_info[7]+"/"+exe_info[9])
			currentfile = load.Loadcurrent("./testcurrent/"+exe_info[7]+"/"+exe_info[9]+".result")
			
			testresult := test.Testcase(expectfile, currentfile)
			if len(testresult) != 0{
				local.Recordresult("scene "+exe_info[0])
				local.Recordresult("zento caseid : "+exe_info[1])
				local.Recordresult("test on dbscale : "+exe_info[4])
				local.Recordresult("test type : "+exe_info[11])
				local.Recordresult("test case sql : ./testcase/"+exe_info[7]+"/"+exe_info[8])
				local.Recordresult("test expect file : ./testexpect/"+exe_info[7]+"/"+exe_info[9])
				local.Recordresult("test current file : ./testcurrent/"+exe_info[7]+"/"+exe_info[9]+".result")
				local.Recordresult("test err info :")
				for _,value := range testresult{
					fmt.Println("err info : "+value)
					local.Recordresult(value)
				}
				fmt.Println("There is sth wrong in your test!")
			}
			if len(testresult) == 0{
				local.Recordresult("scene "+exe_info[0])
				local.Recordresult("zento caseid : "+exe_info[1])
				local.Recordresult("test on dbscale : "+exe_info[4])
				local.Recordresult("test type : "+exe_info[11])
				local.Recordresult("test case sql : ./testcase/"+exe_info[7]+"/"+exe_info[8])
				local.Recordresult("test expect file : ./testexpect/"+exe_info[7]+"/"+exe_info[9])
				local.Recordresult("test current file : ./testcurrent/"+exe_info[7]+"/"+exe_info[9]+".result")
				local.Recordresult("The test is no wrong!")
				fmt.Println("The test is no wrong!")
			}
			local.Recordresult("\n")
		}
		
	}

	
}

