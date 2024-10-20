package main

import (
	// "fmt"
	"hash/fnv"
	"strconv"
)

type Key interface {
	string | int
}

type Value interface {
	string | int
}

type ListNode[K Key, V Value] struct {
	key  K
	val  V
	next *ListNode[K, V]
}

type HashMap[K Key, V Value] struct {
	maxLen  uint64
	hashMap []ListNode[K, V]
}

func newHashMap[K Key, V Value]() *HashMap[K, V] {
	const maxLen uint64 = 31
	return &HashMap[K, V]{
		maxLen:  maxLen,
		hashMap: make([]ListNode[K, V], maxLen),
	}
}

func (H *HashMap[K, V]) insert(key K, value V) {
	idx := H._hashIdx(key)
	item := H.hashMap[idx]

	if item.key == key {
	}
}

func (H *HashMap[K, V]) get() {
}

func (H *HashMap[K, V]) _hashIdx(key K) uint64 {
	hash := fnv.New64()
	var keyBytes []byte

	switch v := any(key).(type) {
	case int:
		keyBytes = []byte(strconv.Itoa(v))
	case string:
		keyBytes = []byte(v)
	default:
		panic("unsupported type")
	}

	hash.Write(keyBytes)
	return hash.Sum64() % H.maxLen
}

func (H *HashMap[K, V]) _reHash() {
}

func main() {
	// hashMap := newHashMap[string, int]()
}
