package main

import (
	"fmt"
)

type NestedMap map[any]any

func (nMap NestedMap) C(key any) NestedMap {
	if v, ok := nMap[key].(NestedMap); ok {
		return v
	} else {
		return NestedMap{0: nMap[key]}
	}
}

func (nMap NestedMap) Get() any {
	// fmt.Println("len", len(nMap))
	// fmt.Println("elem", reflect.TypeOf(nMap).Elem())
	// fmt.Println("nMap", nMap)
	if len(nMap) == 1 {
		// fmt.Println("single value")
		for _, v := range nMap {
			return v
		}
	} else if len(nMap) > 1 {
		// fmt.Println("child value")
		return nMap
	}

	return nil
}

func main() {
	myMap := NestedMap{
		"child1": NestedMap{
			"foo1": NestedMap{
				1: 2,
				2: 3,
			},
			"foo2": "bar2",
		},
		"child2": NestedMap{1: 2},
	}

	fmt.Println(myMap.C("child1").Get())
	fmt.Println(myMap.C("child1").C("foo1").Get())
	fmt.Println(myMap.C("child1").C("foo1").C(1).Get())
	fmt.Println(myMap.C("child1").C("foo1").C(2).Get())
	fmt.Println(myMap.C("child1").C("foo2").Get())
	fmt.Println(myMap.C("child2").Get())
}
