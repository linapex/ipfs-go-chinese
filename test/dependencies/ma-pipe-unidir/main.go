
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:45</date>
//</624460186613977088>

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	ma "gx/ipfs/QmNTCey11oxhb1AxDnQBRHtdhap6Ctud872NjAYPYYXPuc/go-multiaddr"
	manet "gx/ipfs/QmZcLBXKaFe8ND5YHPkJRAwmhJGrVsi1JqDZNyJ4nRK5Mj/go-multiaddr-net"
)

const USAGE = "ma-pipe-unidir [-l|--listen] [--pidFile=path] [-h|--help] <send|recv> <multiaddr>\n"

type Opts struct {
	Listen  bool
	PidFile string
}

func app() int {
	opts := Opts{}
	flag.BoolVar(&opts.Listen, "l", false, "")
	flag.BoolVar(&opts.Listen, "listen", false, "")
	flag.StringVar(&opts.PidFile, "pidFile", "", "")
	flag.Usage = func() {
		fmt.Print(USAGE)
	}
	flag.Parse()
	args := flag.Args()

if len(args) < 2 { //<模式> ADDR>
		fmt.Print(USAGE)
		return 1
	}

	mode := args[0]
	addr := args[1]

	if mode != "send" && mode != "recv" {
		fmt.Print(USAGE)
		return 1
	}

	if len(opts.PidFile) > 0 {
		data := []byte(strconv.Itoa(os.Getpid()))
		err := ioutil.WriteFile(opts.PidFile, data, 0644)
		if err != nil {
			return 1
		}

		defer os.Remove(opts.PidFile)
	}

	maddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return 1
	}

	var conn manet.Conn

	if opts.Listen {
		listener, err := manet.Listen(maddr)
		if err != nil {
			return 1
		}

		conn, err = listener.Accept()
		if err != nil {
			return 1
		}
	} else {
		var err error
		conn, err = manet.Dial(maddr)
		if err != nil {
			return 1
		}
	}

	defer conn.Close()
	switch mode {
	case "recv":
		io.Copy(os.Stdout, conn)
	case "send":
		io.Copy(conn, os.Stdin)
	default:
		return 1
	}
	return 0
}

func main() {
	os.Exit(app())
}

