package main

import (
	"fmt"
	"go-common/pkg/utility"
	"strconv"
)

func main() {

	s:=utility.NewSnowflakeId(1,1)
	for i:=0; i<2; i++{
		//fmt.Printf("Id:%d\n",worker.GetId())
		fmt.Printf("Id:%d\n",s.NextId())
	}

	var v1 int64=100
	v2:=strconv.FormatInt(v1,2)
	v3,_:=strconv.ParseInt("10011010",2,10)
	fmt.Printf("v1:%d,二进制:%s,v3:%d\n",v1,v2,v3)
}
