package main

import "fmt"

func main () {
//	a := 1
//	for i:=0;i<=3;i++{
//		a++
//		fmt.Println(a)
//	}
//
//	switch  {
//	case a == 1:
//		fmt.Println("a=1")
//	case a>1:
//		fmt.Println("a>1")
//		fallthrough
//	default:
//		fmt.Println("None")
//	}
//
////LABEL1:
//	for{
//		for i:=0;i<10;i++{
//			fmt.Println(i)
//			if i>3{
//				goto LABEL2
//			}
//		}
//	}
//LABEL2:
//	fmt.Println("LABEL2")
//LABEL3:
//	for i:=0;i<10;i++{
//		for{
//			continue LABEL3
//			fmt.Println("for")
//		}
//	}

LABEL1:
	for i := 0;i<10;i++	{
		for{
			fmt.Println(i)
			continue LABEL1
		}
	}
LABEL2:
	for i := 0;i<10;i++	{
		for{
			fmt.Println(i)
			goto LABEL2
		}
	}
fmt.Println("Over")
}