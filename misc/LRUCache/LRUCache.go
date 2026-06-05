package LRUCache

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

// Constructor create new instance of LRUCache struct
func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) insertTail(node *Node) {
	node.next = c.tail
	node.prev = c.tail.prev
	c.tail.prev.next = node
	c.tail.prev = node
}

func (c *LRUCache) Get(key int) int {
	node, exist := c.cache[key]
	if !exist {
		return -1
	}
	c.remove(node)
	c.insertTail(node)
	return node.val
}

func (c *LRUCache) Put(key int, value int) {
	if node, exist := c.cache[key]; exist {
		node.val = value
		c.remove(node)
		c.insertTail(node)
		return
	}
	if len(c.cache) == c.capacity {
		lru := c.head.next
		c.remove(lru)
		delete(c.cache, lru.key)
	}
	newNode := &Node{key: key, val: value}
	c.insertTail(newNode)
	c.cache[key] = newNode
}
