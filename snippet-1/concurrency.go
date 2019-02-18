package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//并发
//通过通信来共享内存，而不是共享内存来通信

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//c:=make(chan bool,10)
	//for i:=0;i<10;i++{
	//	go Go(c,i)
	//}
	//
	//for i:=0;i<10;i++{
	//	<-c
	//}
	//wg:=sync.WaitGroup{}
	//wg.Add(10)
	//for i := 0;i<10;i++{
	//	go Go1(&wg,i)
	//}
	//
	//wg.Wait()

	//c1,c2 :=make(chan int),make(chan string)
	//o := make(chan bool,2)//通信channel
	//go func() {
	//	for { //无限循环,不使用，则指挥接受一次，这里for 使之不断的执行信息接收和发送
	//		select{
	//		case v,ok := <-c1:
	//			if !ok{
	//				fmt.Println("c1")
	//				o<- true
	//				break
	//			}
	//			fmt.Println("c1",v)
	//		case v,ok :=<-c2:
	//			if !ok{
	//				fmt.Println("c2")
	//				break
	//			}
	//			fmt.Println("c2",v)
	//		}
	//	}
	//}()
	//
	//c1<- 1
	//c2<- "Hi"
	//c1<- 3
	//c2<- "Hello"
	//
	//close(c1)
	//close(c2)
	//
	//for i:=0;i<2;i++{
	//	<-o
	//}

	//ch := make(chan int)
	//go func(){
	//	for v :=range ch{
	//		fmt.Println(v)
	//	}
	//}()
	//
	//// <-ch ch作为发送者
	//// ch<- xx ch作为接收者
	//for {
	//	select{//select
	//	case ch<-0:
	//	case ch<-1:
	//	}
	//}

	ct := make(chan bool)
	select {
	case v := <-ct:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Time Out!")
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true

}

func Go1(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}
