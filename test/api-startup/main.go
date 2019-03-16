
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:44</date>
//</624460185636704256>

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	when := make(chan time.Time, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	for _, port := range []string{"5001", "8080"} {
		go func(port string) {
			defer wg.Done()
			for {
r, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s”，端口）
				if err != nil {
					continue
				}
				t := time.Now()
				when <- t
				log.Println(port, t, r.StatusCode)
				break
			}
		}(port)
	}
	wg.Wait()
	first := <-when
	second := <-when
	log.Println(second.Sub(first))
}

