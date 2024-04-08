package main

import "fmt"

type person struct {
	name string
	age  int
}

func (tmp person) print() {
	fmt.Println(tmp)
}

// 要改变得传指针
func (p *person) set_(mname string, mage int) {
	p.age = mage
	p.name = mname
}

func main() {
	a := person{"mike", 22}
	a.print()

	a.set_("titi", 88)
	a.print()

}
