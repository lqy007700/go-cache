package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element

	// 回调函数 可以在清楚数据后执行
	OnEvicted func(key string, val Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),

		OnEvicted: onEvicted,
	}
}

// Get
// 从字典中找到对应的双向链表节点
// 将节点移动到队尾
func (c *Cache) Get(key string) (Value, bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		v := ele.Value.(*entry)
		return v.value, true
	}
	return nil, false
}

// RemoveOldest
// 移除最少访问的节点
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)

		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())

		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add 添加
func (c *Cache) Add(key string, val Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(len(kv.key)) + int64(kv.value.Len())
		kv.value = val
	} else {
		ele := c.ll.PushFront(&entry{key: key, value: val})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(val.Len())
	}

	// 超出最大容量则淘汰
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
