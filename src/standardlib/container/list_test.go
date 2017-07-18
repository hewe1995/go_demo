/*
测试list,list是一个链表,不支持数组的下标索引
 */
package container

import (
	"testing"
	"container/list"
	"fmt"
)
/*
添加
 */
func TestAdd(t *testing.T)  {
	lis := list.New()

	//从后面添加
	lis.PushBack(1)
	lis.PushBack(2)
	lis.PushBack(3)
	lis.PushBack(4)
	//从前面添加
	lis.PushFront(0)
	lis.PushFront(-1)
	lis.PushFront(-2)
	//添加一个list到当前list后面
	lis2 := list.New()
	lis2.PushBack(11)
	lis2.PushBack(12)

	lis.PushBackList(lis2)

	//添加一个list到当前list前面
	lis.PushFrontList(lis2)

	for i := lis.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

/*
删除
 */
func TestRemove(t *testing.T)  {
	lis := list.New()

	//从后面添加
	elem := lis.PushBack(1)
	lis.PushBack(2)
	lis.PushBack(3)
	lis.PushBack(4)

	//删除,删除的是element不是value
	lis.Remove(elem)
	for i := lis.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}