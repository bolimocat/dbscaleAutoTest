//对执行结果进行数据清洗
package tool

import (
//	"fmt"
	"strings"
	local "dbscaleAutoTest/local"
)

func WashSysbench(sysbenchLine string,recordfile string){
//	[ 3480s ] thds: 32 tps: 665.52 qps: 15170.72 (r/w/o: 9885.64/2690.78/2594.30) lat (ms,99%): 189.93 err/s: 322.86 reconn/s: 0.00
	
	if strings.Contains(sysbenchLine, "tps: 0 qps: 0 (") {
		local.RecordSysbench(sysbenchLine, recordfile)
	}
	if strings.Contains(sysbenchLine, "tps: 0 qps") {
		local.RecordSysbench(sysbenchLine, recordfile)
	}
	if strings.Contains(sysbenchLine, "qps: 0 (") {
		local.RecordSysbench(sysbenchLine, recordfile)
	}
	
	
}
