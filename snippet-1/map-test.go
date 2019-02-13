package main

import (
	"fmt"
	"sort"
)

func main()  {
	var m map[int]string
	m = map[int]string{}
	m1 := make(map[int]string)
	var m2 map[int]string = make(map[int]string)
	m3 := make(map[int]string)
	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	m[1] = "OK"
	a := m[1]
	fmt.Println(a)

	delete(m,1)
	fmt.Println(a)

	var m4 map[int]map[int]string
	m4 = make(map[int]map[int]string)
	m4[1] = make(map[int]string)//important
	m4[1][1] = "OK"
	b := m4[1][1]
	c,ok := m4[2][1]
	if !ok {
		m4[2] = make(map[int]string)
	}
	m4[2][1] = "GOOD"
	d := m4[2][1]
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	/*
	for i,v:=range slice {
		i 是索引
	    v 是slice 值的拷贝
	}
	for k,v:= range map{
	}
	*/
	sm := make([]map[int]string,5)
	for i :=range sm{
		sm[i] = make(map[int]string,1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	//map 无序，间接排序
	mt :=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	//fmt.Println(len(mt))
	st :=make([]int,len(mt))
	//fmt.Println(st)
	i :=0
	for k,_ := range mt{
		st[i] = k
		i++
	}
	fmt.Println(st)
	//排序 key
	sort.Ints(st)
	fmt.Println(st)
	//然后根据key 取map的值

	m5 := map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	m6 := map[string]int{}
	for k,v :=range m5{
		m6[v] = k
	}
	fmt.Println(m6)

}