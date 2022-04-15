package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
}

type sortedByAge []user
type sortedByName []user

func (a sortedByAge) Len() int { return len(a) }
func (a sortedByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a sortedByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (n sortedByName) Len() int { return len(n) }
func (n sortedByName) Less(i, j int) bool { return n[i].First < n[j].First }
func (n sortedByName) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(sortedByAge(users))
	fmt.Println(users)
	sort.Sort(sortedByName(users))
	fmt.Println(users)

}
