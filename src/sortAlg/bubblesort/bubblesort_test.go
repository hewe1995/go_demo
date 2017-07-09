package bubblesort

import (
	"testing"
)

func TestBbleSort_test1(t *testing.T)  {
	values := []int{5, 4, 3, 2, 1}
	expect := []int{1, 2, 3, 4, 5}
	BubbleSort(values)
	for idx, v := range values{
		if v != expect[idx]{
			t.Error("not expect!")
			break
		}
	}
}
func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	expect := []int{1, 2, 3, 5, 5}
	BubbleSort(values)
	for idx, v := range values{
		if v != expect[idx]{
			t.Error("not expect!")
			break
		}
	}
}
func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	expect := []int{5}
	BubbleSort(values)
	for idx, v := range values{
		if v != expect[idx]{
			t.Error("not expect!")
			break
		}
	}
}
func TestBubbleSort4(t *testing.T) {
	values := []int{}
	expect := []int{}
	BubbleSort(values)
	for idx, v := range values{
		if v != expect[idx]{
			t.Error("not expect!")
			break
		}
	}
}