package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	if c, ok := g.m[key]; ok {
		c.wg.Wait() // 正在请求中等待
		return c.val, c.err
	}

	c := &call{}
	c.wg.Add(1)
	g.m[key] = c

	c.val, c.err = fn()
	c.wg.Done()

	delete(g.m, key)
	return c.val, c.err
}
