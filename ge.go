//ge for Generate expectations 
//使用方法：./ge ./testcase/jdbc/test1.sql
package main


import (
	"fmt"
	"os"
	"strings"
	load "dbscaleAutoTest/load"
	remote "dbscaleAutoTest/remote"
	transmit "dbscaleAutoTest/transmit"

)

func main(){
	var casesql string
	var caseout string
	var attachfile string
	
	//加载参数：top文件位置和pid号
	for index,value := range os.Args{
//		fmt.Println("arg[] :"+value)
		if index == 1 {
			casesql = value
		}
		if index == 2 {
			attachfile = value
		}

	}
	
	casenamelength := len(strings.Split(casesql, "."))
	//判断选择的sql文件是否以sql作为扩展名
	if strings.Split(casesql, ".")[casenamelength - 1] != "sql" {
		fmt.Println("需要sql文件")
		
	}
	if strings.Split(casesql, ".")[casenamelength - 1] == "sql" {
		fmt.Println("要执行的sql文件 : "+casesql)
		caseout = strings.Split(casesql, ".")[casenamelength - 2]+".out"
		var dbscale_host string
		var user string
		var password string
		var port string
		var dbscalepass string
		
		var properties_info []string = load.Loadproperties("./file/properties")
		for index,value := range properties_info{
		if value != "" {
			if index == 3 {
				user = strings.Split(value, "=")[1]
			}
			if index == 4 {
				password = strings.Split(value, "=")[1]
			}
			if index == 6 {
				dbscale_host = strings.Split(strings.Split(value, "=")[1], ",")[0]
			}
			if index == 9 {
				port = strings.Split(value, "=")[1]
			}
			if index == 10 {
				dbscalepass = strings.Split(value,"=")[1]
			}
		}
		
	}
//		将sql文件放入远端的/tmp
		
		realcasesql := strings.Split(casesql, "/")
		realcaseout := strings.Split(caseout, "/")
		
		transmit.Upload(user, password, dbscale_host, 22,casesql,"/tmp",casesql,"向 "+dbscale_host+" 发送sql")
		
		if attachfile != "" {
			transmit.Upload(user, password, dbscale_host, 22, attachfile, "/tmp", attachfile, "向 "+dbscale_host+" 发送 "+attachfile)
		}
		
//		执行sql
		caselist := load.Loadcase(casesql)
		for _,value := range caselist {
//			fmt.Println(" -- "+strings.Split(value, ":")[1])
			remote.Nodemission(user, password, dbscale_host, 22, "echo '"+strings.Split(value, ":")[1]+"' >> /tmp/"+realcaseout[len(realcaseout)-1]  )
			fmt.Println("mysql  --binary-mode=1 -uroot -p"+dbscalepass+" -h"+dbscale_host+" -P"+port+" -e '"+strings.Split(value, ":")[1]+" ' >> /tmp/"+realcaseout[len(realcaseout)-1]  )
			remote.Nodemission(user, password, dbscale_host, 22, "mysql -uroot -p"+dbscalepass+" -h"+dbscale_host+" -P"+port+" -e '"+strings.Split(value, ":")[1]+"' >> /tmp/"+realcaseout[len(realcaseout)-1] +" 2>&1"  )	
			remote.Nodemission(user, password, dbscale_host, 22, "echo '\n' >> /tmp/"+realcaseout[len(realcaseout)-1]   )
		}
		remote.Nodemission(user, password, dbscale_host, 22, "sed -i \"s/\\x0//g\" /tmp/"+realcasesql[len(realcasesql)-1])
		
		

		
	}
	

}