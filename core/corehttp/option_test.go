
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:41</date>
//</624460172118462464>

package corehttp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	version "github.com/ipfs/go-ipfs"
)

type testcasecheckversion struct {
	userAgent    string
	uri          string
	shouldHandle bool
	responseBody string
	responseCode int
}

func (tc testcasecheckversion) body() string {
	if !tc.shouldHandle && tc.responseBody == "" {
		return fmt.Sprintf("%s (%s != %s)\n", errAPIVersionMismatch, version.ApiVersion, tc.userAgent)
	}

	return tc.responseBody
}

func TestCheckVersionOption(t *testing.T) {
	tcs := []testcasecheckversion{
		{"/go-ipfs/0.1/", APIPath + "/test/", false, "", http.StatusBadRequest},
		{"/go-ipfs/0.1/", APIPath + "/version", true, "check!", http.StatusOK},
		{version.ApiVersion, APIPath + "/test", true, "check!", http.StatusOK},
		{"Mozilla Firefox/no go-ipfs node", APIPath + "/test", true, "check!", http.StatusOK},
		{"/go-ipfs/0.1/", "/webui", true, "check!", http.StatusOK},
	}

	for _, tc := range tcs {
		t.Logf("%#v", tc)
		r := httptest.NewRequest("POST", tc.uri, nil)
r.Header.Add("User-Agent", tc.userAgent) //旧版本，应该失败

		called := false
		root := http.NewServeMux()
		mux, err := CheckVersionOption()(nil, nil, root)
		if err != nil {
			t.Fatal(err)
		}

		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			called = true
			if !tc.shouldHandle {
				t.Error("handler was called even though version didn't match")
			} else {
				io.WriteString(w, "check!")
			}
		})

		w := httptest.NewRecorder()

		root.ServeHTTP(w, r)

		if tc.shouldHandle && !called {
			t.Error("handler wasn't called even though it should have")
		}

		if w.Code != tc.responseCode {
			t.Errorf("expected code %d but got %d", tc.responseCode, w.Code)
		}

		if w.Body.String() != tc.body() {
			t.Errorf("expected error message %q, got %q", tc.body(), w.Body.String())
		}
	}
}

