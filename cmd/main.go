package main

import (
	"DI-container/pkg"
	"fmt"
)

type I interface {
	Get() string
}

type Ia struct{}

func (Ia) Get() string { return "" }

type Ib struct{}

func (Ib) Get() string { return "" }

type A struct {
	S    []I `inject:"Ia,Ib"`
	test *B  `inject:"B"`
}

func (a *A) SetTest(test *B) {
	a.test = test
}

type B struct {
	C string
	D *D `inject:"D"`
}

type D struct {
	E string
}

func main() {
	container := pkg.NewContainer()
	a := A{}
	b := B{C: "asd"}
	d := D{"dsa"}
	ia := Ia{}
	ib := Ib{}
	container.Put(&a, &b, &d, ia, ib)
	container.Init()
	bean := container.Get(A{}).(*A)
	fmt.Print("C: ", bean.test.C)
}
