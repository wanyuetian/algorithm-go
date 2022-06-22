package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	// Sort a slice of ints, float64s or strings
	// sort.Ints
	// sort.Float64s
	// sort.Strings
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3 4]

	// Sort with custom comparator
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]

	// Sort custom data structures
	persons := []Person{
		{"Eve0", 2},
		{"Alice", 23},
		{"Eve1", 2},
		{"Eve2", 2},
		{"Eve3", 2},
		{"Eve4", 2},
		{"Eve5", 2},
		{"Eve6", 2},
		{"Eve7", 2},
		{"Eve8", 2},
		{"Eve9", 2},
		{"Eve10", 2},
		{"Eve11", 2},
		{"Eve12", 2},
		{"Eve13", 2},
		{"Eve14", 2},
		{"Eve15", 2},
		{"Eve16", 2},
		{"Eve17", 2},
		{"Eve18", 2},
		{"Eve19", 2},
		{"Eve20", 2},
		{"Eve21", 2},
		{"Eve22", 2},
		{"Eve23", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(persons)) // 不稳定的排序
	fmt.Println(persons)      // [{Eve 2} {Alice 23} {Bob 25}]

	// Sort a map by key or value
	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
