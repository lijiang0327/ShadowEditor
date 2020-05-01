// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package helper

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

func TestDownload(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Download(w, "./download.go", "下载")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	strlen := res.Header.Get("Content-Length")

	t.Logf("Content-Disposition: %v", res.Header.Get("Content-Disposition"))
	t.Logf("Content-Length: %v", strlen)
	t.Logf("Content-Type: %v", res.Header.Get("Content-Type"))

	stat, err := os.Stat("./download.go")
	if err != nil {
		t.Error(err)
		return
	}

	size := int(stat.Size())
	length, _ := strconv.Atoi(strlen)

	if size != length {
		t.Errorf("size should be %v, get %v", size, length)
	}
}
