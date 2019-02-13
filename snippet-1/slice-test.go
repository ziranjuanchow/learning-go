package main

import "fmt"

func main(){
	var s1 [10]int
	fmt.Println(s1)

	a := [10]int{0,1,2,3,4,5,6,7,8,9 }
	s2 := a[5:]
	fmt.Println(s2)

	s3 := make([]int,3,10)
	fmt.Println(len(s3),cap(s3))


	b := []byte{'a','b','c','d','e','f','g','h','i','j','k'}
	s4 := b[2:5]
	s5 := b[3:5]
	s6 := s4[1:3]
	s7 := s4[3:5]
	fmt.Println(string(s4),string(s5))
	fmt.Println(string(s6))
	fmt.Println(string(s7))
	fmt.Println(len(s4),cap(s4))

	s8 := make([]int,3,6)
	fmt.Printf("%p\n",s8)
	s8 = append(s8, 1,2,3)
	fmt.Printf("%v %p\n",s8,s8)
	s8 = append(s8, 4,5,6)
	fmt.Printf("%v %p\n",s8,s8)

	c :=  []int{1,2,3,4,5}
	s9 := c[2:5]
	s10 := c[1:3]
	s11 := c[1:3]
	fmt.Println(s9,s10,s11)
	s11 = append(s11,1,1,1,1,1,1,1,1,1)
	s9[0] = 9
	fmt.Println(s9,s10,s11)

	s12 := []int{1,2,3,4,5,6}
	s13 := []int{7,8,9}
	copy(s12,s13[1:3])
	fmt.Println(s12)
	copy(s12[3:4],s13[1:3])
	fmt.Println(s12)
	//s14 := s13[:]//引用
	s14 := []int{1,2,3,4,5}
	s15 := s14
	s14[0] = 9
	fmt.Println(s15)
}
