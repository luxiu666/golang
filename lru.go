package main

import "fmt"

/*
通过map+双向链表实现lru缓冲算法
1 get操作：如果存在key，返回value并移动到头部，表示最新使用
          如果不存在key，返回-1
2 put操作：如果存在key，更新value并移动到头部，表示最新使用
          如果不存在，判断容量是否满，如果满了，删除尾部节点并添加新节点到头部

*/

// node节点
type Node struct {
	key   string
	value string
	prev  *Node
	next  *Node
}

// lru节点
type LRUCache struct {
	capacity int //容量
	cache    map[string]*Node
	head     *Node
	tail     *Node
}

// 初始化构造函数
func Init(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[string]*Node),
		head:     head,
		tail:     tail,
	}
}

// 添加节点
func (lru *LRUCache) addToHead(node *Node) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

// 删除节点
func (lru *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 移动到头部
func (lru *LRUCache) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// 删除尾部节点
func (lru *LRUCache) popTail() {
	lastNode := lru.tail.prev
	lru.removeNode(lastNode)
}

// get操作
func (lru *LRUCache) Get(key string) string {
	node, ok := lru.cache[key]
	if !ok {
		return "-1"
	} else {
		lru.moveToHead(node)
		return node.value
	}
}

// put 操作
func (lru *LRUCache) Put(key string, value string) {
	node, ok := lru.cache[key]
	if ok {
		node.value = value
		lru.moveToHead(node)

	} else {
		newNode := &Node{
			key:   key,
			value: value,
		}
		if len(lru.cache) >= lru.capacity {
			lru.popTail()
			delete(lru.cache, key)
		}
		lru.cache[key] = newNode
		lru.addToHead(newNode)
	}
}

// 显示当前缓存状态（用于调试）
func (lru *LRUCache) Display() {
	fmt.Print("缓存状态 (新→旧): ")
	current := lru.head.next
	for current != lru.tail {
		fmt.Printf("[%s:%s] ", current.key, current.value)
		current = current.next
	}
	fmt.Printf("(大小: %d/%d)\n", len(lru.cache), lru.capacity)
}

// 测试函数
func testLRUCache() {
	// 创建容量为2的LRU缓存
	lru := Init(2)
	fmt.Println("创建容量为2的LRU缓存")

	// 测试用例1
	fmt.Println("\n测试用例1:")
	lru.Put("p1", "java")
	lru.Display()

	lru.Put("p2", "c#")
	lru.Display()

	result := lru.Get("p2")
	lru.Display()
	fmt.Printf("get(p2) = %s\n", result)

}
