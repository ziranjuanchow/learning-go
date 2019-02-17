package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}
type TVConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string{
	return pc.name
}

func (pc PhoneConnecter) Connect(){
	fmt.Println("Connect",pc.name)
}
func (tv TVConnecter) Connect(){
	fmt.Println("Connect",tv.name)
}

func main(){
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)

	pc := PhoneConnecter{"PhoneConnecter"}
	var b Connecter
	b = Connecter(pc)
	b.Connect()
    pc.name = "ppppp"
    b.Connect()

    var c interface{}
    fmt.Println(c == nil)

    var d *int = nil
    c = d//指向对象为nill,但是他本身是指针
    fmt.Println(c == nil)

	//tv:=TVConnecter{"TVConnecter"}
	//var c USB
	//c = USB(tv)//不能降级
	//c.Connect()
}

//func Disconnect(usb USB){
func Disconnect(usb interface{}){
	//if pc,ok:= usb.(PhoneConnecter);ok{//判断类型 类型断言
	//	fmt.Println("Disconnected ",pc.name)
	//	return
	//}

	switch v:=usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:",v.name)
	default:
		fmt.Println("Unknow decvice. ")
	}


}