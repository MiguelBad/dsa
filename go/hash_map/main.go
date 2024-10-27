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

type Entry[K Key] struct {
	key   K
	value any
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
	idx := h.hashIdx(key)
	curr := h.table[idx]

	if curr == nil {
		h.table[idx] = newHashItem(key, val)
		h.length++
		return
	}

	for curr != nil {
		if curr.key == key {
			curr.value = val
			return
		}

		if curr.next == nil {
			break
		}

		curr = curr.next
	}

	curr.next = newHashItem(key, val)
	h.length++
	return
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

func (h *HashMap[K]) delete(key K) {
	idx := h.hashIdx(key)
	curr := h.table[idx]

	if curr != nil && curr.key == key {
		h.table[idx] = curr.next
		h.length--
		return
	}

	var prev *HashItem[K]
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			h.length--
			return
		}

		prev = curr
		curr = curr.next
	}

	fmt.Printf("Key <%v> does not exist", key)
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

func (h *HashMap[K]) values() []any {
	values := make([]any, 0, h.length)
	for _, item := range h.table {
		for curr := item; curr != nil; curr = curr.next {
			values = append(values, curr.value)
		}
	}

	return values
}

func (h *HashMap[K]) entries() []Entry[K] {
	entries := make([]Entry[K], 0, h.length)
	for _, item := range h.table {
		for curr := item; curr != nil; curr = curr.next {
			entries = append(entries, Entry[K]{curr.key, curr.value})
		}
	}

	return entries
}

func main() {
	hashMap := newHashMap[int]()

	hashMap.set(1, 1)
	hashMap.set(2, 3)
	hashMap.set(4, 6)
	hashMap.set(6, 9)
	hashMap.set(8, 12)
	hashMap.set(10, 15)

	fmt.Println(hashMap.entries())

	hashMap.set(1, 3)
	hashMap.set(2, 5)
	hashMap.set(4, 8)
	hashMap.set(6, 11)
	hashMap.set(8, 14)
	hashMap.set(10, 17)

	fmt.Println(hashMap.entries())

	hashMap.delete(1)

	fmt.Println(hashMap.entries())
}
