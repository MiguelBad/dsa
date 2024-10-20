package main

import (
// "fmt"
// "hash"
)

type ListNode[K any, V any] struct {
	key  K
	val  V
	next *ListNode[K, V]
}

type HashMap[K any, V any] struct {
	maxLen  int
	hashMap []ListNode[any, any]
}

func newHashMap(maxLen int) *HashMap[any, any] {
	return &HashMap[any, any]{
		maxLen:  31,
		hashMap: make([]ListNode[any, any], maxLen),
	}
}

func (H *HashMap[K, V]) insert(key K, value V) {
}

func (H *HashMap[K, V]) get() {
}

func (H *HashMap[K, V]) _hashIdx() {
}

func (H *HashMap[K, V]) _reHash() {
}

func main() {
}
