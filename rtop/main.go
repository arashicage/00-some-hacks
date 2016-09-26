/*

rtop - the remote system monitoring utility

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	//	"os/signal"
	"os/user"
	"path/filepath"
	//	"sort"
	"strconv"
	"strings"
	//	"syscall"
	"time"

	"rtop/misc"

	"net/http"
)

const VERSION = "1.0"
const DEFAULT_REFRESH = 5 // default refresh interval in seconds

//----------------------------------------------------------------------------
// Command-line processing

func usage(code int) {
	fmt.Printf(
		`rtop %s - (c) 2015 RapidLoop - MIT Licensed - http://rtop-monitor.org
rtop monitors server statistics over an ssh connection

Usage: rtop [-i private-key-file] [user@]host[:port] [interval]

	-i private-key-file
		PEM-encoded private key file to use (default: ~/.ssh/id_rsa if present)
	[user@]host[:port]
		the SSH server to connect to, with optional username and port
	interval
		refresh interval in seconds (default: %d)

`, VERSION, DEFAULT_REFRESH)
	os.Exit(code)
}

func usage2() {
	fmt.Printf(
		"open http://hostname:8080/stat?host=<hostname> to view stat\n")
}

func shift(q []string) (ok bool, val string, qnew []string) {
	if len(q) > 0 {
		ok = true
		val = q[0]
		qnew = q[1:]
	}
	return
}

func parseCmdLine() (key, username, addr string, interval time.Duration) {
	ok, arg, args := shift(os.Args)
	if len(args) == 0 {
		key = ""
		usr, _ := user.Current()
		username = usr.Username
		addr = ""
		interval = DEFAULT_REFRESH * time.Second
		return
	}

	var argKey, argHost, argInt string
	for ok {
		ok, arg, args = shift(args)
		if !ok {
			break
		}
		if arg == "-h" || arg == "--help" || arg == "--version" {
			usage(0)
		}
		if arg == "-i" {
			ok, argKey, args = shift(args)
			if !ok {
				usage(1)
			}
		} else if len(argHost) == 0 {
			argHost = arg
		} else if len(argInt) == 0 {
			argInt = arg
		} else {
			usage(1)
		}
	}
	if len(argHost) == 0 || argHost[0] == '-' {
		usage(1)
	}

	// key
	usr, err := user.Current()
	if err != nil {
		log.Print(err)
		usage(1)
	}
	if len(argKey) == 0 {
		key = filepath.Join(usr.HomeDir, ".ssh", "id_rsa")
		if _, err := os.Stat(key); os.IsNotExist(err) {
			key = ""
		}
	} else {
		key = argKey
	}
	// username, addr
	if i := strings.Index(argHost, "@"); i != -1 {
		username = argHost[:i]
		if i+1 >= len(argHost) {
			usage(1)
		}
		addr = argHost[i+1:]
	} else {
		username = usr.Username
		addr = argHost
	}
	if i := strings.Index(addr, ":"); i == -1 {
		addr += ":22"
	}
	// interval
	if len(argInt) == 0 {
		interval = DEFAULT_REFRESH * time.Second
	} else {
		i, err := strconv.ParseUint(argInt, 10, 64)
		if err != nil {
			log.Print(err)
			usage(1)
		}
		interval = time.Duration(i) * time.Second
	}

	return
}

//----------------------------------------------------------------------------

func main() {
	usage2()
	http.HandleFunc("/stat", stat)
	http.ListenAndServe(":8080", nil)
}

func stat(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	host := r.FormValue("host")

	log.SetPrefix("rtop: ")
	log.SetFlags(0)

	//	keyPath, username, addr, interval := parseCmdLine()
	// keyPath, username, addr, _ := parseCmdLine()

	keyPath := ""

	usr, _ := user.Current()
	username := usr.Username

	addr := host + ":22"

	client := misc.SSHConnect(username, addr, keyPath)

	//	hostname := strings.Split(addr, ":")[0]

	s := showStats(client, host)

	fmt.Fprint(w, s)
}

func showStats(client *ssh.Client, hostname string) string {
	stats := misc.Stats{}

	misc.GetAllStats(client, &stats, hostname)

	used := stats.MemTotal - stats.MemFree - stats.MemBuffers - stats.MemCached

	return fmt.Sprintf(
		//	fmt.Printf(
		stats.RedisInfo+
			`%s up %s

Load:
    %s %s %s

CPU:
    %.2f%% user, %.2f%% sys, %.2f%% nice, %.2f%% idle, %.2f%% iowait, %.2f%% hardirq, %.2f%% softirq, %.2f%% guest

Processes:
    %s running of %s total

Memory:
    free    = %s
    used    = %s
    buffers = %s
    cached  = %s
    swap    = %s free of %s

`,
		stats.Hostname, misc.FmtUptime(&stats),
		stats.Load1, stats.Load5, stats.Load10,
		stats.CPU.User,
		stats.CPU.System,
		stats.CPU.Nice,
		stats.CPU.Idle,
		stats.CPU.Iowait,
		stats.CPU.Irq,
		stats.CPU.SoftIrq,
		stats.CPU.Guest,
		stats.RunningProcs,
		stats.TotalProcs,
		misc.FmtBytes(stats.MemFree),
		misc.FmtBytes(used),
		misc.FmtBytes(stats.MemBuffers),
		misc.FmtBytes(stats.MemCached),
		misc.FmtBytes(stats.SwapFree),
		misc.FmtBytes(stats.SwapTotal),
	)
}
