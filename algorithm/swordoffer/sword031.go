package swordoffer

type LRUCache struct {
	keyIndex map[int]int
	capacity int
	queue    []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		queue:    []int{},
	}
}

func (this *LRUCache) Get(key int) int {
	index := this.keyIndex[key]
	return this.queue[index]
}

func (this *LRUCache) Put(key int, value int) {

}
