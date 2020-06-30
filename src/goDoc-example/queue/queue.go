package queue

// 一个队列，可以存放任何类型的元素
type Queue []interface{}

// 向队列添加一个元素
func (q *Queue) Push(n interface{}) {
	*q = append(*q, n)
}

// 移除队列的第一个元素
func (q *Queue) Pop() interface{}  {
	head := (*q)[0]
	*q = (*q)[1:]
	return head

}
