package load

import (
//	"fmt"
	"os"
    "log"
    "bufio"
    "strings"
)

func Loadscene(profile string) []string{
//	fmt.Println("读取配置文件的go")
	var sliceinfo []string
	sliceinfo,err := HandleTextScene(profile)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
    }
	

}

func HandleTextScene(textfile string) ([]string,error) {	//返回切片和error信息
    file, err := os.Open(textfile)
    	var infoSlice []string =make([]string,0)
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }
    defer file.Close()

    //声明一个切片，用来存放所有配置信息

    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text()  // or
      //line := scanner.Bytes()
	     
	     //加载操作用户
	     if strings.Contains(line,"TRUE"){
	     	//加载场景信息
	     	infoSlice = append(infoSlice,""+line)

	     }
	}

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }

//    return nil
//	fmt.Println("infoSliceinfo: ",len(infoSlice),cap(infoSlice),infoSlice)
	return infoSlice,nil
}