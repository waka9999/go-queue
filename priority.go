package queue

import "container/heap"

// 定义优先级队列
type PriorityQueue []*Item

// NewPriorityQueue 新建优先级队列
// cap 为队列容量
func NewPriorityQueue(cap int) *PriorityQueue {
	items := make(PriorityQueue, 0, cap)
	return &items
}

// 返回队列长度
func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

// 队列存储项比较
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

// 队列存储项目交换方式
func (pq *PriorityQueue) Swap(i, j int) {
	if i < 0 || j < 0 {
		return
	}

	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].setIndex(int64(i))
	(*pq)[j].setIndex(int64(j))
}


// 存放队列存储项
func (pq *PriorityQueue) Push(val interface{}) {
	n := len(*pq)
	c := cap(*pq)

	// 扩展队列容量
	if n+1 >= c {
		npq := make(PriorityQueue, n, c*2)
		copy(npq, *pq)
		*pq = npq
	}

	i := val.(*Item)
	i.setIndex(int64(n))

	*pq = append(*pq, i)
}

// 获取队列存储项
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	if n == 0 {
		return nil
	}
	c := cap(*pq)

	// 收缩队列容量
	if n < (c/3) && c > 10 {
		npq := make(PriorityQueue, n, c/2)
		copy(npq, *pq)
		*pq = npq
	}

	// 返回最后一个队列存储项
	i := (*pq)[n-1]

	if n == 1 {
		*pq = (*pq)[:0]
	} else {
		*pq = (*pq)[0 : n-1]
	}

	i.setIndex(-1)

	return i
}

// 返回队列中的首个存储项
func (pq *PriorityQueue) First() *Item {
	return (*pq)[0]
}

// 清空队列
func (pq *PriorityQueue) Clear() {
	for (*pq).Len() != 0 {
		heap.Pop(pq)
	}
}