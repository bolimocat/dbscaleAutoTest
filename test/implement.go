//执行具体测试
package test

import (
//	"os/exec"
//	"fmt"
	"strings"
//	"log"
	"strconv"
)

func Testcase(expectfile []string,currentfile []string) []string{
	var expectlist []string = make([]string,0)
	var currentlist []string = make([]string,0)
	var errlist []string = make([]string,0)

	for index,value := range expectfile {
		expectlist = append(expectlist,"line:"+strconv.Itoa(index+1)+","+value)
		
	}
	for index,value := range currentfile {
		currentlist = append(currentlist,"line:"+strconv.Itoa(index+1)+","+value)
	}
	
	for i:=0;i<len(expectlist);i++ {
		if strings.Split(currentlist[i], ",")[1] != strings.Split(expectlist[i], ",")[1] {
			errlist = append(errlist,"current:[ "+currentlist[i]+" ]\nexpect: [ "+expectlist[i]+" ]")
	}
	}
	return errlist
	
}
