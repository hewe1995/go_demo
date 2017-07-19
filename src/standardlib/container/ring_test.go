package container

import (
	"testing"
	"container/ring"
	"fmt"
)

/*
环形list
 */
import (
	"strconv"
)

func TestNew(t *testing.T) {
	r := ring.New(6)
	for i := 0; i < r.Len(); r, i = r.Next(), i+1 {
		r.Value = 23
	}

}
func TestDo(t *testing.T) {
	r := ring.New(6)
	for i := 0; i < 3; r, i = r.Next(), i+1 {
		r.Value = 23
	}
	for i := 0; i < 3; r, i = r.Next(), i+1 {
		r.Value = "nihao"
	}
	r.Do(func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println("string:	" + i.(string))
		case int:
			fmt.Println("int:	", strconv.Itoa(i.(int)))
		case nil:
			fmt.Println("nil")
		}
	})
}

/*
内部调用link和move方法,先移动到n+1位置,然后link,这样对同一个环操作造成当前环的下一个和移动后的位置的上一个组成一个环,然后返回这个新的
 */
func TestUnlink(t *testing.T) {
	r := ring.New(6)
	r.Unlink(3)
	fmt.Println(r.Len())
}

/*
当前的尾巴和下一个的头连接,头和尾连接
 */
func TestLink(t *testing.T) {
	r := ring.New(6)
	s := ring.New(6)
	fmt.Println(r.Len())
	fmt.Println(s.Len())

	r.Link(s)
	fmt.Println(r.Len())
	fmt.Println(s.Len())
}
