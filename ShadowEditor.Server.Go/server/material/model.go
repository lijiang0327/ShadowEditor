// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package material

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Model 材质模型
type Model struct {
	// ID
	ID string
	// 名称
	Name string
	// 类别ID
	CategoryID string
	// 类别名称
	CategoryName string
	// 全拼
	TotalPinYin string
	// 首字母拼音
	FirstPinYin string
	// 创建时间
	CreateTime time.Time
	// 更新时间
	UpdateTime time.Time
	Data       bson.M
	// 缩略图
	Thumbnail string
}
