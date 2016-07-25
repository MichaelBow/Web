/*Use https://godoc.org/sort to sort the following:

(1)
type people []string
studyGroup := people{"Zeno", "John", "Al", "Jenny"}

(2)
s := []string{"Zeno", "John", "Al", "Jenny"}

(3)
n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

Also sort the above in reverse order */

package main

import (
	"fmt"
	"sort"
)

type people []string

//type name interface {
//	area() float64
//}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Jenny"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println("B4")
	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)

	sort.Strings(s)
	sort.Strings(studyGroup)
	sort.Ints(n)

	fmt.Println("after")
	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println("after2")
	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)

}
