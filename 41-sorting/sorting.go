package main

import (
	"fmt"
	"sort"
)

func main()  {
	strs := []string{"a", "b", "c"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{1, 2, 3}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}