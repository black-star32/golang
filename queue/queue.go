package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}){
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{}{
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//// 强制限制类型
//func (q *Queue) Push(v int){
//	*q = append(*q, v)
//}
//
//func (q *Queue) Pop() int{
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head.(int)
//}

func (q *Queue) IsEmpyt() bool{
	return len(*q) == 0
}
