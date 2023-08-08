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

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type ByLast []user

func (p ByLast) Len() int           { return len(p) }
func (p ByLast) Less(i, j int) bool { return p[i].Last < p[j].Last }
func (p ByLast) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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

	// fmt.Println(users)

	fmt.Println("------------------------------")
	sort.Sort(ByAge(users))

	fmt.Println("sorted by age")
	for k, v := range users {
		fmt.Println("\nuser #", k, "is", v.First, v.Last, ",age -", v.Age, " says:")
		sort.Strings(v.Sayings)
		for k1, v1 := range v.Sayings {
			fmt.Println(k1, "\t", v1)
		}
	}
	fmt.Println("------------------------------")
	sort.Sort(ByLast(users))

	fmt.Println("sorted by last")
	for k, v := range users {
		fmt.Println("\nuser #", k, "is", v.First, v.Last, ",age -", v.Age, " says:")
		for k1, v1 := range v.Sayings {
			fmt.Println(k1, "\t", v1)
		}
	}

	// 	fmt.Println("------------------------------")
	// 	fmt.Println("sort the sayings")

	//	for k, v := range users {
	//		fmt.Println("\nuser #", k, "is", v.First, v.Last, ",age -", v.Age, " says:")
	//		sort.Strings(v.Sayings)
	//		for k1, v1 := range v.Sayings {
	//			fmt.Println(k1, "\t", v1)
	//		}
	//	}
}

/*
Hands-on exercise #5 Starting with this code, sort the []user by
● age
● last
Also sort each []string “Sayings” for each user
● print everything in a way that is pleasant
*/
