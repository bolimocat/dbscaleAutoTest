//临时使用scp从远程向本地下载需要的文件
package transmit

import (
	 "os/exec"
//	 "fmt"
)

func Download(remoteUser string,remoteHost string,remoteFile string,localPath string){

	command := "scp -r "+remoteUser+"@"+remoteHost+":"+remoteFile+" "+localPath
	cmd := exec.Command("/bin/bash", "-c", command)

    _, err := cmd.Output()
    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
//    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
