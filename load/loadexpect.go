package load

import (
	"os"
    "log"
    "bufio"
//    "strings"
//	"fmt"
)

func Loadexpect(expectfile string) []string{
	var sliceinfo []string
	sliceinfo,err := HandleTextExpect(expectfile)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
    }
    
    
}

func HandleTextExpect(textfile string) ([]string,error) {	//返回切片和error信息
    file, err := os.Open(textfile)
    	var infoSlice []string =make([]string,0)
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text()  // orprofile
      //line := scanner.Bytes()
	     infoSlice = append(infoSlice,line)
	 }

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }

	return infoSlice,nil
}

