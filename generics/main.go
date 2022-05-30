package main

import (
	"fmt"
	"reflect"
)

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | UNum
}

type UNum interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func ConvertNum[N Num, To Num](nem N) To {
	return To(nem)
}

func SliceContains[S comparable](slice []S, contain S) bool {
	for _, v := range slice {
		if v == contain {
			return true
		}
	}

	return false
}

func main() {
	num := float32(123456)
	convert := ConvertNum[float32, uint64](num)
	fmt.Println("value before", convert)
	fmt.Println("kind before", reflect.ValueOf(num).Kind())
	fmt.Println("is convert float32 to int64 success", reflect.ValueOf(convert).Kind() == reflect.Int64)
	fmt.Println("kind after", reflect.ValueOf(convert).Kind())
	fmt.Println("value after", convert)
	// sliceAny := []interface{}{"foo", "bar"}
	sliceString := []string{"foo", "bar"}
	sliceNumber := []int{1, 2, 3}
	fmt.Println("contains foo?", SliceContains(sliceString, "foo"))
	fmt.Println("contains 1?", SliceContains(sliceNumber, 1))
}
