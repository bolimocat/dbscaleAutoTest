package main

import (
	"fmt"
	"os/exec"
	"os"
)

func main(){
	
	var killtype string
	var killport string
	
	//加载参数：top文件位置和pid号
	for index,value := range os.Args{
//		if index == 1 {
//			filepath = value
//		}
		if index == 1 {
			killtype = value
		}
		if index == 2 {
			killport = value
		}

	}
	
	
	command := "kill_pid.sh "+killtype+" "+killport
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
   
}

