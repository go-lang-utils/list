package list

type List[K any] list[K]

type list[K any] struct {
	slice *[]K
}

func New[K any]() list[K] {
	return list[K]{
		slice: &[]K{},
	}
}

func (l list[K]) Add(item ...K) {
	l.AddAll(item)
}

func (l list[K]) Get(index int) K {
	return (*l.slice)[index]
}

func (l list[K]) Remove(index int) {
	s := *l.slice
	*l.slice = append(s[:index], s[index+1:]...)
}

func (l list[K]) AddAll(anotherList []K) {
	*l.slice = append(*l.slice, anotherList...)
}

func (l list[K]) Size() int {
	return len(*l.slice)
}

func (l list[K]) isEmpty() bool {
	return len(*l.slice) == 0
}

func (l list[K]) ForEach(predicate func(index int, item K)) {
	for i, k := range *l.slice {
		predicate(i, k)
	}
}

func (l list[K]) Reverse() {
	left := 0
	right := len(*l.slice) - 1

	for left < right {
		tmp := (*l.slice)[left]
		(*l.slice)[left] = (*l.slice)[right]
		(*l.slice)[right] = tmp
		left++
		right--
	}
}
