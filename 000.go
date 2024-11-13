package main

import "fmt"

func main() {

	fmt.Println("Start Loop...")
	b := 1
	for b <= 5 {

		if b == 3 {

			break

		}

		fmt.Println(b)
		b += 1

	}

	fmt.Println("End Loop...")

	for i := 1; i <= 5; i++ {

		if i == 3 {

			continue
		}

		fmt.Println(i)
	}
}
