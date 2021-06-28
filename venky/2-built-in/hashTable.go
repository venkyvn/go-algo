package main

import "fmt"

type hashTable struct {
	size    int
	buckets [][]bucket
}

type bucket struct {
	key, value string
}

type isHashTable interface {
	hash(key string) int

	get(key string) interface{}
	put(key string, value string)
}

func NewHashTable(size int) *hashTable {
	return &hashTable{
		size:    size,
		buckets: make([][]bucket, size),
	}
}

func main() {
	table := NewHashTable(2)
	table.put("vuong", "toi la vuong")
	table.put("nguyen", "toi la nguyen")
	table.put("nguyen2", "toi la nguyen2")
	table.put("dan", "toi la dan")
	table.put("my", "toi la my")

	fmt.Println("total", table.buckets)

	fmt.Println(table.get("vuong"))
	fmt.Println(table.get("nguyen"))
	fmt.Println(table.get("dan"))
	fmt.Println(table.get("my"))

}

func (h *hashTable) hash(key string) int {
	hash := 0
	keyChar := []rune(key)
	for i := 0; i < len(keyChar); i++ {
		hash = (hash + int(keyChar[i])*i) % h.size
	}
	return hash
}

func (h *hashTable) get(key string) interface{} {
	address := h.hash(key)

	currentBucket := h.buckets[address]

	if len(currentBucket) > 0 {
		for _, b := range currentBucket {
			if b.key == key {
				return b.value
			}
		}
	}
	return nil
}

func (h *hashTable) put(key string, value string) {
	address := h.hash(key)
	newItem := bucket{key, value}
	bucket := h.buckets[address]
	bucket = append(bucket, newItem)
	h.buckets[address] = bucket
}
