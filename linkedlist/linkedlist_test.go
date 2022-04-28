package linkedlist

import (
	"log"
	"testing"
)

func TestLinkedList_Size(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	list.Add(listSlice...)

	if list.Size() != len(listSlice) {
		t.Fatalf("list.Size() should return %d", len(listSlice))
	}
}

func TestLinkedListInstantiation(t *testing.T) {
	list := New[int]()

	if list.head != nil {
		t.Fatal("list.head should be nil initially")
	}
	if list.tail != nil {
		t.Fatal("list.tail should be nil initially")
	}
	if list.length != 0 {
		t.Fatalf("list.length should initially be %d", 0)
	}
}

func TestLinkedList_Add(t *testing.T) {
	list := New[int]()

	// check with single node
	list.Add(1)
	if list.head.Val != 1 {
		t.Fatalf("list.head.Val should equal %d", 1)
	}
	if list.head.Next != nil {
		t.Fatal("list.head.Next should be nil")
	}
	if list.tail.Val != 1 {
		t.Fatalf("list.tail.Val should equal %d", 1)
	}
	if list.tail.Next != nil {
		t.Fatal("list.head.Next should be nil")
	}
	if list.length != 1 {
		t.Fatalf("list.length should be %d", 1)
	}
}

func TestLinkedList_AddTwo(t *testing.T) {
	list := New[int]()

	// check with two nodes
	list.Add(1, 2)
	if list.head.Val != 1 {
		t.Fatalf("list.head.Val should equal %d", 1)
	}
	if list.head.Prev != nil {
		t.Fatal("list.head.Prev should be nil")
	}
	if list.head.Next != list.tail {
		t.Fatal("list.head.Next should be list.tail")
	}

	// check tail val, prev and next
	if list.tail.Val != 2 {
		t.Fatalf("list.tail.Val should equal %d", 2)
	}
	if list.tail.Prev != list.head {
		t.Fatal("list.tail.Prev should be list.head")
	}
	if list.tail.Next != nil {
		t.Fatal("list.tail.Next should be nil")
	}
}

func TestLinkedList_AddMany(t *testing.T) {
	list := New[int]()

	// check with 3+ nodes
	listSlice := []int{1, 2, 3, 4, 5}
	list.Add(listSlice...)
	if list.head.Val != listSlice[0] {
		t.Fatalf("list.head.Val should equal %d", listSlice[0])
	}
	if list.head.Prev != nil {
		t.Fatal("list.head.Prev should be nil")
	}
	if list.head.Next.Val != listSlice[1] {
		t.Fatalf("list.head.Next should be %d", listSlice[1])
	}

	// check tail val, prev and next
	if list.tail.Val != listSlice[len(listSlice)-1] {
		t.Fatalf("list.tail.Val should equal %d", listSlice[len(listSlice)-1])
	}
	if list.tail.Prev.Val != listSlice[len(listSlice)-2] {
		t.Fatalf("list.tail.Prev should be %d", listSlice[len(listSlice)-2])
	}
	if list.tail.Next != nil {
		t.Fatal("list.tail.Next should be nil")
	}
}

func TestLinkedList_Get(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	list.Add(listSlice...)

	if list.Get(0) != listSlice[0] {
		t.Fatalf("list.Get(0) should get the first element")
	}

	if list.Get(3) != listSlice[3] {
		t.Fatalf("list.Get(n) should get the nth element")
	}
}

func TestLinkedList_GetNotExists(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("list.Get should panic when accessing a non-existing index")
		}
	}()

	list := New[int]()
	list.Get(0)
}

func TestLinkedList_RemoveWithOne(t *testing.T) {
	list := New[int]()
	list.Add(1)

	list.Remove(0)

	if list.length != 0 {
		t.Fatalf("list.length should be %d after removal", 0)
	}
	if list.head != nil {
		t.Fatal("list.head should be nil")
	}
	if list.tail != nil {
		t.Fatal("list.tail should be nil")
	}
}

func TestLinkedList_RemoveWithTwo(t *testing.T) {
	list := New[int]()
	list.Add(1, 2)

	list.Remove(1)

	if list.length != 1 {
		t.Fatalf("list.length should be %d", 1)
	}

	if list.head != list.tail {
		t.Fatal("list.head should equal list.tail")
	}

	if list.head.Next != nil {
		t.Fatal("list.head.Next should be nil")
	}
	if list.head.Prev != nil {
		t.Fatal("list.head.Prev should be nil")
	}

	if list.tail.Next != nil {
		t.Fatal("list.tail.Next should be nil")
	}
	if list.tail.Prev != nil {
		t.Fatal("list.tail.Prev should be nil")
	}

	TestLinkedList_RemoveWithOne(t)
}

func TestLinkedList_RemoveWithMany(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	list.Add(listSlice...)

	list.Remove(2)

	if list.length != len(listSlice)-1 {
		t.Fatalf("list.length should be %d", len(listSlice)-1)
	}
	if list.Get(2) != listSlice[3] {
		t.Fatalf("list[%d] should equal listSlice[%d]", 2, 3)
	}

	TestLinkedList_RemoveWithTwo(t)
}

func TestLinkedList_AddAll(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	list.AddAll(listSlice)

	if list.length != len(listSlice) {
		t.Fatalf("list.length should be %d", len(listSlice))
	}

	if list.head.Val != listSlice[0] {
		t.Fatal("list.head.Val should equal first value in listSlice")
	}
	if list.tail.Val != listSlice[len(listSlice)-1] {
		t.Fatal("list.tail.Val should equal last value in listSlice")
	}

	TestLinkedList_Get(t)
}

func TestLinkedList_IsEmpty(t *testing.T) {
	list := New[int]()
	if !list.IsEmpty() {
		t.Fatal("list.IsEmpty() should return true for an empty list")
	}

	list2 := New[int]()
	list2.Add(1)
	if list2.IsEmpty() {
		t.Fatal("list.IsEmpty() should return false for a non-empty list")
	}
}

func TestLinkedList_ForEach(t *testing.T) {
	var listSliceTwo []int

	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	list.Add(listSlice...)

	list.ForEach(func(index int, item int) {
		listSliceTwo = append(listSliceTwo, item)
	})

	if len(listSliceTwo) != list.length {
		t.Fatalf("listSlice length should equal %d", list.length)
	}
	for index, item := range listSlice {
		expected := list.Get(index)
		if item != expected {
			t.Fatalf("listSlice[%d] should equal %d", index, expected)
		}
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	reversedSlice := []int{5, 4, 3, 2, 1}

	list.Add(listSlice...)
	list.Reverse()

	list.ForEach(func(index int, item int) {
		if item != reversedSlice[index] {
			t.Fatalf("list[%d] should equal %d after reversing", index, reversedSlice[index])
		}
	})
}

func TestLinkedList_ToSlice(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	list.Add(listSlice...)

	listSliceActual := list.ToSlice()

	for i, expected := range listSlice {
		if listSliceActual[i] != expected {
			t.Fatalf("listSliceActual[%d] should match listSlice[%d]", i, i)
		}
	}
}

func TestLinkedList_GetFirst(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	list.Add(listSlice...)

	if list.GetFirst() != 1 {
		t.Fatalf("list.getFirst should be %d", 1)
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 5}
	list.Add(listSlice...)

	if list.GetLast() != 5 {
		t.Fatalf("list.getFirst should be %d", 5)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	list := New[int]()
	listSlice := []int{1, 2, 3, 4, 6}

	list.Add(listSlice...)

	err := list.Insert(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = list.Insert(list.Size()-1, 5)
	if err != nil {
		log.Fatal(err)
	}

	for i, val := range list.ToSlice() {
		if i != val {
			log.Fatal("Expected ", i, " but got ", val)
		}
	}
}
