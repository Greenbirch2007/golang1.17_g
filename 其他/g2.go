package main

import "fmt"

func main(){
	m := map[string]int{
	}
	m["a"]=1
	m["b"]=
	m["c"]=1
	m["d"]=1
	fmt.Printf("%d \n",m["c"])
	delete(m,"d")
	for k,v := range m{
		fmt.Printf("%s,%d \n",k,v)

	}
	fmt.Printf("%+v\n",m)
}
