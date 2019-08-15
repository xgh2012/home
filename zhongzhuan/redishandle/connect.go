/**
* redis 操作
**/
package redishandle

import (
	"home/zhongzhuan/common"
)

var (
	pool = common.RedisPool
)

func Test() {
	client := pool.Get()
	defer client.Close()
	client.Do("SET", "xghsdfsdfds", 1232432)
	client.Do("HMSET", "SSS", "key1", "V1", "key2", "V2", "key3", "V3")
	//使用newPool构建一个redis连接池
	/*for i := 1;i <= 9;i++ {
		go func() {
			//从pool里面获取一个可用的redis连接
			c := pool.Get()
			defer c.Close()
			fmt.Println(c)
			//return
			//mset mget
			fmt.Printf("ActiveCount:%d IdleCount:%d\r\n",pool.Stats().ActiveCount,pool.Stats().IdleCount)
			c.Do("mset","name","biaoge","url","http://xxbandy.github.io")
			errCheck("setErr",setErr)
			if r,mgetErr := redis.Strings(c.Do("mget","name","url")); mgetErr == nil {
				for _,v := range r {
					fmt.Println("mget ",v)
				}
			}
		}()
	}*/

	//time.Sleep(1*time.Second)
	//r,err:=client.Do("GET","xghsdfsdfds")
	//fmt.Println(err)
	//fmt.Println(reflect.TypeOf(r))
	//username, err := redis.String(client.Do("GET", "xghsdfsdfds"))
	//fmt.Println(username)
	//fmt.Println(redis.Int64s(client.Do("Get","xghsdfsdfds")))
}
