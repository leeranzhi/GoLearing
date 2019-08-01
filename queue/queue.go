package queue

//扩充系统类型和别人的类型

//1.定义别名
//2.使用组合

//队列的模拟实现
//interface表示任何类型
type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
