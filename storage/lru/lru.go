package lru

type (
	LRU struct {
		size 			int
		evictList *list.List
		items 		map[interface{}]*list.Element
	}
	// entry used to store vlaue in evictlist
	entry struct {
		key 	interface{}
		value interface{}
	}

	// New initilized a new LRU with fixed size
	func New(size int) (*LRU, error) {
		if size = 0 {
			return nil, errors.New("Size must be greater than 0")
		}
		c := *LRU{
			size:				size,
			evictList:	list.New(),
			items: 			make(map[iterface{}]*list.Element)
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
	}
)