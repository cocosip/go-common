package utility

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

var staticMachine = getMachineHash()
var staticIncrement = getRandomNumber()
var staticPid = int32(os.Getpid())

type ObjectId struct {
	timestamp int64
	machine   int32
	pid       int32
	increment int32
}

func NewObjectId() ObjectId {
	timestamp := time.Now().Unix()
	return ObjectId{timestamp, staticMachine, staticPid, atomic.AddInt32(&staticIncrement, 1) & 0xffffff}
}

func Parse(input string) (objId ObjectId, err error) {
	if objId, err = tryParse(input); err == nil {
		return
	}
	return objId, fmt.Errorf("%s is not a valid 24 digit hex string", input)
}

func (objId ObjectId) Timestamp() int64 {
	return objId.timestamp
}

func (objId ObjectId) Machine() int32 {
	return objId.machine
}

func (objId ObjectId) Pid() int32 {
	return objId.pid
}

func (objId ObjectId) Increment() int32 {
	return objId.increment & 0xffffff
}

// String returns the ObjectID id as a 24 byte hex string representation.
func (objId ObjectId) String() string {
	array := []byte{
		byte(objId.timestamp >> 0x18),
		byte(objId.timestamp >> 0x10),
		byte(objId.timestamp >> 8),
		byte(objId.timestamp),
		byte(objId.machine >> 0x10),
		byte(objId.machine >> 8),
		byte(objId.machine),
		byte(objId.pid >> 8),
		byte(objId.pid),
		byte(objId.increment >> 0x10),
		byte(objId.increment >> 8),
		byte(objId.increment),
	}
	return hex.EncodeToString(array)
}

func getMachineHash() int32 {
	machineName, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	buf := md5.Sum([]byte(machineName))
	return (int32(buf[0])<<0x10 + int32(buf[1])<<8) + int32(buf[2])
}

func getRandomNumber() int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31()
}

func tryParse(input string) (objId ObjectId, err error) {
	if len(input) != 0x18 {
		return objId, errors.New("invalid input length")
	}
	array, err := hex.DecodeString(input)
	if err != nil {
		return objId, err
	}
	return ObjectId{
		timestamp: int64(array[0])<<0x18 + int64(array[1])<<0x10 + int64(array[2])<<8 + int64(array[3]),
		machine:   int32(array[4])<<0x10 + int32(array[5])<<8 + int32(array[6]),
		pid:       int32(array[7])<<8 + int32(array[8]),
		increment: int32(array[9])<<0x10 + (int32(array[10]) << 8) + int32(array[11]),
	}, nil
}