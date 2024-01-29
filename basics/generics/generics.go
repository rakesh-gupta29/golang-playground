package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Num interface {
	int
}

// add int or floats
func Add[T Num](a T, b T) T {
	return a + b
}

func AddAny[T constraints.Ordered](a T, b T) T {
	return a + b
}

func mapArray(values []int, mapFunc func(int) int) []int {
	var newValues []int

	for _, elem := range values {
		newValue := mapFunc(elem)
		newValues = append(newValues, newValue)
	}
	return newValues
}

type CustomMap[Key comparable, Value string | int] map[Key]Value

// this will be allowing all the values which can be compared directly.

func UseGenerics() {
	sample := Add(1, 2)

	sample2 := AddAny(10, 30)

	fmt.Println(mapArray([]int{1, 200, 3}, func(elem int) int { return elem * 10 }))
	fmt.Println(sample, sample2)

	maps := make(CustomMap[string, string])
	maps["hello"] = "hi"
}
