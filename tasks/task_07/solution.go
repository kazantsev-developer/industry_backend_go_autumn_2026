package main

import (
	"container/list"
	"sync"
)

type LRU[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
}
type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       list.List
	items    map[K]*list.Element
}

func entryOf[K comparable, V any](elem *list.Element) *entry[K, V] {
	return elem.Value.(*entry[K, V])
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity <= 0 {
		return
	}

	elem, exists := c.items[key]
	if !exists {
		return
	}

	c.ll.MoveToBack(elem)
	return entryOf[K, V](elem).value, true
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity <= 0 {
		return
	}

	if elem, exists := c.items[key]; exists {
		entryOf[K, V](elem).value = value
		return
	}

	if len(c.items) >= c.capacity {
		front := c.ll.Front()
		delete(c.items, entryOf[K, V](front).key)
		c.ll.Remove(front)
	}

	elem := c.ll.PushBack(&entry[K, V]{key: key, value: value})
	c.items[key] = elem
}
