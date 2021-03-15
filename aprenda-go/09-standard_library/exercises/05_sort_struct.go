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

	sort.Slice(users, func(i, y int) bool {
		return users[i].Age < users[y].Age
	})

	fmt.Println("================ Ordenador por Age")

	for _, v := range users {
		fmt.Println(v)
	}

	sort.Slice(users, func(i, y int) bool {
		return users[i].Last < users[y].Last
	})

	fmt.Println("================ Ordenador por Sobrenome")

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
