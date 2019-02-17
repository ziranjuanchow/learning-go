package main

import "fmt"

func main() {
	/*
	a,b := 9,8
	A(a,b)
	fmt.Println(a,b)

	s1 :=[]int{1,2,3,4,5}
	B(s1)
	fmt.Println(s1)

	c := 1
	C(&c)
	fmt.Println(c)

	d := D
	d()

	e := func(){
		fmt.Println("匿名函数")
	}
	e()

	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
	*/

	/*
	for i := 0 ; i < 5 ; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("+++++++++++++++++++++++++++++++++++++")
	for i := 0 ; i < 5 ; i ++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	*/

	//E()
	//F()
	//G()

	var fs = [4]func(){}//函数数组
	for i := 0 ; i < 4 ; i++{
		defer fmt.Println("defer i = ",i)
		defer func(){
			fmt.Println("defer_closure i = ", i)
		}()
		fs[i] = func() {
			fmt.Println("closure i =", i)//闭包
		}
	}
	for _,f := range fs{
		f()
	}
}
//引用传递，但是是值拷贝
func A(s ...int) {
	s[0] = 99
	s[1] = 88
	fmt.Println(s)
}
//直接传递slice是要改变调用出的slice的,是对slice的内存地址进行了拷贝
func B(s []int){
	s[0] = 99
	s[1] = 98
	s[2] = 97
	s[3] = 96
	s[4] = 95
	fmt.Println(s)
}
//传地址
func C(c *int){
	*c = 2
	fmt.Println(c)
}
//
func D(){
	fmt.Println("Call D")
}
//
func closure(x int) func(int) int{
	return func(y int) int {
		return x+y
	}
}

func E(){
	fmt.Println("Func E")
}
func F(){
	defer func(){
		if err:=recover();err != nil {
			fmt.Println("Recover F")
		}
	}()
	panic("Panic F")//执行到这里 程序就已经停止了
}
func G(){
	fmt.Println("Func G")
}