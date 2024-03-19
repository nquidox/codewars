package main

import (
	"fmt"
	"math"
	"math/big"
	"slices"
	"sort"
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

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, j := range a {
		if j != b[i] {
			return false
		}
	}
	return true
}

func InAscOrder(numbers []int) bool {
	s := make([]int, len(numbers))
	copy(s, numbers)
	sort.Ints(s)
	fmt.Println(numbers, s)
	return compareSlices(numbers, s)
}

func HighAndLow(in string) string {
	nums := strings.Fields(in)
	var ints []int

	for _, j := range nums {
		s, _ := strconv.Atoi(j)
		ints = append(ints, s)
	}

	sort.Ints(ints)
	return strconv.Itoa(ints[len(ints)-1]) + " " + strconv.Itoa(ints[0])
}

func Divisors(n int) int {
	counter := 1
	for i := 1; i < n; i++ {
		if n%i == 0 {
			counter++
		}
	}
	return counter
}

func SumCubes(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

func Spacify(s string) string {
	var s2 string

	for _, j := range s {
		s2 += fmt.Sprintf("%c ", j)
	}

	return strings.TrimSuffix(s2, " ")
}

func HasUniqueChar(str string) bool {
	for i := range str {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}

func FindShort(s string) int {
	sep := strings.Split(s, " ")
	shortest := len(sep[0])

	for _, j := range sep {
		if len(j) < shortest {
			shortest = len(j)
		}
	}

	return shortest
}

func MoveZeros(arr []int) []int {
	counter := 0
	var arr2 []int

	for _, j := range arr {
		if j == 0 {
			counter++
		} else {
			arr2 = append(arr2, j)
		}
	}

	for i := 0; i < counter; i++ {
		arr2 = append(arr2, 0)
	}

	return arr2
}

func DivisibleCount(x, y, k uint64) uint64 {
	//math.Ceil and convertions to float are unecessary, but I'll keep em
	counter := uint64(math.Ceil(float64(x)/float64(k)) * float64(k))
	if counter > y {
		return 0
	}
	return ((y - counter) / k) + 1
}

func SquarePi(digits int) int {
	pi := "31415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"
	var sum int

	for i := 0; i < digits; i++ {
		sum += int(pi[i]-'0') * int(pi[i]-'0')
	}

	return int(math.Ceil(math.Sqrt(float64(sum))))
}

func Gcd(x, y uint32) uint32 {
	for x != y {
		if x > y {
			x, y = y, x
		}
		y = y - x
	}
	return x
}

func FindMissingLetter(chars []rune) rune {
	for i := 0; i < len(chars)-1; i++ {
		if chars[i+1] != chars[i]+1 {
			return chars[i] + 1
		}
	}
	return 'a'
}

func FindNextSquare(sq int64) int64 {
	sqr := math.Sqrt(float64(sq))

	if math.Mod(sqr, 1) == 0 {
		return int64(math.Pow(sqr+1, 2))
	}

	return -1
}

func Solution(str string) []string {
	// Split Strings
	var arr []string
	if len(str)%2 != 0 {
		str += "_"
	}

	for i := 0; i < len(str); i = i + 2 {
		arr = append(arr, string(str[i])+string(str[i+1]))
	}

	return arr
}

func ReverseLetters(s string) (s2 string) {
	fmt.Println(s[0], s[1])
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > 96 && s[i] < 123 || s[i] > 64 && s[i] < 91 {
			s2 += string(s[i])
		}
	}
	return
}

func FindEvenIndex(arr []int) int {
	sum, leftSum := 0, 0

	for _, j := range arr {
		sum += j
	}

	for i := range arr {
		sum -= arr[i]

		if sum == leftSum {
			return i
		}

		leftSum += arr[i]
	}
	return -1
}

func Pell(n int) *big.Int {
	// works not good on high nubmers cuz of type conversion
	// array should be big.Int type instead of uint64
	pell := []uint64{0, 1}

	for i := 2; i <= n; i++ {
		pell = append(pell, 2*pell[i-1]+pell[i-2])
	}

	return new(big.Int).SetUint64(pell[n])
}

func Invert(arr []int) []int {
	var res []int
	for _, j := range arr {
		res = append(res, 0-j)
	}
	return res
}

func AmountOfPages(summary int) int {
	if summary > 38889 {
		return 9999 + (summary-38889)/5
	} else if summary > 2889 {
		return 999 + (summary-2889)/4
	} else if summary > 189 {
		return 99 + (summary-189)/3
	} else if summary > 9 {
		return 9 + (summary-9)/2
	} else if summary <= 9 {
		return summary
	}
	return 0
}

func StringToNumber(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func NumberToString(n int) string {
	return strconv.Itoa(n)
}

func HowMuchILoveYou(i int) string {
	switch i % 6 {
	case 1:
		return "I love you"
	case 2:
		return "a little"
	case 3:
		return "a lot"
	case 4:
		return "passionately"
	case 5:
		return "madly"
	default:
		return "not at all"
	}
}

func monkeyCount(n int) []int {
	monkeys := []int{}
	for i := 1; i <= n; i++ {
		monkeys = append(monkeys, i)
	}
	return monkeys
}

func Maps(x []int) []int {
	var res []int
	for _, j := range x {
		res = append(res, j*2)
	}
	return res
}

func Grow(arr []int) int {
	res := 1
	for _, j := range arr {
		res *= j
	}
	return res
}

func PowersOfTwo(n int) []uint64 {
	var res []uint64
	for i := 0; i <= n; i++ {
		res = append(res, uint64(math.Pow(2, float64(i))))
	}
	return res
}

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	sumNeg, countPos := 0, 0

	for _, j := range numbers {
		if j < 0 {
			sumNeg += j
		} else if j > 0 {
			countPos++
		}
	}

	res = append(res, countPos)
	res = append(res, sumNeg)

	return res
}

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, j := range s {
		if j > 96 && j < 123 {
			return false
		}
	}
	return true
}

func Feast(beast string, dish string) bool {
	return beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1]
}

func MakeNegative(x int) int {
	switch {
	case x < 0:
		return x
	case x == 0:
		return 0
	}
	return -x
}

func PositiveSum(numbers []int) int {
	sum := 0
	for _, j := range numbers {
		if j > 0 {
			sum += j
		}
	}
	return sum
}

func ReverseString(word string) string {
	runes := []rune(word)
	size := len(runes)
	for i, j := 0, size-1; i < size>>1; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func Rps(p1, p2 string) string {
	r, p, s := "rock", "paper", "scissors"
	switch {
	case p1 == r && p2 == p:
		{
			return "Player 2 won!"
		}
	case p1 == r && p2 == s:
		{
			return "Player 1 won!"
		}
	case p1 == p && p2 == r:
		{
			return "Player 1 won!"
		}
	case p1 == p && p2 == s:
		{
			return "Player 2 won!"
		}
	case p1 == s && p2 == r:
		{
			return "Player 2 won!"
		}
	case p1 == s && p2 == p:
		{
			return "Player 1 won!"
		}
	}
	return "Draw!"
}

func CountBy(x, n int) (res []int) {
	for i := 1; i <= n; i++ {
		res = append(res, x*i)
	}
	return res
}

func CountSheeps(numbers []bool) int {
	counter := len(numbers)
	for _, j := range numbers {
		if j != true {
			counter--
		}
	}
	return counter
}

func Well(x []string) string {
	good := 0
	for _, j := range x {
		if j == "good" {
			good++
		}
	}

	switch {
	case good > 0 && good < 3:
		return "Publish!"
	case good > 2:
		return "I smell a series!"
	}
	return "Fail!"
}

func GetSize(w, h, d int) [2]int {
	return [2]int{2 * (w*h + h*d + w*d), w * h * d}
}

func FindMultiples(integer, limit int) (res []int) {
	for i := 1; i <= limit; i++ {
		if i%integer == 0 {
			res = append(res, i)
		}
	}
	return
}

func Between(a, b int) (res []int) {
	for i := a; i <= b; i++ {
		res = append(res, i)
	}
	return
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func Hero(bullets, dragons int) bool {
	return bullets >= dragons*2
}

func Points(games []string) (res int) {
	for _, j := range games {
		p := strings.Split(j, ":")
		switch {
		case p[0] > p[1]:
			res += 3
		case p[0] == p[1]:
			res++
		}
	}
	return
}

func CorrectTail(body string, tail rune) bool {
	return rune(body[len(body)-1]) == tail
}

func SumMix(arr []any) (sum int) {
	for _, j := range arr {
		num, t := j.(int)
		if t == true {
			sum += num
		} else {
			n, _ := strconv.Atoi(j.(string))
			sum += n
		}
	}
	return sum
}

func HowManyDalmatians(number int) string {
	switch {
	case number < 10:
		return "Hardly any"
	case number >= 10 && number <= 50:
		return "More than a handful!"
	case number > 50 && number < 101:
		return "Woah that's a lot of dogs!"
	case number == 101:
		return "101 DALMATIONS!!!"
	}
	return ""
}

func Digitize(n int) (rev []int) {
	if n > 0 {
		last := 0
		for i := 10; i < n*10; i *= 10 {
			rev = append(rev, ((n%i)-last)/(i/10))
		}
		return rev
	}
	return []int{0}
}

func MergeArrays(arr1, arr2 []int) []int {
	var r []int
	res := append(arr1, arr2...)
	sort.Ints(res)
	for i := 0; i < len(res)-1; i++ {
		if res[i] != res[i+1] {
			r = append(r, res[i])
		}
	}
	r = append(r, res[len(res)-1])
	return r
}

func GetPlanetName(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	case 9:
		return "Pluto"
	}
	return ""
}

func CloseCompare(a, b, margin float64) int {
	if math.Abs(a-b) <= margin {
		return 0
	} else {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
	}
	return 42
}

func SmallestIntegerFinder(numbers []int) (res int) {
	res = numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < res {
			res = numbers[i]
		}
	}
	return
}

func FakeBin(x string) (res string) {
	for _, j := range x {
		switch {
		case j < 53:
			res += "0"
		case j >= 53:
			res += "1"
		}
	}
	return
}

func SquareSum(numbers []int) (res int) {
	for _, j := range numbers {
		res += j * j
	}
	return
}

func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}

func DoubleInteger(i int) int {
	return i * 2
}

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

func GreetJ(name string) string {
	if name == "Johnny" {
		return fmt.Sprint("Hello, my love!")
	}
	return fmt.Sprintf("Hello, %v!", name)
}

func ShortLongShort(a, b string) string {
	if len(a) < len(b) {
		return a + b + a
	}
	return b + a + b
}

func MakeUpperCase(str string) (res string) {
	for _, j := range str {
		if j > 96 && j < 123 {
			res += fmt.Sprintf("%c", j-32)
		} else {
			res += fmt.Sprintf("%c", j)
		}
	}
	return
}

func SquareOrSquareRoot(arr []int) (res []int) {
	for _, j := range arr {
		m, n := math.Modf(math.Sqrt(float64(j)))
		if n != 0 {
			res = append(res, j*j)
		} else {
			res = append(res, int(m))
		}
	}
	return
}

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}

func multipleOfIndex(ints []int) (res []int) {
	for i := 1; i < len(ints); i++ {
		if ints[i]%i == 0 {
			res = append(res, ints[i])
		}
	}
	return
}

func Multiply(a, b int) int {
	return a * b
}

func BoolToWord(word bool) string {
	if word == true {
		return "Yes"
	}
	return "No"
}

func Opposite(value int) int {
	return value * -1
}

func ReverseSeq(n int) (res []int) {
	for i := n; i > 0; i-- {
		res = append(res, i)
	}
	return
}

func Litres(time float64) int {
	return int(time * 0.5)
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	r := sonYearsOld*2 - dadYearsOld
	if r < 0 {
		return -r
	}
	return r
}

func Xor(a, b bool) bool {
	return (a || b) && !(a && b)
}

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func AbbrevName(name string) string {
	n := strings.Fields(name)
	return strings.ToUpper(string(n[0][0]) + "." + string(n[1][0]))
}

func main() {
	fmt.Println("Codewars")
	fmt.Println(AbbrevName("David Mendieta"), "D.M")
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
		fmt.Println(WordsToMarks("attitude"), 100)
		fmt.Println(InAscOrder([]int{1, 2, 4, 7, 19}))
		fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"), "42 -9")
		fmt.Println(Divisors(12), 6)
		fmt.Println(SumCubes(3), 36)
		fmt.Println(Spacify("Hello world!"))
		fmt.Println(HasUniqueChar("Hello world!"))
		fmt.Println(FindShort("Test Example"))
		fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
		fmt.Println(DivisibleCount(6, 11, 2))
		fmt.Println(SquarePi(3))
		fmt.Println(FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'}))
		fmt.Println(FindNextSquare(121))
		fmt.Println(Solution("kekwait"))
		fmt.Println(ReverseLetters("AZ"))
		fmt.Println(FindEvenIndex([]int{-1, -2, -3, -4, -3, -2, -1}), 3)
		fmt.Println(AmountOfPages(11367), 3118)
		fmt.Println(HowMuchILoveYou(7))
		fmt.Println(monkeyCount(10))
		fmt.Println(Maps([]int{1, 2, 3}))
		fmt.Println(Grow([]int{1, 2, 3}), 6)
		fmt.Println(PowersOfTwo(4), "1,2,4,8,16")
		fmt.Println(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}), "[10, -65]")
		fmt.Println(MyString("a").IsUpperCase(), false)
		fmt.Println(Feast("great blue heron", "garlic naan"))
		fmt.Println(MakeNegative(42), -42)
		fmt.Println(PositiveSum([]int{1, -2, 3, 4, 5}), 13)
		fmt.Println(ReverseString("word"), "drow")
		fmt.Println(RepeatStr(3, "hello "))
		fmt.Println(Rps("rock", "scissors"), "p1 won")
		fmt.Println(CountBy(1, 5), 12345)
		fmt.Println(CountSheeps([]bool{true, true, true, false, true, true, true, true, true, false, true, false,
			true, false, false, true, true, true, true, true, false, false, true, true, }), 17)
		fmt.Println(Well([]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}), "I smell a series!")
		fmt.Println(GetSize(4, 2, 6), 88, 48)
		fmt.Println(FindMultiples(2, 6), "2, 4, 6")
		fmt.Println(Between(1, 4), "[1 2 3 4]")
		fmt.Println(EvenOrOdd(2), "Even")
		fmt.Println(Greet("kek"))
		fmt.Println(Hero(10, 5), true)
		fmt.Println(Points([]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}), 10)
		fmt.Println(CorrectTail("Fox", 'x'), true)
		fmt.Println(SumMix([]any{9, 3, "7", "3"}), 22)
		fmt.Println(HowManyDalmatians(80), "Woah that's a lot of dogs!")
		fmt.Println(Digitize(35231), "[1,3,2,5,3]")
		fmt.Println(MergeArrays([]int{1, 3, 3, 3, 3, 5, 7, 9, 11, 12}, []int{8, 6, 1, 2, 3, 4, 5, 10, 12}))
		fmt.Println(GetPlanetName(3), "Earth")
		fmt.Println(CloseCompare(8.1, 5.0, 3.0), 1)
		fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2, -4}), -4)
		fmt.Println(FakeBin("45385593107843568"), "01011110001100111")
		fmt.Println(SquareSum([]int{0, 3, 4, 5}), 50)
		fmt.Println(IsDivisible(12, 3, 6), true)
		fmt.Println(GetVolumeOfCuboid(1.0, 2.0, 2.0), 4.0)
		fmt.Println(DoubleInteger(2), 4)
		fmt.Println(RemoveChar("word"), "or")
		fmt.Println(GreetJ("Alfred"), "Hello, Alfred!")
		fmt.Println(ShortLongShort("aaa", "bbbb"), "aaabbbbaaa")
		fmt.Println(MakeUpperCase("hello"), "HELLO")
		fmt.Println(SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}), "[2 9 3 49 4 1]")
		fmt.Println(StringToArray("I love arrays they are my favorite"),
		"[I love arrays they are my favorite]")
		fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25}), "-6, 32, 25")
		fmt.Println(TwiceAsOld(55, 30), 5)
		fmt.Println(DNAtoRNA("GCAT"), "GCAU")
	*/
}
