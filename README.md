# noGCMapIndex

对于大型map，比如总数达到千万级别的map,如果键或者值中包含引用类型，那么这个MAP在垃圾回收的时候就会非常慢，GC的周期回收时间可以达到秒级。
这里的引用类型包括string类型，结构体类型，或者任何基本类型+指针的定义（*int, *float）。
所以对于这种map需要进行优化，把诸如 map[string]intercace{} 转换为 map[uint64]int+[]intercace{}的形式
也就是说把复杂的不利于GC的map转化为基础类型的map+外部二级索引切片的形式。
