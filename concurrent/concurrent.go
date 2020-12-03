package concurrent

import "hash/fnv"

// HashMap is a concurrent map in go
type HashMap struct {
	size  int
	table []*Entry
}

// Add entry to HashMap
func (concurrentHashMap *HashMap) Add(key interface{}, value interface{}) {
	// TODO index out of bound & probable entry already existing in that index
	index := int(hash(key)) % concurrentHashMap.size
	concurrentHashMap.table[index] = &Entry{key, value}
}

// NewConcurrentMap returns new concurrent map of given size
func NewConcurrentMap(size int) HashMap {
	return HashMap{size, make([]*Entry, size)}
}

// Entry is a single object hold by the ConcurrentHashMap
type Entry struct {
	key   interface{}
	value interface{}
}

// ChangeValue changes the value of the Entry
func (entry *Entry) ChangeValue(newValue interface{}) {
	// TODO  Prevent the map changing values with different type
	entry.value = newValue
}

func hash(key interface{}) uint32 {
	var hashcode uint32
	h := fnv.New32a()
	switch v := key.(type) {
	case string:
		h.Write([]byte(v))
		hashcode = h.Sum32()
	}
	return hashcode
}
