package main

import (
	"github.com/dynport/gossh"
	"log"
)

// returns a function of type gossh.Writer func(...interface{})
// MakeLogger just adds a prefix (DEBUG, INFO, ERROR)
func MakeLogger(prefix string) gossh.Writer {
	return func(args ...interface{}) {
		log.Println((append([]interface{}{prefix}, args...))...)
	}
}

func main() {
	client := gossh.New("130.9.1.41", "root")
	// my default agent authentication is used. use
	client.SetPassword("6yhnmko0")
	// for password authentication
	client.DebugWriter = MakeLogger("DEBUG")
	client.InfoWriter = MakeLogger("INFO ")
	client.ErrorWriter = MakeLogger("ERROR")

	defer client.Close()
	rsp, e := client.Execute("uptime")
	if e != nil {
		client.ErrorWriter(e.Error())
	}
	client.InfoWriter(rsp.String())

	rsp, e = client.Execute("echo -n $(cat /proc/loadavg); cat /does/not/exists")
	if e != nil {
		client.ErrorWriter(e.Error())
		client.ErrorWriter("STDOUT: " + rsp.Stdout())
		client.ErrorWriter("STDERR: " + rsp.Stderr())
	}
}
