package redis

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"os"
	"errors"
)

var rd redis.Conn
var isConnect bool

var isNotConnectError = errors.New("until:redis:redis没有初始化")

func Init(address string) {
	var err error
	rd, err = redis.Dial("tcp", address)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
	isConnect = true
}

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if !isConnect {
		return nil, isNotConnectError
	}

	return rd.Do(commandName, args...)
}

func Int(commandName string, args ...interface{}) (int, error) {
	return redis.Int(Do(commandName, args...))
}
