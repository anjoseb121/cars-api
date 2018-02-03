package lru

import (
	"container/list"
	"errors"
)

type (
	// LRU used to store last recently used
	LRU struct {
		size      int
		evictList *list.List
		items     map[interface{}]*list.Element
	}
	// entry used to store vlaue in evictlist
	entry struct {
		key   interface{}
		value interface{}
	}
)

// New initilized a new LRU with fixed size
func New(size int) (*LRU, error) {
	if size == 0 {
		return nil, errors.New("Size must be greater than 0")
	}
	c := &LRU{
		size:      size,
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element),
	}
	return c, nil
}

// Add adds a value to the cache. Return true if evection occured
func (l *LRU) Add(key, value interface{}) bool {
	if ent, ok := l.items[key]; ok {
		l.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return false
	}
	ent := &entry{key, value}
	entry := l.evictList.PushFront(ent)
	l.items[key] = entry
	evict := l.evictList.Len() > l.size
	if evict {
		l.removeOldest()
	}
	return evict
}

func (l *LRU) removeOldest() {
	ent := l.evictList.Back()
	if ent != nil {
		l.removeElement(ent)
	}
}
func (l *LRU) removeElement(e *list.Element) {
	l.evictList.Remove(e)
	kv := e.Value.(*entry)
	delete(l.items, kv.key)
}
