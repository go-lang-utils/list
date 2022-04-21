package arraylist

type arrayList[K any] struct {
	slice *[]K
}

func New[K any](initValues ...K) arrayList[K] {
	return arrayList[K]{
		slice: &initValues,
	}
}

func (l arrayList[K]) Add(item ...K) {
	l.AddAll(item)
}

func (l arrayList[K]) Get(index int) K {
	return (*l.slice)[index]
}

func (l arrayList[K]) Remove(index int) {
	s := *l.slice
	*l.slice = append(s[:index], s[index+1:]...)
}

func (l arrayList[K]) AddAll(anotherarrayList []K) {
	*l.slice = append(*l.slice, anotherarrayList...)
}

func (l arrayList[K]) Size() int {
	return len(*l.slice)
}

func (l arrayList[K]) IsEmpty() bool {
	return len(*l.slice) == 0
}

func (l arrayList[K]) ForEach(predicate func(index int, item K)) {
	for i, k := range *l.slice {
		predicate(i, k)
	}
}

func (l arrayList[K]) Reverse() {
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

func (l arrayList[K]) ToSlice() []K {
	return *l.slice
}
