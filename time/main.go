// 时间戳
// 当前时间戳
//
// fmt.Println(time.Now().Unix())
// # 1389058332
//
// str格式化时间
// 当前格式化时间
//
// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))  // 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
// # 2014-01-07 09:42:20
//
// 时间戳转str格式化时间
//
// str_time := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")
// fmt.Println(str_time)
// # 2014-01-07 09:32:12
//
//
// str格式化时间转时间戳
// 这个比较麻烦
//
// the_time := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)
// unix_time := the_time.Unix()
// fmt.Println(unix_time)
// # 389045004
//
//
// 还有一种方法,使用time.Parse
//
// the_time, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
// if err == nil {
//         unix_time := the_time.Unix()
// 	fmt.Println(unix_time)
// }
// # 1389171881

package main

import (
	"fmt"
	"time"
)

func main() {

	//获取时间戳 到秒
	timestamp1 := time.Now().Unix()
	fmt.Println(timestamp1)

	//获取时间戳 到纳秒
	timestamp2 := time.Now().UnixNano()
	fmt.Println(timestamp2)

	//格式化为字符串,tm为Time类型

	tm := time.Unix(timestamp1, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))

	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm2, _ := time.Parse("01/02/2006", "02/08/2015")
	fmt.Println(tm2.Unix())

}
