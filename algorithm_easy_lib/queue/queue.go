package queue

import "github.com/opay-org/lib-common/xlog"

type Queue []interface{}

func (q *Queue) Push(i interface{}) {
	*q = append(*q, i)
	xlog.Debug("len = %v", len(*q))
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Empty() bool {
	return q.Len() == 0
}

func (q *Queue) Pop() interface{} {
	i := (*q)[0]
	*q = (*q)[1:]
	return i
}
