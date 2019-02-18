package main

import "fmt"

var c chan string

func PingPang()  {
	i := 0
	for{
		fmt.Println(<-c)
		c<- fmt.Sprintf("From PingPang:Hi,#%d.\n",i)
		i++
	}
}

func main()  {
	c = make(chan string)
	go PingPang()
	for i:=0;i<10;i++{
		c<- fmt.Sprintf("From main :Hello,#%d.\n",i)
		fmt.Println(<-c)
	}
}
