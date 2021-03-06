
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:37</date>
//</624460155081199616>

package main

import (
	"fmt"
	"net"
	"os"

	logging "gx/ipfs/QmcuXC5cxs79ro2cUuHs4HQ2bkDLJUYokwL8aivcX6HW3C/go-log"
)

var log = logging.Logger("seccat")

func exit(format string, vals ...interface{}) {
	if format != "" {
		fmt.Fprintf(os.Stderr, "seccat: error: "+format+"\n", vals...)
	}
	Usage()
	os.Exit(1)
}

func out(format string, vals ...interface{}) {
	if verbose {
		fmt.Fprintf(os.Stderr, "seccat: "+format+"\n", vals...)
	}
}

type logConn struct {
	net.Conn
	n string
}

func (r *logConn) Read(buf []byte) (int, error) {
	n, err := r.Conn.Read(buf)
	if n > 0 {
		log.Debugf("%s read: %v", r.n, buf)
	}
	return n, err
}

func (r *logConn) Write(buf []byte) (int, error) {
	log.Debugf("%s write: %v", r.n, buf)
	return r.Conn.Write(buf)
}

func (r *logConn) Close() error {
	return r.Conn.Close()
}

