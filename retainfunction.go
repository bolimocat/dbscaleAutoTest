package dbscaleAutoTest

import (

)

func re(){
	//读取场景中的配置状态
//			fmt.Println("--"+exe_info[11])
			dbproperties := strings.Split(exe_info[11], "||")
			if (dbproperties[0] == "null")&& (dbproperties[1] == "null") {
				fmt.Println("本场景不需要修改dbscale或mysql的配置")
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i '1i\\tee  "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+"' "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n'>>"+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo 'notee;'>>"+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -u"+exe_info[2]+" -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e 'source "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8] +"  '  ")
			}
			if (dbproperties[0] != "null") && (dbproperties[1] == "null") {
				fmt.Println("本场景需要修改dbscale的配置")
				local.Localstoprole("dbscale");
				
				local.Localstartrole("dbscale");
				
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i '1i\\tee  "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+"' "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n'>>"+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo 'notee;'>>"+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -u"+exe_info[2]+" -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e 'source "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8] +"  '  ")
			}
			if (dbproperties[0] == "null") && (dbproperties[1] != "null") {
				fmt.Println("本场景需要修改mysql的配置")
				local.Localstoprole("mysql");
				
				local.Localstartrole("mysql");
				
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i '1i\\tee  "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+"' "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n'>>"+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo 'notee;'>>"+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -u"+exe_info[2]+" -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e 'source "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8] +"  '  ")
			}
			if (dbproperties[0] != "null") && (dbproperties[1] != "null") {
				fmt.Println("本场景需要修改dbscale和mysql的配置")
				local.Localstoprole("all");
				
				local.Localstartrole("all");
				
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i '1i\\tee  "+testroot+testpath+"/testresult/"+exe_info[7]+"/"+exe_info[9]+"' "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "sed -i \"s/\\x0//g\" "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo '\n'>>"+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "echo 'notee;'>>"+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8])
				remote.Nodemission(exe_info[2], password, exe_info[4], 22, "mysql -u"+exe_info[2]+" -p"+exe_info[3]+" -h"+exe_info[4]+" -P"+exe_info[5]+" -e 'source "+testroot+testpath+"/testcase/"+exe_info[7]+"/"+exe_info[8] +"  '  ")
			}
}