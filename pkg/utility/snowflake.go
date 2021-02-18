package utility

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type SnowflakeId struct {
	m *sync.Mutex
	workerId int64
	datacenterId int64
	sequence int64
	lastTimestamp int64
}

const (
	//开始时间截(2015-01-01)
	Twepoch int64 = 1420041600000
	//机器id所占的位数
	WorkerIdBits int32 = 5
	//数据标识id所占的位数
	DatacenterIdBits int32 = 5
	//支持的最大机器id，结果是31 (这个移位算法可以很快的计算出几位二进制数所能表示的最大十进制数)
	MaxWorkerId int64= -1 ^ (-1 << WorkerIdBits)
	//支持的最大数据标识id，结果是31
	MaxDatacenterId int64 = -1 ^ (-1 << DatacenterIdBits)
	//序列在id中占的位数
	SequenceBits int32 = 12
	//机器ID向左移12位
	WorkerIdShift = SequenceBits
	//数据标识id向左移17位(12+5)
	DatacenterIdShift = SequenceBits + WorkerIdBits
	//时间截向左移22位(5+5+12)
	TimestampLeftShift = SequenceBits + WorkerIdBits + DatacenterIdBits
	//生成序列的掩码，这里为4095 (0b111111111111=0xfff=4095)
	SequenceMask int64 = -1 ^ (-1 << SequenceBits)
)

func (s *SnowflakeId) NextId() int64 {
	s.m.Lock()
	defer s.m.Unlock()
	timestamp := timeGen()
	//fmt.Printf("timestamp:%d，sequence:%d\n",timestamp,s.sequence)
	//如果当前时间小于上一次ID生成的时间戳，说明系统时钟回退过这个时候应当抛出异常
	if timestamp < s.lastTimestamp {
		err := errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", s.lastTimestamp-timestamp))
		panic(err)
	}
	//如果是同一时间生成的，则进行毫秒内序列
	if s.lastTimestamp == timestamp {
		s.sequence = (s.sequence + 1) & SequenceMask
		//毫秒内序列溢出
		if s.sequence == 0 {
			//阻塞到下一个毫秒,获得新的时间戳
			timestamp = tilNextMillis(s.lastTimestamp)
		}
	} else {
		//时间戳改变，毫秒内序列重置
		s.sequence = 0
	}
	s.lastTimestamp = timestamp
	//fmt.Printf("timestamp:%d,lastTimestamp:%d,sequence:%d\n",timestamp,s.lastTimestamp,s.sequence)
	id := ((timestamp - Twepoch) << TimestampLeftShift) | (s.datacenterId << DatacenterIdShift) | (s.workerId << WorkerIdShift) | s.sequence
	return id
}

func timeGen() int64 {
	return 	time.Now().UTC().UnixNano()/1000000
}

func tilNextMillis(lastTimestamp int64) int64 {
  timestamp := timeGen()
  for{
	  if timestamp <= lastTimestamp {
		  timestamp = timeGen()
	  }else{
	  	break
	  }
  }
	return timestamp
}

func NewSnowflakeId(workerId int64,datacenterId int64) SnowflakeId {
	if workerId > MaxWorkerId || workerId < 0 {
		panic(errors.New(fmt.Sprintf("worker Id can't be greater than %d or less than 0", MaxWorkerId)))
	}
	if datacenterId > MaxDatacenterId || datacenterId < 0 {
		panic(errors.New(fmt.Sprintf("datacenter Id can't be greater than %d or less than 0", MaxDatacenterId)))
	}
	s := SnowflakeId{workerId: workerId, datacenterId: datacenterId,sequence: 0,lastTimestamp: -1,m: new(sync.Mutex)}
	return s
}

var DefaultSnowflakeId SnowflakeId = NewSnowflakeId(1,1)


