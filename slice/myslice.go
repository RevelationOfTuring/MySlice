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

const TAG = 8

type Myslice struct {
	Data unsafe.Pointer
	len  int
	cap  int
}

//创建线性表
func (s *Myslice) Create(cap int, data ...int) {
	if len(data) == 0 {
		return
	}
	s.len = len(data)
	s.cap = cap
	s.Data = C.malloc(C.ulong(cap) * TAG)

	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		*(*int)(unsafe.Pointer(p)) = data[i]
		p += TAG
	}
}

//打印线性表中的数据
func (s *Myslice) Print() {
	if s == nil {
		return
	}
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		fmt.Printf("%d ", *(*int)(unsafe.Pointer(p)))
		p += TAG
	}
	fmt.Println()
}

//获取线性表中的某一索引的元素值
func (s *Myslice) GetData(index int) (int, error) {
	if index < 0 || index >= s.len {
		err := errors.New("Error:Index out of range")
		return 0, err
	}
	if s == nil {
		err := errors.New("Myslice is empty")
		return 0, err
	}

	p := uintptr(s.Data)
	p += uintptr(index * TAG)
	return *(*int)(unsafe.Pointer(p)), nil
}

//修改元素的值
func (s *Myslice) UpdateData(index, value int) (error) {
	if s == nil {
		return errors.New("Myslice is empty")
	}
	if index < 0 || index >= s.len {
		return errors.New("Error:Index out of range")
	}
	p := uintptr(s.Data) + uintptr(index*TAG)
	*(*int)(unsafe.Pointer(p)) = value
	return nil
}

//删除元素
func (s *Myslice) DeleteData(index int) (error) {
	if s == nil {
		return errors.New("Myslice is empty")
	}
	if index < 0 || index >= s.len {
		return errors.New("Error:Index out of range")
	}
	p := uintptr(s.Data) + uintptr(index*TAG)
	for i := index; i < s.len-1; i++ {
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(p + TAG))
		p += TAG
	}
	s.len--
	return nil
}

//追加单个元素
func (s *Myslice) AppendData(value int) {
	if s == nil {
		return
	}

	var p unsafe.Pointer
	if s.len == s.cap {
		p = C.realloc(s.Data, C.ulong(s.cap*2))
		s.cap *= 2
	} else {
		p = s.Data
	}
	pInt := uintptr(p) + uintptr(s.len*TAG)

	*(*int)(unsafe.Pointer(pInt)) = value
	s.len++
	return
}
//追加多个元素
func (s *Myslice) AppendDatas(values ...int) {
	if s == nil {
		return
	}
	nValue := len(values)
	if nValue==0{
		return
	}
	var p unsafe.Pointer
	//根据添加的数据长度，来扩容
	if s.len+nValue <= s.cap {
		p = s.Data
	} else {
		for s.len+nValue > s.cap {
			p = C.realloc(s.Data, C.ulong(s.cap)*2)
			s.Data = p
			s.cap *= 2
		}
	}
	pInt:=uintptr(p)+uintptr(s.len*TAG)
	for i:=0;i<nValue;i++{
		*(*int)(unsafe.Pointer(pInt))=values[i]
		pInt+=TAG
	}
	s.len+=nValue
	return
}
