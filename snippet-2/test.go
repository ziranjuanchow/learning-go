package main

import "fmt"
var temp map[int]int
func main(){
	temp = make(map[int]int,1)
	temp[1] = 1
	if temp1,ok:=temp[2];ok{
		fmt.Println(temp1)
	}else{
		fmt.Println(ok)
		fmt.Println("Error")
	}
}


