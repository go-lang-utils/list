package list

type List[K any] interface {
	Add(item ...K)

	Get(index int) K

	Remove(index int)

	AddAll(anotherList []K)

	Size() int

	IsEmpty() bool

	ForEach(predicate func(index int, item K))

	Reverse()

	ToSlice() []K
}
