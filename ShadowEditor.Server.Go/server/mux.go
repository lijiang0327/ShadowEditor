// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package server

import (
	"net/http"

	"github.com/dimfeld/httptreemux"

	"github.com/tengge1/shadoweditor/helper"
)

var (
	// Mux is a tree mux.
	Mux *httptreemux.TreeMux
)

func init() {
	mux := httptreemux.New()
	mux.OptionsHandler = corsHandler
	Mux = mux
}

func corsHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	helper.EnableCrossDomain(w, r)
}
