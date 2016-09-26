package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	// "strings"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	// url := "https://www.baidu.com/"
	url := "https://www.ss-link.com/my/free"

	httpGetBody(url)

}

func main2() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	// url := "https://www.baidu.com/"
	url := "http://192.168.5.102/fpcy.html"
	N := 5000
	sem := make(chan int, N)
	begin := time.Now().Add(time.Second * 30)

	// httpGetBody(url)

	fmt.Println(begin)

	for i := 0; i < N; i++ {
		go httpGet(i, url, begin, sem)
	}

	for i := 0; i < N; i++ {
		<-sem
	}

}

func httpGet(i int, url string, begin time.Time, sem chan<- int) {
	done := true
	for done {
		if time.Now().Sub(begin) >= 0 {
			timeBegin := time.Now()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			timeEnd := time.Now()
			fmt.Printf("# %8d total time %s\n	begin:%s \n	end  :%s\n", i, timeEnd.Sub(timeBegin), timeBegin, timeEnd)
			// fmt.Println(string(body))
			sem <- 1

			done = false
		}
	}

}

func httpGetBody(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}
