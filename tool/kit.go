package tool

import (
	"os"
//	"io/ioutil"
	"log"
	"fmt"
	"path/filepath"
)

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