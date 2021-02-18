package utility

import "testing"

func TestNewSnowflakeId(t *testing.T) {
	s1:=NewSnowflakeId(1,2)
	if s1.workerId!=1 || s1.datacenterId !=2{
		t.Fatalf("NewSnowflakeId failed,workerId:%d,datacenterId:%d\n", s1.workerId, s1.datacenterId)
	}

}

func TestSnowflakeId_NextId(t *testing.T) {
	s:=NewSnowflakeId(1,1)
	v1:=s.NextId()
	v2:=s.NextId()
	if v1 >= v2{
		t.Fatalf("NewSnowflakeId  NextId failed,v1 %d less than v2%d\n", v1,v2)
	}
	if v1<0 || v2<0{
		t.Fatalf("NewSnowflakeId  NextId value less than 0, v1:%d,v2%d\n", v1,v2)
	}
}