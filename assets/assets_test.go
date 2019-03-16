
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:36</date>
//</624460150085783552>

package assets

import (
	"bytes"
	"io/ioutil"
	"sync"
	"testing"
)

//TestEmbeddedDocs确保文档更改后不会忘记重新生成
func TestEmbeddedDocs(t *testing.T) {
	testNFiles(initDocPaths, 5, t, "documents")
}

func TestDirIndex(t *testing.T) {
	t.Skip("skipping for now, code being tested is currently unused")
//TODO:在初始化期间导入资产。
//这需要弄清楚如何为
//从其GX路径引用代码
	testNFiles(initDirIndex, 2, t, "assets")
}

func testNFiles(fs []string, wantCnt int, t *testing.T, ftype string) {
	if len(fs) < wantCnt {
		t.Fatalf("expected %d %s. got %d", wantCnt, ftype, len(fs))
	}

	var wg sync.WaitGroup
	for _, f := range fs {
		wg.Add(1)
//比较资产
		go func(f string) {
			defer wg.Done()
			testOneFile(f, t)
		}(f)
	}
	wg.Wait()
}

func testOneFile(f string, t *testing.T) {
//从文件系统（Git）加载数据
	vcsData, err := ioutil.ReadFile(f)
	if err != nil {
		t.Errorf("asset %s: could not read vcs file: %s", f, err)
		return
	}

//从EMDEDED源加载数据
	embdData, err := Asset(f)
	if err != nil {
		t.Errorf("asset %s: could not read vcs file: %s", f, err)
		return
	}

	if !bytes.Equal(vcsData, embdData) {
		t.Errorf("asset %s: vcs and embedded data isnt equal", f)
		return
	}

	t.Logf("checked %s", f)
}

