// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package scene

import "time"

// Model 场景模型
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
	// 表名
	CollectionName string
	// 版本号
	Version int
	// 创建时间
	CreateTime time.Time
	// 更新时间
	UpdateTime time.Time
	// 缩略图
	Thumbnail string
	// 是否公开
	IsPublic bool
}
