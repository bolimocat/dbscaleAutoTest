package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//dataservers表结构体定义
type Dataservers struct {
	 servername string
	 host string
	 port string
	 username string
	 status string
	 master_online_status string
	 master_backup string
	 remote_user string
	 remote_port string
	 max_needed_conn_max_mysql_conn string
	 master_priority string
}

func GrantAllToRoot(DB *sql.DB){
	err := DB.QueryRow("grant all on *.* to root@'%';")
	if err != nil{
		 fmt.Printf("grant all on *.* to root@'%' :%v\n", err)
       return
	}
	 fmt.Println("grant all on *.* to root@'%' successd:")
}

func NewGlobalDatasource(DB *sql.DB){
	err := DB.QueryRow("dbscale dynamic add schema dataspace global_db datasource = global_ds;")
	if err != nil{
		 fmt.Printf("dbscale dynamic add schema dataspace global_db datasource = global_ds;", err)
       return
	}
	 fmt.Println("dbscale dynamic add schema dataspace global_db datasource = global_ds; successd")
}

func FetchDataserversInfo (DB *sql.DB) bool {
	fmt.Println("search status from dataservers")
	rows,err := DB.Query("dbscale show dataservers where status like '%normal%';")
	var result = false
	var count = 0
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
//		var dataserver Dataservers
//		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
//		fmt.Println("servername: "+dataserver.servername+" host: "+dataserver.host+" port: "+dataserver.port+"  username: "+dataserver.username+" status: "+dataserver.status+" master_online_status: "+dataserver.master_online_status)
		count = count + 1
	}
	if count == 24{
			result = true
		}else{
			result = false
		}
	return result
}

func Killdataserversmaster (DB *sql.DB,port string) string {
	rows,err := DB.Query("dbscale show dataservers where status like '%normal%';")
	var killed string
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var dataserver Dataservers
		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
//		fmt.Println("servername: "+dataserver.servername+" host: "+dataserver.host+" port: "+dataserver.port+"  username: "+dataserver.username+" status: "+dataserver.status+" master_online_status: "+dataserver.master_online_status)
//		if ((dataserver.port == port)&&(dataserver.master_online_status == 'Master_Online')){
//			killed = dataserver.host
//		}
		if ((dataserver.port == port)&&(dataserver.master_online_status == "Master_Online")) {
				killed = dataserver.host
			}
		
	}
	return killed
}

func Killdataserverslastslave (DB *sql.DB,port string) string {
	rows,err := DB.Query("dbscale show dataservers where status like '%normal%';")
	var killed string
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var dataserver Dataservers
		
//		var slavelist []string = make([]string,1)
		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
		if ((dataserver.port == port)&&(dataserver.master_online_status == "Not_A_Master")) {
				killed = dataserver.host
			}
	}
	return killed
}

func DisableServerInfo (DB *sql.DB,servername string,rolename string) []string {
	rows,err := DB.Query("dbscale show dataservers where servername like '%"+servername+"%';")
	var disableInfo []string = make([]string,0)
//	var disableInfo string
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var dataserver Dataservers
		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
//		disablelist = append(disablelist,dataserver.servername+","+dataserver.host+","+dataserver.port+","+dataserver.master_online_status)
//		if (rolename == "master")&&(ataserver.master_online_status == "Master_Online") {
//			disableInfo = append(disableInfo,dataserver.servername+","+dataserver.host+","+dataserver.port)
//		}
//		if (rolename == "slave")&&(ataserver.master_online_status == "Not_A_Master") {
//			disableInfo = append(disableInfo,dataserver.servername+","+dataserver.host+","+dataserver.port)
//		}
		
	}
	return disableInfo
}



//,username,status,master_online_status,master_backup,remote_user,remote_port,max_needed_conn_max_mysql_conn,master_priority
