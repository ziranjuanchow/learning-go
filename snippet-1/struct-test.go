package main

import "fmt"

type person struct {
	Name string
	Age int
}

type man struct {
	Name string
	Age int
	Contant struct{
		 Phone,City string
	}
}

type woman struct {
	string
	int
}//struct 匿名字段

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age int
}

type student struct {
	human
	Name string
	Age int
}

type teacherA struct {
	human
	Sex string
	Name string
	Age int
}

func main(){
	a := &person{
		Name:"Chow",
		Age: 19,
	}
	fmt.Println(a)
	Ap(a)
	fmt.Println(a)

	b := &struct {
		Name string
		Age int
	}{
		Name: "Chow",
		Age: 20,
	}//匿名结构体 初始化，取地址
	fmt.Println(b)
	c := man{Name:"Chow",Age:21}
	c.Contant.Phone = "123456"
	c.Contant.City = "BJ"
	fmt.Println(c)

	d := woman{"Wo",22}//顺序和对应类型重要
	var e = d
	fmt.Println(d)
	fmt.Println(e)

	f := teacher{Name:"Jim",Age:33,human:human{Sex:0}}
	g := student{Name:"Tom",Age:15,}
	g.human.Sex = 1
	g.Sex = 2
	fmt.Println(f)
	fmt.Println(g)

	h := teacherA{Sex:"F",Name:"Jeffy",Age:40,human:human{Sex:1}}
	h.Sex = "M"
	h.human.Sex = 0

	fmt.Println(h)
}

func Ap(per *person){
	per.Age = 13
	fmt.Println("Ap",per)
}