package Polymorphism

import "fmt"

type Shape interface {
	Area()
}

type Circle struct {
	r float64
}

func (c *Circle) SetR(r float64) {
	c.r = r
}

func (c Circle) Area() {
	fmt.Println(3.14 * c.r * c.r)
}

type Rectangle struct {
	l float64
	b float64
}

func (r *Rectangle) SetLB(l, b float64) {
	r.l = l
	r.b = b
}
func (r Rectangle) Area() {
	fmt.Println(r.l * r.b)
}
