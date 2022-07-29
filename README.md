# noGCMapIndex

废弃，见项目见 https://github.com/yudeguang/noGcStaticMap
```go
package main

import (
	"github.com/yudeguang/noGCMapIndex"
	"log"
	"strconv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	tstring()
	tint()
}

func tstring() {
	//旧MAP
	type col struct {
		a string
		b string
		c int
	}
	oldMap := make(map[string]col)
	for i := 0; i < 10; i++ {
		oldMap[strconv.Itoa(i)] = col{strconv.Itoa(i), strconv.Itoa(i), i}
	}

	//旧MAP转换为新形式的map+二级索引切片的形式
	data := make([]col, len(oldMap))
	m := noGCMapIndex.NewString(len(oldMap))
	for k, v := range oldMap {
		index := m.CreateIndex(k)
		data[index] = v
	}
	//查询数据
	key := "3"
	index := m.GetIndex(key)
	if index != -1 {
		log.Println("key:", key, "对应的值为:", data[index])
	} else {
		log.Println(key, "不存在")
	}
}

func tint() {
	//旧MAP
	type col struct {
		a string
		b string
		c int
	}
	oldMap := make(map[int]col)
	for i := 0; i < 10; i++ {
		oldMap[i] = col{strconv.Itoa(i), strconv.Itoa(i), i}
	}

	//旧MAP转换为新形式的map+二级索引切片的形式
	data := make([]col, len(oldMap))
	m := noGCMapIndex.NewInt(len(oldMap))
	for k, v := range oldMap {
		index := m.CreateIndex(k)
		data[index] = v
	}
	//查询数据
	key := 3
	index := m.GetIndex(key)
	if index != -1 {
		log.Println("key:", key, "对应的值为:", data[index])
	} else {
		log.Println(key, "不存在")
	}
}

```
