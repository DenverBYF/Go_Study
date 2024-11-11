package structture

type MyListNode struct {
	Val   int
	Key   int
	Next  *MyListNode
	Front *MyListNode
}

type LRUCache struct {
	Cap  int
	Size int
	Hash map[int]*MyListNode
	Head *MyListNode
	Tail *MyListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Size: 0,
		Hash: make(map[int]*MyListNode),
		Head: nil,
		Tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.Hash[key]; ok {
		this.moveToHead(val)
		return val.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.Hash[key]; ok {
		val.Val = value
		this.moveToHead(val)
		return
	}
	node := &MyListNode{
		Val: value,
		Key: key,
	}
	this.Hash[key] = node
	this.addToHead(node)
	this.Size++
	if this.Size > this.Cap {
		tail := this.deleteTail()
		delete(this.Hash, tail)
		this.Size--
	}
}

func (this *LRUCache) addToHead(node *MyListNode) {
	if this.Head == nil {
		this.Head = node
		this.Tail = node
		return
	}
	node.Next = this.Head
	this.Head.Front = node
	this.Head = node
}

func (this *LRUCache) moveToHead(node *MyListNode) {
	if this.Head == node {
		return
	}
	if this.Tail == node {
		this.Tail = node.Front
	} else {
		node.Next.Front = node.Front
	}
	node.Front.Next = node.Next
	node.Next = this.Head
	this.Head.Front = node
	this.Head = node
	node.Front = nil
}

func (this *LRUCache) deleteTail() int {
	v := this.Tail.Key
	if this.Tail.Front == nil {
		this.Tail = nil
		return v
	}
	this.Tail.Front.Next = nil
	this.Tail = this.Tail.Front
	return v
}
