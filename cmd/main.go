package main

import "fmt"

func main() {
	a := Test{
		ID:   "123",
		name: "aaa",
	}

	stores := store{
		obj: map[string]*Test{},
	}

	stores.obj[a.ID] = &a

	fmt.Println(stores.obj["123"])

	changeName(&a)
	b := a
	b.name = "ccc"
	fmt.Println(a)
	fmt.Println(stores.obj["123"])
}

type store struct {
	obj map[string]*Test
}

type Test struct {
	ID   string
	name string
}

func changeName(target *Test) {
	target.name = "gg"
}
