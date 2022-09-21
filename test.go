package main

import (
	ab "dbscaleAutoTest/ha"
)

func main(){
	ab.Killpid("greatdb", "abc123", "192.168.2.18",22,"16318")
}
