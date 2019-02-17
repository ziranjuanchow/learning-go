package main

import "fmt"

type TZ int

type IN int

type A struct {
	Name string
}

type B struct {
	Name string
}

func main(){
	a:= A{}
	a.Print()

	b:= B{}
	b.Print()

	var c TZ
	c.Print()//method value
	(*TZ).Print(&c)//method exp

	var in IN
	in = 0
	in.Increase(100)
	fmt.Println(in)

	in1 := IN(0)
	in1.Increase(200)
}

func (a1 A) Print() {
	fmt.Println("A")
}

func (b B) Print() {
	fmt.Println("B ")
}

func (c *TZ) Print(){
	fmt.Println("TZ")
}

func (in *IN) Increase(num int){
	*in += IN(num)
	fmt.Println(*in)
}