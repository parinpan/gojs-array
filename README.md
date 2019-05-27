# gojs-array
Who doesn't love Javascript iteration method? It keeps us away from unreadable conventional looping that sometimes hurts the eyes.

Everyone loves map, filter, reduce, and forEach Javascript's iteration methods. Find examples on main.go file:
```go
package main

import (
	"fmt"

	. "github.com/parinpan/gojs-array/src"
)

func main() {
	ages := JSIntSlice([]int{5, 6, 7, 10, 20, 30, 40})

	oldAges := ages.Filter(func(age int, index int) bool {
		return age > 10
	})

	sumOfAges := ages.Reduce(func(total int, age int, index int) int {
		return total + age
	})

	doubledAges := ages.Mapper(func(age int, index int) int {
		return age * 2
	})

	fmt.Printf("ages: %#v\n", ages)
	fmt.Printf("oldAges: %#v\n", oldAges)
	fmt.Printf("sumOfAges: %#v\n", sumOfAges)
	fmt.Printf("doubledAges: %#v\n", doubledAges)

	JSIntSlice(doubledAges).ForEach(func(age int, index int) {
		fmt.Printf("[%d] Doubled Age: %d\n", index, age)
	})
}
```
