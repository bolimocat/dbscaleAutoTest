package load

import (
	 "os"
    "log"
    "bufio"
    "strings"	
    "fmt"
)

func Loadproperties(profile string) []string{
	fmt.Println("读取配置文件的go")
	var sliceinfo []string
	sliceinfo,err := HandleTextProperties(profile)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
    }
}

func HandleTextProperties(textfile string) ([]string,error) {	//返回切片和error信息
	 fmt.Println("properties : "+textfile )
    file, err := os.Open(textfile)
    	var infoSlice []string =make([]string,3)
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text() 
	     
	     //过滤普通用户用户名
	     if strings.Contains(line,"USER"){
	     	//加载欲操作
	     //user_name := strings.Split(line,"=")
	     	infoSlice = append(infoSlice,""+line)
		}
	      //过滤普通用户密码
	     if strings.Contains(line,"PASSWORD"){
	     	//加载欲操作
//	     	user_pass := strings.Split(line,"=")
	     	infoSlice = append(infoSlice,""+line)
		}
	      //过滤ROOT用户密码
	     if strings.Contains(line,"ROOTPASS"){
	     	//加载欲操作
//	     	root_pass := strings.Split(line,"=")
	     	infoSlice = append(infoSlice,""+line)
		}
	      //过滤所有dbscale节点
	     if strings.Contains(line,"DBSCALE_HOST"){
	     	//加载欲操作
//	     	dbscale_host := strings.Split(line,"=")
	     	infoSlice = append(infoSlice,""+line)
		}
	      if strings.Contains(line,"ALL_HOST"){
	     	//加载欲操作
//	     	dbscale_host := strings.Split(line,"=")
	     	infoSlice = append(infoSlice,""+line)
		}
	     if strings.Contains(line, "TESTROOT"){
	     	infoSlice = append(infoSlice,""+line)
	     }
	     if strings.Contains(line, "TESTPATH"){
	     	infoSlice = append(infoSlice,""+line)
	     }
	     if strings.Contains(line,"DBSCALEPORT"){
	     	infoSlice = append(infoSlice,""+line)
	     }
	     if strings.Contains(line, "DBSCALEPASS"){
	     	infoSlice = append(infoSlice,""+line)
	     }
		   
	 }

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }

	return infoSlice,nil
}