package tool

import (
	"os"
//	"io/ioutil"
	"time"
	"log"
	"fmt"
	"path/filepath"
)

//计算超时时间，当前时间大于一个目标时间。
func Overtime(atime int64,otime int64) bool{
	//input a aim time,that is fixed,
		var result bool
	if( otime + time.Now().Unix()) > atime {
		result = true
	}else{
		result = false
	}
	return result
}

func Listfolder(folderPath string) (listname []string){
	pwd,_ := os.Getwd()
	filepathNames,err := filepath.Glob(filepath.Join(pwd,folderPath+"*"))
	if err != nil {
		log.Fatal(err)
	}
 
//	for i := range filepathNames {
//		fmt.Println(filepathNames[i]) //打印path
//	}
	return filepathNames
}

func Listfolder1() {
	pwd,_ := os.Getwd()
	filepathNames,err := filepath.Glob(filepath.Join(pwd,"./testcase/*"))
	if err != nil {
		log.Fatal(err)
	}
 
	for i := range filepathNames {
		fmt.Println(filepathNames[i]) //打印path
	}
//	return filepathNames
}