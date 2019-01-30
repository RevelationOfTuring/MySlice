package main

import (
	"dataStructure/MySlice/slice"
	"fmt"
)

func main() {
	var s slice.Myslice
	//s.Create(5,99,88,77,66,55)
	//fmt.Println(s)
	//s.Print()
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
	//s.Print()
	//if err:=s.DeleteData(4);err!=nil {
	//	panic(err)
	//}
	//s.Print()
	//fmt.Print(s)
	//s.Print()
	//s.AppendData(999)
	//s.Print()
	//fmt.Println(s)
	//s.Create(1,99)
	//s.AppendDatas(2,3)
	//s.Print()
	//fmt.Println(s)
	s.Create(4,1,2,3,4)
	s.Print()
	if err:=s.InsertData(0,99);err!=nil {
		panic(err)
	}
	s.Print()
	fmt.Println(s)

}
