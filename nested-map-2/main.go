package main

import (
	"fmt"
)

type N[T any] struct {
	val   T
	ok    bool
	child map[any]*N[T]
}

func (n *N[T]) C(key any) *N[T] {
	// fmt.Println(reflect.TypeOf(n.child[key]).Elem())
	if n != nil {
		return n.child[key]
	}

	return nil
}

func (n *N[T]) Get() (T, bool) {
	return n.val, n.ok
}

func main() {
	myMap := N[string]{
		child: map[any]*N[string]{
			"foo": {val: "1", ok: true},
		},
	}

	fmt.Println(myMap.C("child1"))
	fmt.Println(myMap.C("child2"))
	fmt.Println(myMap.C("child2"))

}
