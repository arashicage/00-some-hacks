package main

import (
	"log"

	"code.google.com/p/mahonia"
	"github.com/garyburd/redigo/redis"
)

func main() {

	rs, err := redis.Dial("tcp", "130.9.1.22:6379")
	if err != nil {
		log.Println(err)
	}
	defer rs.Close()

	v, err := redis.Values(rs.Do("HGETALL", "0X:120013414001873645"))

	if err != nil {
		panic(err)
	}

	m, err := redis.StringMap(v, err)

	for k, v := range m {
		log.Println(k, v)
	}

	enc := mahonia.NewDecoder("gbk")
	for k, v := range m {
		log.Println("+++++++", k, enc.ConvertString(v))
	}

}
