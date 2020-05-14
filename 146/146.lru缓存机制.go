// package main

// import "fmt"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start

type Node struct{
	key int
	value int
	next *Node
	pre  *Node
}

type LRUCache struct {
		capacity int
		current int
		kvBinding map[int]*Node
		head *Node
}

func Constructor(capacity int) LRUCache {
		return LRUCache{
			capacity: capacity,
			current : 0,
			kvBinding: make(map[int]*Node),
			head : nil,
		}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.kvBinding[key];ok {
		if this.head != node {
			// 脱离 node
			node.pre.next = node.next
			node.next.pre = node.pre

			// 修改node指针
			node.next = this.head
			node.pre = this.head.pre

			// 修改队首和队尾的指针
			this.head.pre.next = node		
			this.head.pre = node

			// 修改队首
			this.head = node
		}
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.kvBinding[key]; ok{
		node.value = value
		if this.head != node {
			node.pre.next = node.next
			node.next.pre = node.pre
			node.next = this.head
			node.pre = this.head.pre
			this.head.pre.next = node		
			this.head.pre = node
			this.head = node
		}
	} else {
		n := &Node{
			key : key,
			value : value,
			pre : nil,
			next : nil,
		}
		if this.current == 0 {
			n.pre = n
			n.next = n
			this.head = n
			this.current = 1
		}else {
			n.pre = this.head.pre
			n.next = this.head
			this.head.pre.next = n
			this.head.pre = n
			this.head = n
			this.current = this.current + 1
			if this.current > this.capacity {
					shouldDelete := this.head.pre
					this.head.pre.pre.next = this.head
					this.head.pre = this.head.pre.pre
					this.current = this.current - 1
					delete(this.kvBinding, shouldDelete.key)
			}
		}
		this.kvBinding[key] = n
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end


// func main(){
// 	cache := Constructor(2)
// 	cache.Put(2,1)
// 	cache.Put(3,2)
// 	fmt.Println(cache)
// 	fmt.Println(cache.Get(3))
// 	fmt.Println(cache)
// 	fmt.Println(cache.Get(2))
// 	fmt.Println(cache)
// 	cache.Put(4,3)
// 	fmt.Println(cache)
// 	fmt.Println(cache.Get(2))
// 	fmt.Println(cache.Get(3))
// 	fmt.Println(cache.Get(4))
	// cache.Put(4,4)
	// fmt.Println(cache)
	// fmt.Println(cache.Get(3))
	// fmt.Println(cache.Get(2))
	// fmt.Println(cache.Get(1))
	// fmt.Println(cache)
	// cache.Put(5,5)
	// fmt.Println(cache.Get(1))
	// fmt.Println(cache.Get(2))
	// fmt.Println(cache.Get(3))
	// fmt.Println(cache.Get(4))
	// fmt.Println(cache.Get(5))
// }