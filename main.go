package main

import (
	"fmt"
	"slices"
	"strconv"
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

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	} else {
		x, y := 0, 0
		for _, j := range walk {
			switch j {
			case 'n':
				y++
			case 's':
				y--
			case 'e':
				x--
			case 'w':
				x++
			}
		}
		return x == 0 && y == 0
	}
}

func DirReduc(arr []string) []string {
	var short []string
	flag := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == "NORTH" && arr[i+1] == "SOUTH" {
			arr[i] = "o"
			arr[i+1] = "o"
		} else if arr[i] == "SOUTH" && arr[i+1] == "NORTH" {
			arr[i] = "o"
			arr[i+1] = "o"
		} else if arr[i] == "EAST" && arr[i+1] == "WEST" {
			arr[i] = "o"
			arr[i+1] = "o"
		} else if arr[i] == "WEST" && arr[i+1] == "EAST" {
			arr[i] = "o"
			arr[i+1] = "o"
		}
	}

	for _, j := range arr {
		if j == "o" {
			flag = false
		} else {
			short = append(short, j)
		}
	}

	if flag == false {
		return DirReduc(short)
	}

	if len(short) == 0 {
		short = []string{}
	}
	return short
}

func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

func neihbours(m, n int) bool {
	if m == n+1 || m == n-1 {
		return true
	}
	return false
}

func RangeExtraction(list []int) string {
	res := ""
	start, stop := 0, 0
	for i := 0; i < len(list); i++ {
		start = list[i]
		counter := 0
		for j := i; j < len(list)-1; j++ {
			if neihbours(list[j], list[j+1]) {
				stop = list[j+1]
				counter++
			} else {
				stop = list[j]
				break
			}
		}
		i = i + counter

		if start != stop && counter >= 2 {
			res += fmt.Sprint(start) + "-" + fmt.Sprint(stop) + ","
		} else if start == stop {
			res += fmt.Sprint(start) + ","
		} else if start != stop && start < stop {
			res += fmt.Sprint(start) + "," + fmt.Sprint(stop) + ","
		} else if start == list[len(list)-1] && start > stop {
			res += fmt.Sprint(start) + ","
		}

	}
	return res[:len(res)-1]
}

func PrinterError(s string) string {
	counter := 0
	runes := []rune(s)
	for _, j := range runes {
		if j < 97 || j > 109 {
			fmt.Println("err", j)
			counter++
		}

	}
	return fmt.Sprint(counter) + "/" + fmt.Sprint(len(runes))
}

func contains(n int, list []int) bool {
	for _, m := range list {
		if m == n {
			return true
		}
	}
	return false
}

func SumOfIntervalsBrute(intervals [][2]int) int {
	// Bruteforce edition for mad mans only
	// 100% accurate, but consumes tons of memory
	unsorted := []int{}
	sorted := []int{}
	for _, j := range intervals {
		for i := j[0]; i < j[1]; i++ {
			unsorted = append(unsorted, i)
		}
	}

	for _, j := range unsorted {
		if !contains(j, sorted) {
			sorted = append(sorted, j)
		}
	}

	return len(sorted)
}

func GetCount(str string) (count int) {
	// Vowel Count
	for _, j := range str {
		if j == 'a' || j == 'e' || j == 'i' || j == 'o' || j == 'u' {
			count++
		}
	}
	return count
}

func PartsSums(ls []uint64) []uint64 {
	// Sums of Parts
	var sum uint64 = 0
	new_ls := []uint64{}

	for _, j := range ls {
		sum += j
	}

	new_ls = append(new_ls, sum)

	for _, j := range ls {
		new_ls = append(new_ls, sum-j)
		sum -= j
	}

	return new_ls
}

func DeleteDigit(n int) int {
	digits := []string{}
	ans := []string{}

	for _, j := range fmt.Sprint(n) {
		digits = append(digits, string(j))
	}

	for i := range digits {
		s := ""
		for j := range digits {
			if i != j {
				s += digits[j]
			}
		}
		ans = append(ans, s)
	}

	max := 0
	for _, j := range ans {
		num, _ := strconv.Atoi(j)
		if num > max {
			max = num
		}
	}

	fmt.Println(max)
	return 0
}

func ArrowArea(a, b int) float64 {
	return (float64(a) * float64(b) / 2.0) / 2.0
}

func Fib(n int) int {
	fib := []int{0, 1}

	for i := 0; i < n; i++ {
		fib = append(fib, fib[i]+fib[i+1])
	}

	return fib[len(fib)-2]
}

func SequenceSum(start, end, step int) int {
	sum := 0
	for i := start; i <= end; i += step {
		sum += i
	}
	return sum
}

func WordsToMarks(s string) int {
	runes := []rune(s)
	sum := 0
	for _, j := range runes {
		sum += int(j) - 96
	}
	return sum
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
		fmt.Println(FirstNonRepeating("sTreSS"))
		fmt.Println(IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}))
		fmt.Println(DirReduc([]string{"NORTH", "SOUTH", "EAST", "WEST"}))
		fmt.Println(IsLeapYear(2100))
		fmt.Println(RangeExtraction([]int{40, 44, 48, 51, 52, 54, 55, 58, 67, 73}))
		fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
		fmt.Println("result: ", SumOfIntervalsBrute([][2]int{{1, 4}, {7, 10}, {3, 5}}))
		fmt.Println(PartsSums([]uint64{1, 2, 3, 4, 5, 6}))
		fmt.Println(DeleteDigit(1001))
		fmt.Println(ArrowArea(25, 25), " test: 156.25")
		fmt.Println(Fib(3))
		fmt.Println(SequenceSum(1, 5, 3))
	*/
	fmt.Println(WordsToMarks("attitude"), 100)
}
