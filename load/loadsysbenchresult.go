//加载sysbench的执行结果
package load

import (
//	"fmt"
	"os"
    "log"
    "bufio"
//    "strings"
)

func Loadsysbench(sysbenchfile string,linenum int) string{
//	fmt.Println("读取配置文件的go")
	var resultLine string
	resultLine,err := HandleSysbenchFile(sysbenchfile,linenum)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return resultLine
    }
	

}

func HandleSysbenchFile(textfile string,num int) (string,error) {	//返回切片和error信息
    file, err := os.Open(textfile)
//    var infoSlice []string =make([]string,0)
	var result string
    count := 0
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return result,err
    }
    defer file.Close()

    //声明一个切片，用来存放所有配置信息

    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text()  
	     count = count + 1
		  if count == num{
		  	result = line
		  }
	}

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return result,err
    }

//    return nil
//	fmt.Println("infoSliceinfo: ",len(infoSlice),cap(infoSlice),infoSlice)
	return result,nil
}

