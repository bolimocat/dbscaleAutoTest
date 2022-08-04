package load

import (
	 "os"
    "log"
    "bufio"
    "strings"	

)

func Loadattachment(attachfile string) []string{
//	fmt.Println("读取配置文件的go")
	var sliceinfo []string
	sliceinfo,err := HandleFile(attachfile)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
    }
}

func HandleFile(attachfile string) ([]string,error) {	//返回切片和error信息
    file, err := os.Open(attachfile)
    	var infoSlice []string =make([]string,0)
//    	var case_info string
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", attachfile, err)
        return infoSlice,err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text()  // orprofile
	     if strings.Contains(line,"TRUE"){
	     	attachname := strings.Split(line, ",")[12]
	     	if attachname != ""{
	     		infoSlice = append(infoSlice,line)
	     	}

	     }
       }

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", attachfile, err)
        return infoSlice,err
    }

	return infoSlice,nil
}

