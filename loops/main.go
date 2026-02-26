package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i, day := range days {
		fmt.Println(i, day)
	}

lco:
	println("Even number here")

	i := 4
	for i < 10 {
		if i == 2 {
			goto lco
		}
		println(i)
		i++
	}
}
