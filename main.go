package main

import (
	"dataStructure/MySlice/slice"
	"fmt"
)

func main() {
	var s slice.Myslice
	s.Create(10,99,88,77,66,55)
	fmt.Println(s)
	s.Print()
	//ret,err:=s.GetData(6)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(ret)
	//if err:=s.UpdateData(5,100);err!=nil {
	//	panic(err)
	//}
	//s.Print()
	s.Print()
	if err:=s.DeleteData(4);err!=nil {
		panic(err)
	}
	s.Print()
	fmt.Print(s)

}
