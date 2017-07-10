package object_demo

import (
	"testing"
	"fmt"
	"log"
)

type Integer int

func (a Integer) Less(b Integer)  bool{
	return a < b
}
func (a Integer) Add(b Integer)  {
	a += b
}
func (a *Integer) Append(b Integer)  {
	*a += b
}

/**
给类添加方法
 */
func TestObjAddFunc(t *testing.T)  {
	var a Integer

	a = 23

	t.Log(a.Less(23))

	a.Add(22)
	t.Log(a)
	a.Append(22)
	t.Log(a)
}

type Rect struct {
	x,y float64
	width, height float64
}

func NewRect(x, y, width, height float64)  *Rect{
	fmt.Println("new Func")
	return &Rect{x, y, width, height}
}
/**
类初始化
 */
func TestObjInit(t *testing.T)  {
	rect := new(Rect)
	rect1 := &Rect{}
	rect2 := &Rect{1, 1, 200, 500}
	rect3 := &Rect{width: 299, height: 499}
	fmt.Println(rect.height)
	fmt.Println(rect1.height)
	fmt.Println(rect2.height)
	fmt.Println(rect3.height)
}

type Base struct {
	Name string
}

func (base *Base) Bar()  {
	fmt.Println("base bar")
}

type Foo struct {
	x int
	*Base
	*log.Logger
}

func (f *Foo) Start()  {
	f.Bar()
}


/**
类组合
 */
func TestObjCompoent(t *testing.T)  {
	f := new(Foo)
	f.Bar()
	
}

type Person struct {

}

func (p *Person) Walk() {
	fmt.Println("person walk")
}
func (p *Person) Eat()  {
	fmt.Println("person eat")
}
func (p *Person) Play()  {
	fmt.Println("person play")
}

type Pig interface {
	Walk()
	Eat()

}

func TestInterFac(t *testing.T)  {

}