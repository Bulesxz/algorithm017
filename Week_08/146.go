package main

/*
146. LRU 缓存机制
*/
type LRUCache struct {
	m          map[int]*DLinkedNode
	capacity   int
	head, tail *DLinkedNode //头和尾使用两个节,不存数据
}
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		head:     initDLinkedNode(-1, -1),
		tail:     initDLinkedNode(-1, -1),
		m:        make(map[int]*DLinkedNode),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node := initDLinkedNode(key, value)
		if len(this.m) < this.capacity { //直接插入
			this.addToHead(node)
		} else {
			d := this.removeTail()
			delete(this.m, d.key)
			this.addToHead(node)
		}
		this.m[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
