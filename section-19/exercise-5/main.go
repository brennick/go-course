package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user
type ByLast []user

func (u ByAge) Len() int           { return len(u) }
func (u ByAge) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u ByAge) Less(i, j int) bool { return u[i].Age < u[j].Age }

func (u ByLast) Len() int           { return len(u) }
func (u ByLast) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u ByLast) Less(i, j int) bool { return u[i].Last < u[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, u := range users {
		sort.Strings(u.Sayings)
	}

	fmt.Println(users)
	fmt.Println()

	// your code goes here
	sort.Sort(ByAge(users))
	fmt.Println(users)
	fmt.Println()

	sort.Sort(ByLast(users))
	fmt.Println(users)
}
