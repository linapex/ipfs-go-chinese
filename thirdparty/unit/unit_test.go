
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:46</date>
//</624460192305647616>

package unit

import "testing"

//大多数meta奖授予…

func TestByteSizeUnit(t *testing.T) {
	if 1*KB != 1*1024 {
		t.Fatal(1 * KB)
	}
	if 1*MB != 1*1024*1024 {
		t.Fail()
	}
	if 1*GB != 1*1024*1024*1024 {
		t.Fail()
	}
	if 1*TB != 1*1024*1024*1024*1024 {
		t.Fail()
	}
	if 1*PB != 1*1024*1024*1024*1024*1024 {
		t.Fail()
	}
	if 1*EB != 1*1024*1024*1024*1024*1024*1024 {
		t.Fail()
	}
}

