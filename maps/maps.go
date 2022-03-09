package maps

import(
	"fmt"
)

const TableSize = 7
// Hashtable struct
type HashTable struct {
	keys [TableSize]*bucket
}
// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.keys[index].insert(key)
}
// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.keys[index].delete(key)
}
// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	in_table := h.keys[index].search(key)
	if in_table {
		fmt.Println(key,"is in the table")
	} else {
		fmt.Println(key,"is not in the table")
	}
	return in_table
}

// bucket Linked list struct
type bucket struct {
	head *node
}

type node struct {
	key string
	next *node
}

// insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &node{key: k}
		newNode.next = b.head
		b.head = newNode
	} else { 
		fmt.Println(k, "Already exists")
	}
}
// delete
func (b *bucket) delete(k string) {
	for i := b.head; i != nil; i = i.next {
		if i.key == k {
			*i = *i.next
		}
	}
}
// search
func (b *bucket) search(k string) bool{
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// hash function
func hash(key string) int {
	sum := 0
	for _, i := range key {
		sum += int(i)
	}
	return sum % TableSize
}

// constructor
func MakeMap() *HashTable {
	newmap := &HashTable{}
	for i := range newmap.keys {
		newmap.keys[i] = &bucket{}
	}
	return newmap
}
