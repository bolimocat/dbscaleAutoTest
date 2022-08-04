//
package local

import (
	 "fmt"
    "os/exec"
    "time"
    "strconv"
)

//操作本地文件
func Mvlocalfile(sourcefile string,targetfile string){
	command := "mv "+sourcefile+" "+targetfile
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func Delcurrentfile(currentPath string){
	command := "rm -rf testcurrent/"+currentPath
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func Mvresult(){
	command := "mv testresult/myresult testresult/myresult"+strconv.FormatInt(time.Now().Unix(), 10)
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}