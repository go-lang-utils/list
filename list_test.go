package list

import (
	"testing"
)

func TestAddAndGet(t *testing.T) {
	l := New[int]()

	expectedVal := 1

	l.Add(expectedVal)

	val := l.Get(0)

	if val != expectedVal {
		t.Fatal("Expected", expectedVal, "but got", val)
	}
}

func TestAddAllAndSize(t *testing.T) {
	l := New[int]()

	appendArr := []int{1, 2, 3}

	l.AddAll(appendArr)

	if l.Size() != len(appendArr) {
		t.Fatal("Expected size", l.Size(), "but got", len(appendArr))
	}

	for i := 0; i < l.Size(); i++ {
		val := l.Get(i)
		expectedVal := appendArr[i]

		if val != expectedVal {
			t.Fatal("Expected", expectedVal, "but got", val)
		}
	}
}

func TestForEach(t *testing.T) {
	l := New[int]()

	l.Add(1)

	l.ForEach(func(index int, item int) {
		if l.Get(index) != item {
			t.Fatal()
		}
	})
}

func TestList_Reverse(t *testing.T) {
	l := New[int]()

	l.Add(5, 4, 3, 2, 1, 0)
	l.Reverse()

	for i := 0; i < l.Size(); i++ {
		val := l.Get(i)

		if val != i {
			t.Fatal("Expected", i, "but got", val)
		}
	}
}

func TestRemoveAndIsEmpty(t *testing.T) {
	l := New[int]()

	l.Add(1)

	l.Remove(0)

	if !l.isEmpty() {
		t.Fatal()
	}
}
