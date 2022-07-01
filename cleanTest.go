package main

import (
	"fmt"
	load "dbscaleAutoTest/load"
	"strings"
	remote "dbscaleAutoTest/remote"
)

func main(){
	var properties_info []string = load.Loadproperties("./file/properties")
	var dbscale_host string
	var user string
	var password string
	var testroot string 
	var testpath string
	var rootpass string
	var uif []string = make([]string,7)
	
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
				testroot = strings.Split(value, "=")[1]
				uif = append(uif,testroot)
			}
			if index == 8 {
				testpath = strings.Split(value, "=")[1]
				uif = append(uif,testpath)
			}
			
		}
	}
	
	//clean test path
	host := strings.Split(dbscale_host, ",")
		for _,value := range host {
			fmt.Println("host : " + value)
			
			//在每个dbscale的节点上创建测试目录，存放测试场景和测试用例
			fmt.Println("dbscale的节点上delete test目录")
			remote.Nodemission(user, password, value, 22, "rm  -rf "+testroot+testpath)
			
		}
	
}

