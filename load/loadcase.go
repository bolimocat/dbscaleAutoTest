package load

import (
	 "os"
    "log"
    "bufio"
    "strings"	

)

func Loadcase(casefile string) []string{
//	fmt.Println("读取配置文件的go")
	var sliceinfo []string
	sliceinfo,err := HandleTextCase(casefile)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
    }
}

func HandleTextCase(textfile string) ([]string,error) {	//返回切片和error信息
    file, err := os.Open(textfile)
    	var infoSlice []string =make([]string,0)
//    	var case_info string
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text()  // orprofile
      //line := scanner.Bytes()
	     
	     
	     //加载操作
		  if (!strings.HasPrefix(line, "--")&&!(len(line)==0)){//只要语句不以“--”开头，就都存入队列
		     //加载欲操作
//	     	infoSlice = append(infoSlice,strings.Split(line,":")[1])
			infoSlice = append(infoSlice,line)
		 }
	 }

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }

	return infoSlice,nil
}
