package sort

type SortIf []interface{}

type compare interface {
	Less(interface{}) bool
}

func (l SortIf) Len() int {
	return len(l)
}

func (l SortIf) Less(i, j int) bool {
	cmp := l[i].(compare)
	return (cmp).Less(l[j])
}

func (l SortIf) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
