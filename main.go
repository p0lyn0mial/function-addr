package main

import "fmt"

func main() {
	a := &T{"a"}
	b := &T{"b"}
	fmt.Println(a.String())
	fmt.Println(b.String())
	fmt.Printf("%p\n", a.String)
	fmt.Printf("%p\n", b.String)
        if fmt.Sprintf("%p", a.String) == fmt.Sprintf("%p", b.String) {
         panic("a and b point to the same address!")
        }
}

type T struct {
	S string
}

func (t *T) String() string {
	return t.S
}
