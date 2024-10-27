package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

type Key interface {
	string | int
}

type HashItem[K Key] struct {
	key   K
	value any
	next  *HashItem[K]
}

type HashMap[K Key] struct {
	table     []*HashItem[K]
	length    int
	maxLength int
}

func newHashMap[K Key]() *HashMap[K] {
	return &HashMap[K]{
		maxLength: 101,
		table:     make([]*HashItem[K], 101),
	}
}

func newHashItem[K Key](key K, val any) *HashItem[K] {
	return &HashItem[K]{
		key:   key,
		value: val,
	}
}

func (h *HashMap[K]) set(key K, val any) {
	item := newHashItem(key, val)
	idx := h.hashIdx(key)

	if h.table[idx] == nil {
		h.table[idx] = item
	} else {
		curr := h.table[idx]
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = item
	}

	h.length++
}

func (h *HashMap[K]) get(key K) any {
	idx := h.hashIdx(key)

	for curr := h.table[idx]; curr != nil; curr = curr.next {
		if curr.key == key {
			return curr.value
		}
	}

	fmt.Printf("%v not found\n", key)
	return nil
}

func (h *HashMap[K]) hashIdx(key K) int {
	hash := fnv.New64a()
	var keyByte []byte

	switch k := any(key).(type) {
	case int:
		keyByte = []byte(strconv.Itoa(k))
	case string:
		keyByte = []byte(k)
	default:
		panic("Key must be an int or a string")
	}

	for i := 0; i < 3; i++ {
		_, err := hash.Write(keyByte)

		if err != nil {
			fmt.Println("Failed to create hash")
			continue
		}

		hashVal := hash.Sum64()
		return int(hashVal % uint64(h.maxLength))
	}

	fmt.Println("Failed to create hash value 3 times. Using 0 as the key")
	return 0
}

func (h *HashMap[K]) reHash() {
}

func (h *HashMap[K]) keys() []K {
	keys := make([]K, 0, h.length)
	for _, item := range h.table {
		for curr := item; curr != nil; curr = curr.next {
			keys = append(keys, curr.key)
		}
	}

	return keys
}

func (h *HashMap[K]) values(key K, val any) {
}

func (h *HashMap[K]) entries(key K, val any) {
}

func main() {
	hashMap := newHashMap[int]()
}
