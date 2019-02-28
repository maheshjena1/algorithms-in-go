package main

import "log"

// list all the possible armstrong numbers
// armstrong numbers definition: An Armstrong Number is a number of three digits, such that the sum of the cubes of it's digits is equal to the number itself. It's a specific case of Narcissistic Numbers.

// For example 153 is an Armstrong Number; because 1³ + 5³ + 3³ = 1 + 125 + 27 = 153.

func main() {
	listAllArmNum()
}

func listAllArmNum() {
	for i := 100; i <= 999; i++ {
		num := 0
		for n := i; n > 0; n = n / 10 {
			digit := n % 10
			num += (digit * digit * digit)
		}
		if num == i {
			log.Println("Armstrong number: ", i)
		}
	}
}
