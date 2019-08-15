/**
* redis 操作
**/
package common

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

//构造一个连接池
//url为包装了redis的连接参数ip,port,passwd
func RedisNewPool(ip, port, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:         5, //定义redis连接池中最大的空闲链接为3
		MaxActive:       0, //在给定时间已分配的最大连接数(限制并发数)
		IdleTimeout:     240 * time.Second,
		MaxConnLifetime: 300 * time.Second,
		Dial:            func() (redis.Conn, error) { return redisConn(ip, port, passwd) },
	}
}

//构造一个链接函数，如果没有密码，passwd为空字符串
func redisConn(ip, port, passwd string) (redis.Conn, error) {
	c, err := redis.Dial("tcp",
		ip+":"+port,
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second),
		redis.DialPassword(passwd),
		redis.DialKeepAlive(1*time.Second),
	)
	return c, err
}

//构造一个错误检查函数
func errCheck(tp string, err error) {
	if err != nil {
		fmt.Printf("sorry,has some error for %s.\r\n", tp, err)
		os.Exit(-1)
	}
}
