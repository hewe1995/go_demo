package object_demo

import (
	"testing"
	"fmt"
	"log"
	"strconv"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a Integer) Add(b Integer) {
	a += b
}
func (a *Integer) Append(b Integer) {
	*a += b
}

/**
给类添加方法
 */
func TestObjAddFunc(t *testing.T) {
	var a Integer

	a = 23

	t.Log(a.Less(23))

	a.Add(22)
	t.Log(a)
	a.Append(22)
	t.Log(a)
}

type Rect struct {
	x, y          float64
	width, height float64
}

func NewRect(x, y, width, height float64) *Rect {
	fmt.Println("new Func")
	return &Rect{x, y, width, height}
}

/**
类初始化
 */
func TestObjInit(t *testing.T) {
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

func (base *Base) Bar() {
	fmt.Println("base bar")
}

type Foo struct {
	x int
	*Base
	*log.Logger
}

func (f *Foo) Start() {
	f.Bar()
}

/**
类组合
 */
func TestObjCompoent(t *testing.T) {
	f := new(Foo)
	f.Bar()

}

type Person struct {
}

func (p *Person) Walk() {
	fmt.Println("person walk")
}
func (p *Person) Eat() {
	fmt.Println("person eat")
}
func (p *Person) Play() {
	fmt.Println("person play")
}

type Pig interface {
	Walk()
	Eat()
}

type Engine struct {
	a, b int
}

func (e *Engine) SetA(a int) {
	fmt.Println("engine start")
	e.a = a
}
func (e *Engine) GetA() int {
	return e.a
}

type Car struct {
	Engine
	a int
}
type Format interface {
	String() string
}

func (c *Car) String() string {
	var str string
	str = "{a: " + strconv.Itoa(c.a) + "}"
	return str
}

type Comp interface {
	Less(b Integer) bool
	Add(b Integer)
	Append(b Integer)
}
func TestInterFac(t *testing.T) {
	var i Integer = 1
	var comp Comp = &i
	comp.Add(23)
	comp.Append(11)
}
