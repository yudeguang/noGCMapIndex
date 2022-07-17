// Copyright 2022 rateLimit Author(https://github.com/yudeguang/noGCMapIndex). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/noGCMapIndex.

package noGCMapIndex

import (
	"github.com/cespare/xxhash"
)

type NoGCMapString struct {
	usedDataLen         int            //切片data中当前已存放数据的数量
	mapHasHashCollision map[string]int //int为外部切片的下标,string为存放有hash冲突的第2次出现的key,这个map一般来说是非常小的
	mapNoHashCollision  map[uint64]int //int为外部切片的下标,uint64为存储无hash冲突的key的hash值以及有hash冲突并且是第1次出现的key的hash值
}

//需要预先确定大致数量，在不知道精确值的情况下，一般可以设定一个比实际数量稍大一点的值
func NewString(length int) *NoGCMapString {
	var n NoGCMapString
	n.mapNoHashCollision = make(map[uint64]int, length)
	n.mapHasHashCollision = make(map[string]int)
	return &n
}

//通过KEY获得外部切片的下标
func (n *NoGCMapString) GetIndex(key string) int {
	if v, exist := n.mapHasHashCollision[key]; exist {
		return v
	}
	if v, exist := n.mapNoHashCollision[xxhash.Sum64([]byte(key))]; exist {
		return v
	}
	return -1
}

//通过KEY生成外部切片的下标
func (n *NoGCMapString) CreateIndex(key string) int {
	h := xxhash.Sum64([]byte(key))
	//说明有hash冲突,以前存储过一次
	_, exist := n.mapNoHashCollision[h]
	if exist {
		//尽可能的避免重复加载,如果在mapNoHashCollision加载过，确实也是无法检测的，但是如果加载了3次一定会被检测到
		if _, exist := n.mapHasHashCollision[key]; exist {
			panic("can't add the key '" + key + "' for twice")
		}
		n.mapHasHashCollision[key] = n.usedDataLen + 1
	} else {
		n.mapNoHashCollision[h] = n.usedDataLen + 1
	}
	n.usedDataLen = n.usedDataLen + 1
	return n.usedDataLen - 1
}
