package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store(1,2)
	v, ok := m.LoadOrStore(1,3)
	fmt.Println(v,ok)
	v, ok = m.LoadOrStore(2,3)
	fmt.Println(v,ok)
v,ok = m.Load(2)
	fmt.Println(v,ok)
m.Store(1,67)
v ,ok = m.Load(1)
fmt.Println(v,ok)
}