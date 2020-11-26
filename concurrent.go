package main

import "hash/fnv"

// ConcurrentHashMap is a concurrent map in go
type ConcurrentHashMap struct {
	size  int
	table []*Entry
}

func newConcurrentMap(size int) ConcurrentHashMap {
	return ConcurrentHashMap{size, make([]*Entry, size)}
}

// Entry is a single object hold by the ConcurrentHashMap
type Entry struct {
	key   interface{}
	value interface{}
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
