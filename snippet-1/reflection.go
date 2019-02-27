package main

import (
	"fmt"
	"reflect"
)

type User struct {
 	Id int
 	Name string
 	Age int
 }

type Manager struct {
	User  //匿名字段会当作名称来处理
	title string
}

func (u User) Hello(){
	fmt.Println("Hello wold.")
}

func (u User) Hello1(name string){
	fmt.Println("Hello",name,", my name is",u.Name)
}

func main(){
	u:=User{1,"OK",12}
	Info(u)//值拷贝
	Info(&u)

	//
	m:=Manager{User:User{1,"OK",12},title:"123"}
	t:=reflect.TypeOf(m)

	//获取含有匿名struct/字段的项
	for i:=0;i<t.NumField();i++{
		fmt.Printf("%#v\n",t.Field(i))
	}
	for j:=0;j<t.NumField();j++{
		fmt.Printf("index %#v\n",t.FieldByIndex([]int{0,0}))//按索引来取字段信息
	}
	fmt.Println("_______________________________________")
	fmt.Println(t.FieldByName("Age"))//按照字段名来获取信息 字段名 类型 ？ 索引 ？ ？

	x:=123
	v:=reflect.ValueOf(&x)
	fmt.Println(v)
	v.Elem().SetInt(999)
	fmt.Println(x)

	fmt.Println("________________________________________")
	Set(&u)
	fmt.Println(u)

	u.Hello1("Joe")
	v1  := reflect.ValueOf(u)
	mv := v1.MethodByName("Hello1")

	args:=[]reflect.Value{reflect.ValueOf("Joe")}//组装参数，也可以自己填写,但是类型必须是value
	//args1 := []string{"Op"} //类型不行
	mv.Call(args)
	//mv.Call(args1)//类型不对
}

func Info(o interface{}){
	t:=reflect.TypeOf(o)

	fmt.Println("Type:",t.Name())

	if k:=t.Kind();k!=reflect.Struct{//参数判断 判断传入对象是否是 struct
		fmt.Println("BAD Arg!\n")
		return
	}

	v:=reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i :=0 ; i<t.NumField();i++{
		f:=t.Field(i)
		val:=v.Field(i).Interface()
		fmt.Printf("%6s : %v = %v\n",f.Name,f.Type,val)
	}

	for i:=0;i<t.NumMethod();i++{
		m:=t.Method(i)
		fmt.Printf("%6s: %v\n",m.Name,m.Type)//方法名称 以及
	}
}

func Set(o interface{}){
	v := reflect.ValueOf(o)

	//判断是否是指针类型 / 判断是否能修改
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet(){//注意 &&
		fmt.Println("Bad Argument!")
		return
	}else{
		v = v.Elem()
	}

	f:=v.FieldByName("Name")
	if !f.IsValid(){
		fmt.Println("Cant find field.")
		return
	}

	if f.Kind()==reflect.String{
		f.SetString("PIG")
	}


}