package local

import (
	"fmt"
//	local "dbscaleAutoTest/local"
)

func Localstoprole(stoprole string){
	if stoprole == "dbscale"{
		fmt.Println("停止dbscale角色")

		ControlDb("./goDeploy", "goDbscaleService", "stop")
	}
	if stoprole == "mysql"{
		fmt.Println("停止mysql角色")
		ControlDb("./goDeploy", "goMysqlService", "stop")
	}
	if stoprole == "all"{
		fmt.Println("停止dbscale 和 mysql角色")
		ControlDb("./goDeploy", "goMysqlService", "stop")
		ControlDb("./goDeploy", "goDbscaleService", "stop")
	}
	
}

func Localstartrole(startrole string){
	if startrole == "dbscale"{
		fmt.Println("启动dbscale角色")
		ControlDb("./goDeploy", "goDbscaleService", "start")
	}
	if startrole == "mysql"{
		fmt.Println("启动mysql角色")
		ControlDb("./goDeploy", "goMysqlService", "start")
	}
	if startrole == "all"{
		fmt.Println("启动 all 角色")
		ControlDb("./goDeploy", "goMysqlService", "start")
		ControlDb("./goDeploy", "goDbscaleService", "start")
	}
}
