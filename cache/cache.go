package cache

import (
	"fmt"
	"os"
	"strconv"
)

type LRUCache struct {
	capacity int
	cache    map[string]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key   string
	value string
	prev  *Node
	next  *Node
}

func NewLRUCache() *LRUCache {
	size := os.Getenv("CACHE_SIZE")
	capacity, err := strconv.Atoi(size)

	if err != nil {
		return nil
	}

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*Node, capacity),
	}
}

func (c *LRUCache) AddToFront(node *Node) {
	node.next = c.head
	node.prev = nil

	if c.head != nil {
		c.head.prev = node
	}

	c.head = node

	if c.tail == nil {
		c.tail = node
	}
}

func (c *LRUCache) MoveToFront(node *Node) {
	if node == c.head {
		return
	}

	if node == c.tail {
		c.tail = node.prev
		c.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	node.prev = nil
	node.next = c.head
	c.head.prev = node
	c.head = node
}

func (c *LRUCache) RemoveFromTail() {
	// cache is empty
	if c.tail == nil {
		return
	}

	if c.tail == c.head {
		//empty cache
		c.head = nil
		c.tail = nil
	} else {
		// set the current tail to the previous node
		c.tail = c.tail.prev
		// remove current tail's reference to the removed node
		c.tail.next = nil
	}
}

func (c *LRUCache) Get(key string) string {
	if node, ok := c.cache[key]; ok {
		c.MoveToFront(node)
		return node.value
	}
	return ""
}

func (c *LRUCache) Put(key string, value string) {
	if node, ok := c.cache[key]; ok {
		node.value = value
		c.MoveToFront(node)
	} else {
		if len(c.cache) >= c.capacity {
			delete(c.cache, c.tail.key)
			c.RemoveFromTail()
		}

		newNode := &Node{key: key, value: value}
		c.cache[key] = newNode
		c.AddToFront(newNode)
	}
}

func (c *LRUCache) PrintCurrentCache() {
	if c.head != nil && c.tail != nil {
		node := c.head

		for {
			fmt.Printf("{%s : %s}", node.key, node.value)
			if node.next == nil {
				fmt.Println()
				break
			}
			node = node.next
			fmt.Print(" -> ")
		}
	} else {
		fmt.Println("Cache is empty")
	}
}
