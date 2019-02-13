package main

import "fmt"

func main(){
	a := [2]int{1}
	b := [20]int{19:1}
	c := [...]int{1,2,3,4,5,6}
	d := [...]int{0:1,1:2,2:3,4:99}
	e := [...]int{99:1}
	var p *[100]int = &e

	x,y := 1,2
	f := [...]*int{&x,&y}

	g := new([10]int)//返回指向数组的指针
	g[1] = 99

	h := [2][3]int{
		{1,2,3},
		{3,2,1}}
	//非顶级的维度，长度不能用"..."

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(p)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
