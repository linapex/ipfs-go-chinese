
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:45</date>
//</624460186223906816>

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		usageError()
	}
	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not parse duration: %s\n", err)
		usageError()
	}

	time.Sleep(d)
}

func usageError() {
	fmt.Fprintf(os.Stderr, "Usage: %s <duration>\n", os.Args[0])
	fmt.Fprintln(os.Stderr, `Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".`)
fmt.Fprintln(os.Stderr, "See https://godoc.org/time parseDuration for more.”）
	os.Exit(-1)
}

