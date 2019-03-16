
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:44</date>
//</624460184021897216>

package mfsr

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/ipfs/go-ipfs/thirdparty/assert"
)

func testVersionFile(v string, t *testing.T) (rp RepoPath) {
	name, err := ioutil.TempDir("", v)
	if err != nil {
		t.Fatal(err)
	}
	rp = RepoPath(name)
	return rp
}

func TestVersion(t *testing.T) {
	rp := RepoPath("")
	_, err := rp.Version()
	assert.Err(err, t, "Should throw an error when path is bad,")

	rp = RepoPath("/path/to/nowhere")
	_, err = rp.Version()
	if !os.IsNotExist(err) {
		t.Fatalf("Should throw an `IsNotExist` error when file doesn't exist: %v", err)
	}

	fsrepoV := 5

	rp = testVersionFile(strconv.Itoa(fsrepoV), t)
	_, err = rp.Version()
	assert.Err(err, t, "Bad VersionFile")

	assert.Nil(rp.WriteVersion(fsrepoV), t, "Trouble writing version")

	assert.Nil(rp.CheckVersion(fsrepoV), t, "Trouble checking the version")

	assert.Err(rp.CheckVersion(1), t, "Should throw an error for the wrong version.")
}

