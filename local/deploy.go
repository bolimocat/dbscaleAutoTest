package local

import (
	"fmt"
//	local "dbscaleAutoTest/local"
)

func Initcluster(){
	fmt.Println("init new dbscale cluster")
	ControlDb("./deploy", "goDeploy", "")
	ControlDb("./deploy", "goZkService", "-c start")
	ControlDb("./deploy", "goZkService", "-c init")
	ControlDb("./deploy", "goMysqlService", "-c init")
	ControlDb("./deploy", "goMysqlService", "-c startall")
	ControlDb("./deploy", "goMysqlService", "-c updateuser")
	ControlDb("./deploy", "goDbscaleService", "-c start")
	ControlDb("./deploy", "goDbscaleService", "-c init")
	ControlDb("./deploy", "goDbscaleService", "-c stop")
	ControlDb("./deploy", "goMysqlService", "-c stopall")
	ControlDb("./deploy", "goZkService", "-c stop")

}


func Localstop(stoprole string){
	if stoprole == "dbscale"{
		fmt.Println("停止dbscale角色")
		ControlDb("./deploy", "goDbscaleService", "stop")
	}
	if stoprole == "mysql"{
		fmt.Println("停止mysql角色")
		ControlDb("./deploy", "goMysqlService", "stop")
	}
	if stoprole == "zookeeper"{
		fmt.Println("停止zookeeper角色")
		ControlDb("./deploy", "goZkService", "stop")
	}
	if stoprole == "all"{
		fmt.Println("停止dbscale 和 mysql角色")
		ControlDb("./deploy", "goMysqlService", "stop")
		ControlDb("./deploy", "goZkService", "stop")
		ControlDb("./deploy", "goDbscaleService", "stop")
	}
	
}

func Localstart(startrole string){
	if startrole == "dbscale"{
		fmt.Println("启动dbscale角色")
		ControlDb("./deploy", "goDbscaleService", "start")
	}
	if startrole == "mysql"{
		fmt.Println("启动mysql角色")
		ControlDb("./deploy", "goMysqlService", "start")
	}
	if startrole == "zookeeper"{
		fmt.Println("启动zookeeper角色")
		ControlDb("./deploy", "goZkService", "start")
	}
	if startrole == "all"{
		fmt.Println("启动 all 角色")
		ControlDb("./deploy", "goZkService", "start")
		ControlDb("./deploy", "goMysqlService", "start")
		ControlDb("./deploy", "goDbscaleService", "start")
	}
}
