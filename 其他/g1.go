package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	i int
	j int
}
type IntList struct {
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int{
	if list == nil{
		return 0
	}
	return list.Value + list.Tail.Sum()
}
func myFunction(ms *MyStruct){
	ptr := unsafe.Pointer(ms)
	for i:= 0;i <2 ;i++{
		c := (*int)(unsafe.Pointer(uintptr(ptr)+uintptr(8*i)))
		*c += i+1
		fmt.Printf("[%p] %d \n",c,*c)
	}
}



var any interface{}
any = true
any = 12.34
any = "hello"
any =map[string]int{"one":1}
any = new(bytes.Buffer)

if f,ok := w.(*os.File);ok{
	///
}


if w,ok func main(){
	a := &MyStruct{i:40,j:6}
	myFunction(a)
	fmt.Printf("[%p]%v\n",a,a)
}


