package main

import (
	"fmt"
	"github.com/jakecoffman/cron"
	"time"
)

func main() {
	c := cron.New()

	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println(time.Now().String()[:19], "niehaha")

	}, "IAmUniqueName")

	c.Start()

	e := c.Entries()
	for _, et := range e {
		fmt.Println(et.Schedule, et.Next.String())
	}

	for {
		time.Sleep(120 * time.Second)
	}
}
