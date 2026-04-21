package main

type LinkedList struct {
	head *Node
	tail *Node
}
type Node struct {
	prev  *Node
	next  *Node
	value int
	key   int
}

func InitLinkedList() LinkedList {
	headNode := Node{nil, nil, 0, 0}
	tailNode := Node{&headNode, nil, 0, 0}
	headNode.next = &tailNode

	n := LinkedList{
		head: &headNode,
		tail: &tailNode,
	}
	return n
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	linkedList LinkedList
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity:   capacity,
		cache:      make(map[int]*Node, capacity),
		linkedList: InitLinkedList(),
	}

	return cache
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	if node == nil {
		return -1
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	tmp := this.linkedList.head.next

	node.prev = this.linkedList.head
	node.next = tmp

	this.linkedList.head.next = node
	tmp.prev = node

	return this.cache[key].value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		if len(this.cache) == this.capacity {
			delete(this.cache, this.linkedList.tail.prev.key)
			this.linkedList.tail.prev = this.linkedList.tail.prev.prev
			this.linkedList.tail.prev.next = this.linkedList.tail

		}

		tmp := this.linkedList.head.next
		newNode := &Node{
			next:  tmp,
			prev:  this.linkedList.head,
			value: value,
			key:   key,
		}

		this.linkedList.head.next = newNode
		tmp.prev = newNode
		this.cache[key] = newNode

	} else {
		node.value = value
		node.prev.next = node.next
		node.next.prev = node.prev

		tmp := this.linkedList.head.next

		node.prev = this.linkedList.head
		node.next = tmp

		this.linkedList.head.next = node
		tmp.prev = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
