# noGCMapIndex

对于大型map，比如总数达到千万级别的map,如果键或者值中包含引用类型(string类型，结构体类型，或者任何基本类型+指针的定义 *int, *float 等)，那么这个MAP在垃圾回收的时候就会非常慢，GC的周期回收时间可以达到秒级。

所以对于这种map需要进行优化，把复杂的不利于GC的复杂map转化为基础类型的  map[uint64]int+外部二级索引切片的形式。比如 map[string]intercace{} 转换为 map[uint64]int+[]intercace{}的形式，变成这种形式之后，整个gc基本就不耗时了。

注意，这种做法主要适用于单次加载完后，键值对不再变化的情况。对于键值对在运行过程中还要动态增减的情况则不适合。
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
		log.Println("key:",key,"对应的值为:",data[index])
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
		log.Println("key:",key,"对应的值为:",data[index])
	} else {
		log.Println(key, "不存在")
	}
}
```
