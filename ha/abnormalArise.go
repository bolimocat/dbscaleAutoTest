package ha

import (
	"fmt"
    "os/exec"
)

//
func KillPid(user string,password string,host_ip string,port int,basepth string){
	command := "shfile/establishFolder.sh "+host_ip+" "+basepth
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

