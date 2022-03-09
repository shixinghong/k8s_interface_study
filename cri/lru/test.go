package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

func main() {
	l, _ := lru.New(2)
	for i := 0; i < 256; i++ {
		l.Add(i, nil)
	}
	//if l.Len() != 128 {
	//	panic(fmt.Sprintf("bad len: %v", l.Len()))
	//}
	fmt.Println(l.Keys())
}
