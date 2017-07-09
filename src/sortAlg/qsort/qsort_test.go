package qsort

import (
	"testing"
	"fmt"
)

func TestQsort(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	expect := []int{1, 2, 3, 4, 5}
	for _, v := range values{
		fmt.Print(v, "	")
	}
	fmt.Println()
	Qsort(values)
	for _, v := range values{
		fmt.Print(v, "	")
	}
	for idx, v := range values{
		if v != expect[idx]{
			t.Error("not expect!")
			break
		}
	}

}

func TestQsort1(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	expect := []int{1, 2, 3, 5, 5}
	for _, v := range values{
		fmt.Print(v, "	")
	}
	fmt.Println()
	Qsort(values)
	for _, v := range values{
		fmt.Print(v, "	")
	}
	for idx, v := range values{
		if v != expect[idx]{
			t.Error("not expect!")
			break
		}
	}

}
func TestQsort2(t *testing.T) {
	values := []int{5}
	expect := []int{5}
	for _, v := range values{
		fmt.Print(v, "	")
	}
	fmt.Println()
	Qsort(values)
	for _, v := range values{
		fmt.Print(v, "	")
	}
	for idx, v := range values{
		if v != expect[idx]{
			t.Error("not expect!")
			break
		}
	}

}