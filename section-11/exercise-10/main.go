package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["test"] = []string{"testing", "123"}

	for k, a := range m {
		for i, v := range a {
			fmt.Println(k, i, v)
		}
	}

	fmt.Println()
	fmt.Println("DELETING")
	fmt.Println()

	delete(m, "bond_james")

	for k, a := range m {
		for i, v := range a {
			fmt.Println(k, i, v)
		}
	}
}
