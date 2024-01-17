package main

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {
	/*
		Complete the method/function so that it converts dash/underscore delimited words into camel casing.
		The first word within the output should be capitalized only if the original word was capitalized
		(known as Upper Camel Case, also often referred to as Pascal case).
		The next words should be always capitalized.
	*/
	runes := []rune(s)
	var newString []rune
	for i := 0; i < len(runes); i++ {
		if runes[i] == '_' || runes[i] == '-' {
			runes[i+1] = unicode.ToUpper(runes[i+1])
		} else {
			newString = append(newString, runes[i])
		}

	}

	return string(newString)
}

func Multiple3And5(number int) int {
	/*
		If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
		The sum of these multiples is 23.
		Finish the solution so that it returns the sum of all the multiples of 3 or 5 below the number passed in.
		Note: If the number is a multiple of both 3 and 5, only count it once.
	*/
	acc := 0
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			acc += i
		}
	}
	return acc
}

func IsPrime(n int) bool {
	/*
		Define a function that takes an integer argument and returns a logical value true or false
		depending on if the integer is a prime. Per Wikipedia, a prime number ( or a prime )
		is a natural number greater than 1 that has no positive divisors other than 1 and itself.
	*/

	if n <= 0 || n == 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func StringEndsWith(str, ending string) bool {
	/*
		Complete the solution so that it returns true if the first argument(string)
		passed in ends with the 2nd argument (also a string).
	*/
	return strings.HasSuffix(str, ending)
}

func MinMax(arr []int) [2]int {
	/*
		Write a function that returns both the minimum and maximum number of the given list/array.
	*/
	return [2]int{slices.Min(arr), slices.Max(arr)}
}

func IsTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func FirstNonRepeating(str string) string {
	/*
		Write a function named first_non_repeating_letter that takes a string input
		and returns the first character that is not repeated anywhere in the string.
	*/
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		letter := runes[i]
		counter := 0
		for j := 0; j < len(runes); j++ {
			if unicode.ToLower(letter) == unicode.ToLower(runes[j]) {
				counter++
			}
		}
		if counter == 1 {
			return string(letter)
		}
	}
	return ""
}

func main() {
	fmt.Println("Codewars")
	/*
		fmt.Println(ToCamelCase("to_camel-case"))
		fmt.Println(Multiple3And5(10))
		fmt.Println(IsPrime(8))
		fmt.Println(StringEndsWith("banana", "ana"))
		fmt.Println(MinMax([]int{2334454, 5}))
		fmt.Println(IsTriangle(1, 2, 2))
	*/
	fmt.Println(FirstNonRepeating("sTreSS"))
}
