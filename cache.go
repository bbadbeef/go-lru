package cache

import (
	"fmt"
	"sync"
)

type Item struct {
	value interface{}
}

type Cache struct {
	sync.Mutex
	d      map[interface{}]*Item
	l      *List
	cnt    uint
	maxCnt uint
}

func NewCache(mc uint) *Cache {
	c := &Cache{
		d:      make(map[interface{}]*Item, mc/10),
		l:      NewList(),
		cnt:    0,
		maxCnt: mc,
	}
	return c
}

func (c *Cache) Set(key interface{}, value interface{}) {
	c.Lock()
	v, ok := c.d[key]
	if ok {
		v.value = value
		c.l.update(key)
	} else {
		i := &Item{
			value: value,
		}
		c.d[key] = i
		c.l.add(key)
		c.cnt++
		if c.cnt > c.maxCnt {
			k := c.l.cut()
			delete(c.d, k)
			c.cnt--
		}
	}
	c.Unlock()
}

func (c *Cache) Get(key interface{}) interface{} {
	c.Lock()
	defer c.Unlock()
	v, ok := c.d[key]
	if !ok {
		return nil
	}
	c.l.update(key)
	return v.value
}

func (c *Cache) Del(key interface{}) {
	c.Lock()
	_, ok := c.d[key]
	if ok {
		delete(c.d, key)
		c.l.del(key)
		c.cnt--
	}
	c.Unlock()
}

func (c *Cache) debug() {
	fmt.Println("cache size: ", len(c.d))
	for k, _ := range c.d {
		fmt.Println("map key: ", k)
	}
}
