package list

type List[K any] struct {
	slice *[]K
}

func New[K any]() List[K] {
	return List[K]{
		slice: &[]K{},
	}
}

func (l List[K]) Add(item ...K) {
	l.AddAll(item)
}

func (l List[K]) Get(index int) K {
	return (*l.slice)[index]
}

func (l List[K]) AddAll(anotherList []K) {
	*l.slice = append(*l.slice, anotherList...)
}

func (l List[K]) Size() int {
	return len(*l.slice)
}

func (l List[K]) ForEach(predicate func(index int, item K)) {
	for i, k := range *l.slice {
		predicate(i, k)
	}
}

// Sort

func (l List[K]) Reverse() {
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
