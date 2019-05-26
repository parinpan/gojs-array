package slice

type IntSlice struct {
	slice []int
}

type forEachClosure func(int, int)
type mapClosure func(int, int) int
type filterClosure func(int, int) bool
type reduceClosure func(int, int, int) int

// JSIntSlice aims to make new objects of IntSlice struct
func JSIntSlice(slice []int) *IntSlice {
	return &IntSlice{slice: slice}
}

// use mapper as name instead of map because it's a reserved keyword
func (s IntSlice) Mapper(closure mapClosure) []int {
	slice := make([]int, len(s.slice))

	for index, value := range s.slice {
		slice[index] = closure(value, index)
	}

	return slice
}

func (s IntSlice) Filter(closure filterClosure) []int {
	var slice []int

	for index, value := range s.slice {
		if closure(value, index) {
			slice = append(slice, value)
		}
	}

	return slice
}

func (s IntSlice) Reduce(closure reduceClosure) int {
	var total = 0

	for index, value := range s.slice {
		total = closure(total, value, index)
	}

	return total
}

func (s IntSlice) ForEach(closure forEachClosure) {
	for index, value := range s.slice {
		closure(value, index)
	}
}
