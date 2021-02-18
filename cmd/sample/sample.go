package main

import (
	"fmt"
	"go-common/pkg/utility"
)

func main() {

	s:=utility.NewSnowflakeId(1,1)
	for i:=0; i<2; i++{
		//fmt.Printf("Id:%d\n",worker.GetId())
		fmt.Printf("Id:%d\n",s.NextId())
	}

	var arr [257] int

	for i:=0; i<257; i++ {
		arr[i]= i+1
	}

	for i,v:=range arr{
		fmt.Printf("第%d位,值为:%d\n", i, v)
	}
}
