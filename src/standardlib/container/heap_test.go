/*
堆会对放进去的数据进行排序,具体排序规则自己实现,每次放进去后会进行排序吧最小或最大的放到开头,每次pop后取出最小或最大的,
堆中只有开头是最大或最小,其他数据不一定是顺序的.
 */
package container

import (
	"testing"
	"container/heap"
	"fmt"
)

//定义一个实现heap接口的类型
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (this *IntHeap) Push(x interface{}) {
	*this = append(*this, x.(int))
}
func (this *IntHeap) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n-1]
	*this = old[: n-1]

	return x
}

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}

	heap.Init(h)
	fmt.Print((*h)[0], "	")
	fmt.Println()
	heap.Push(h, 23)
	fmt.Print((*h)[0], "	")
	heap.Push(h, 0)
	fmt.Print((*h)[0], "	")
	fmt.Println()
	heap.Push(h, 3)

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
