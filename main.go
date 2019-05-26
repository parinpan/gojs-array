package main

import (
	"fmt"

	. "github.com/parinpan/jsarray/src"
)

func main() {
	ages := []int{5, 6, 7, 10, 20, 30, 40}

	oldAges := JSIntSlice(ages).Filter(func(age int, index int) bool {
		return age > 10
	})

	sumOfAges := JSIntSlice(ages).Reduce(func(total int, age int, index int) int {
		return total + age
	})

	doubleAges := JSIntSlice(ages).Mapper(func(age int, index int) int {
		return age * 2
	})

	fmt.Printf("ages: %#v\n", ages)
	fmt.Printf("oldAges: %#v\n", oldAges)
	fmt.Printf("sumOfAges: %#v\n", sumOfAges)
	fmt.Printf("doubleAges: %#v\n", doubleAges)

	JSIntSlice(ages).ForEach(func(age int, index int) {
		fmt.Printf("[%d] Age: %d\n", index, age)
	})
}