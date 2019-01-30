package slice

import (
	"unsafe"
	"fmt"
	"errors"

)

/*
#include <stdlib.h>
 */
 import "C"

 const TAG=8
type Myslice struct {
	Data unsafe.Pointer
	len int
	cap int
}
//创建线性表
func (s*Myslice)Create(cap int , data ...int){
	if len(data)==0{
		return
	}
	s.len=len(data)
	s.cap=cap
	s.Data=C.malloc(C.ulong(cap)*TAG)

	p:=uintptr(s.Data)
	for i:=0;i<s.len;i++ {
		*(*int)(unsafe.Pointer(p))=data[i]
		p+=TAG
	}
}
//打印线性表中的数据
func (s*Myslice)Print(){
	if s==nil {
		return
	}
	p:=uintptr(s.Data)
	for i:=0;i<s.len;i++{
		fmt.Printf("%d ",*(*int)(unsafe.Pointer(p)))
		p+=TAG
	}
	fmt.Println()
}
//获取线性表中的某一索引的元素值
func (s*Myslice)GetData(index int)(int,error){
	if index<0||index>=s.len{
		err:=errors.New("Error:Index out of range")
		return 0,err
	}
	if s==nil {
		err:=errors.New("Myslice is empty")
		return 0,err
	}

	p:=uintptr(s.Data)
	p+=uintptr(index*TAG)
	return *(*int)(unsafe.Pointer(p)),nil
}
//修改元素的值
func (s*Myslice)UpdateData(index,value int)(error){
	if s==nil {
		return errors.New("Myslice is empty")
	}
	if index<0||index>=s.len{
		return errors.New("Error:Index out of range")
	}
	p:=uintptr(s.Data)+uintptr(index*TAG)
	*(*int)(unsafe.Pointer(p))=value
	return nil
}