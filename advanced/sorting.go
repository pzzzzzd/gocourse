package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(peopel []Person) {
	ps := &personSorter{
		people: peopel,
		by:     by,
	}
	sort.Sort(ps)
}

func main() {
	people := []Person{
		{"Alice", 30},
		{"Egor", 18},
		{"John", 45},
	}

	age := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}
	By(age).Sort(people)
	fmt.Println("Sort by age =>:", people)
	By(ageDesc).Sort(people)
	fmt.Println("Sort by age <=:", people)
	By(name).Sort(people)
	fmt.Println("Sort by name:", people)
	By(lenName).Sort(people)
	fmt.Println("Sort by lenght of name:", people)

	// ======================

	strSlice := []string{"apple", "banana", "watermelon"}
	sort.Slice(strSlice, func(i, j int) bool {
		return strSlice[i][len(strSlice[i])-1] < strSlice[j][len(strSlice[j])-1]
	})
	fmt.Println("Sorted by last characret:", strSlice)

}

// // === Age ===
// type ByAge []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// // === Name ===
// type ByName []Person

// func (a ByName) Len() int {
// 	return len(a)
// }

// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }

// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func main() {
// 	people := []Person{
// 		{"Alice", 30},
// 		{"Egor", 18},
// 		{"John", 45},
// 	}
// 	sort.Sort(ByAge(people))
// 	fmt.Println("Sort by age:", people)

// }

// func main() {
// 	numbers := []int{5, 3, 4, 2, 1}
// 	sort.Ints(numbers)
// 	fmt.Println("Sorted numbers:", numbers)

// 	strSlice := []string{"John", "Egor", "Matvey", "Artem"}
// 	sort.Strings(strSlice)
// 	fmt.Println("Sorted strings:", strSlice)

// }
