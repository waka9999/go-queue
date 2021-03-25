// queue 包提供支持优先级的队列
package queue

import "sync/atomic"

// Item 队列存储项
type Item struct {
	// 实际存储值
	value interface{}

	// 优先级
	priority int64

	// 索引值
	index int64
}

// NewItem 新建队列存储项
func NewItem(value interface{}, priority int64) *Item {
	return &Item{
		value:    value,
		priority: priority,
		index:    -1,
	}
}

// Value 获取队列存储项中的实际值
func (i *Item) Value() interface{} {
	return i.value
}

// Index 读取存储项的索引值
func (i *Item) Index() int64 {
	return atomic.LoadInt64(&i.index)
}

// Priority 获取存储项的优先级
func (i *Item) Priority() int64 {
	return i.priority
}

// 设置存储项的索引值
func (i *Item) setIndex(index int64) {
	atomic.StoreInt64(&i.index, index)
}
