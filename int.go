// Copyright 2022 rateLimit Author(https://github.com/yudeguang/noGCMapIndex). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/noGCMapIndex.
package noGCMapIndex

import "strconv"

type NoGCMapIndexInt struct {
	usedDataLen int //外部存放数据的切片中当前已存放数据的数量
	m           map[int]int
}

//需要预先确定大致数量，在不知道精确值的情况下，一般可以设定一个比实际数量稍大一点的值
func NewInt(length int) *NoGCMapIndexInt {
	var n NoGCMapIndexInt
	n.m = make(map[int]int, length)
	return &n
}

//通过KEY生成外部切片的下标
func (n *NoGCMapIndexInt) CreateIndex(key int) int {
	_, exist := n.m[key]
	if exist {
		panic("can't add the key '" + strconv.Itoa(key) + "' for twice")
	}
	n.m[key] = n.usedDataLen + 1
	n.usedDataLen = n.usedDataLen + 1
	return n.usedDataLen - 1
}

//通过KEY获得外部切片的下标
func (n *NoGCMapIndexInt) GetIndex(key int) int {
	if v, exist := n.m[key]; exist {
		return v
	}
	return -1
}
