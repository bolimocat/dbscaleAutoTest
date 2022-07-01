package local

import (
	"os/exec"
)

func Recordresult(recordinfo string){
			
	command := "echo '"+recordinfo+"' >> ./testresult/myresult"
	cmd := exec.Command("/bin/bash", "-c", command)

    _, err := cmd.Output()
    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
//    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
