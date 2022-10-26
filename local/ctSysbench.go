//控制本地sysbench
package local

import (
	 "fmt"
    "os/exec"
//    "time"
//    "strconv"
)

//初始化需要的测试库


//在库中初始化sysbench基础数据
func SysbenchInit(scriptfile string,dbhost string,dbuser string,dbpass string,tbsize string,threads string,recordfile string){
	command := "sysbench/initSysbench.sh "+scriptfile+" "+dbhost+" "+dbuser+" "+dbpass+" "+tbsize+" "+threads+" "+recordfile
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			fmt.Printf("Shell stmt : %s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//在库中初始化sysbench基础数据
func SysbenchRun(scriptfile string,dbhost string,dbuser string,dbpass string,tbsize string,threads string,runtime string,recordfile string){
	command := "sysbench/runSysbench.sh "+scriptfile+" "+dbhost+" "+dbuser+" "+dbpass+" "+tbsize+" "+threads+" "+runtime+" "+recordfile
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			fmt.Printf("Shell stmt : %s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//在库中初始化sysbench基础数据
func SysbenchStop(){
	command := "sysbench/stopSysbench.sh "
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			fmt.Printf("Shell stmt : %s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func StressRun(scriptfile string,dbhost string,dbuser string,dbpass string,tbsize string,threads string,runtime string,recordfile string){
	command := "sysbench/runSysbench.sh "+scriptfile+" "+dbhost+" "+dbuser+" "+dbpass+" "+tbsize+" "+threads+" "+runtime+" "+recordfile
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			fmt.Printf("Shell stmt : %s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}