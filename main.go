package main

import (
	"crypto/md5"
	"encoding/hex"
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

func OddCount(n int) int {
	if n%2 == 0 {
		return (n - 1) / 2
	}
	return n / 2
}

func QuarterOf(month int) int {
	return int(math.Ceil(float64(month) / 3))
}

func LoveFunc(flower1, flower2 int) bool {
	return (flower1+flower2)%2 != 0
}

func century(year int) int {
	return ((year - 1) / 100) + 1
}

func ReverseWords(str string) (res string) {
	s := strings.Fields(str)
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] + " "
	}
	return strings.Trim(res, " ")
}

func NthEven(n int) int {
	return (n - 1) * 2
}

func OtherAngle(a int, b int) int {
	return 180 - (a + b)
}

func countSheep(num int) (res string) {
	if num >= 1 {
		for i := 1; i <= num; i++ {
			res += fmt.Sprintf("%d sheep...", i)
		}
	} else {
		res = ""
	}
	return
}

func Past(h, m, s int) int {
	return ((h * 3600) + (m * 60) + s) * 1000
}

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func subSum(n int) int {
	num, sum := n, 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return num - sum
}

func SubtractSum(n int) string {
	num := n
	if num > 100 {
		for num > 100 {
			num = subSum(num)
		}
	} else {
		num = subSum(num)
	}

	switch num {
	case 1:
		return "kiwi"
	case 2:
		return "pear"
	case 3:
		return "kiwi"
	case 4:
		return "banana"
	case 5:
		return "melon"
	case 6:
		return "banana"
	case 7:
		return "melon"
	case 8:
		return "pineapple"
	case 9:
		return "apple"
	case 10:
		return "pineapple"
	case 11:
		return "cucumber"
	case 12:
		return "pineapple"
	case 13:
		return "cucumber"
	case 14:
		return "orange"
	case 15:
		return "grape"
	case 16:
		return "orange"
	case 17:
		return "grape"
	case 18:
		return "apple"
	case 19:
		return "grape"
	case 20:
		return "cherry"
	case 21:
		return "pear"
	case 22:
		return "cherry"
	case 23:
		return "pear"
	case 24:
		return "kiwi"
	case 25:
		return "banana"
	case 26:
		return "kiwi"
	case 27:
		return "apple"
	case 28:
		return "melon"
	case 29:
		return "banana"
	case 30:
		return "melon"
	case 31:
		return "pineapple"
	case 32:
		return "melon"
	case 33:
		return "pineapple"
	case 34:
		return "cucumber"
	case 35:
		return "orange"
	case 36:
		return "apple"
	case 37:
		return "orange"
	case 38:
		return "grape"
	case 39:
		return "orange"
	case 40:
		return "grape"
	case 41:
		return "cherry"
	case 42:
		return "pear"
	case 43:
		return "cherry"
	case 44:
		return "pear"
	case 45:
		return "apple"
	case 46:
		return "pear"
	case 47:
		return "kiwi"
	case 48:
		return "banana"
	case 49:
		return "kiwi"
	case 50:
		return "banana"
	case 51:
		return "melon"
	case 52:
		return "pineapple"
	case 53:
		return "melon"
	case 54:
		return "apple"
	case 55:
		return "cucumber"
	case 56:
		return "pineapple"
	case 57:
		return "cucumber"
	case 58:
		return "orange"
	case 59:
		return "cucumber"
	case 60:
		return "orange"
	case 61:
		return "grape"
	case 62:
		return "cherry"
	case 63:
		return "apple"
	case 64:
		return "cherry"
	case 65:
		return "pear"
	case 66:
		return "cherry"
	case 67:
		return "pear"
	case 68:
		return "kiwi"
	case 69:
		return "pear"
	case 70:
		return "kiwi"
	case 71:
		return "banana"
	case 72:
		return "apple"
	case 73:
		return "banana"
	case 74:
		return "melon"
	case 75:
		return "pineapple"
	case 76:
		return "melon"
	case 77:
		return "pineapple"
	case 78:
		return "cucumber"
	case 79:
		return "pineapple"
	case 80:
		return "cucumber"
	case 81:
		return "apple"
	case 82:
		return "grape"
	case 83:
		return "orange"
	case 84:
		return "grape"
	case 85:
		return "cherry"
	case 86:
		return "grape"
	case 87:
		return "cherry"
	case 88:
		return "pear"
	case 89:
		return "cherry"
	case 90:
		return "apple"
	case 91:
		return "kiwi"
	case 92:
		return "banana"
	case 93:
		return "kiwi"
	case 94:
		return "banana"
	case 95:
		return "melon"
	case 96:
		return "banana"
	case 97:
		return "melon"
	case 98:
		return "pineapple"
	case 99:
		return "apple"
	case 100:
		return "pineapple"
	}
	// or just comment all above =)
	return "apple"
}

func BonusTime(salary int, bonus bool) (res string) {
	if bonus {
		salary *= 10
	}
	return "\u00A3" + strconv.Itoa(salary)
}

func MultiTable(number int) (res string) {
	for i := 1; i <= 10; i++ {
		m := i * number
		res += fmt.Sprintf("%d * %d = %d\n", i, number, m)
	}
	return strings.TrimRight(res, "\n")
}

func Namevar() string {
	var a string = "code"
	var b string = "wa.rs"
	var name string = a + b
	return name
}

func Derive(coefficient, exponent int) (res string) {
	if exponent > 2 {
		res = strconv.Itoa(coefficient*exponent) + "x^" + strconv.Itoa(exponent-1)
	} else {
		res = strconv.Itoa(coefficient * exponent)
	}
	return
}

func CalculateYears(years int) (result [3]int) {
	switch years {
	case 1:
		return [3]int{years, 15, 15}
	case 2:
		return [3]int{years, 15 + 9, 15 + 9}
	default:
		return [3]int{years, 15 + 9 + ((years - 2) * 4), 15 + 9 + ((years - 2) * 5)}
	}
}

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	var rev string

	for _, j := range str {
		rev = string(j) + rev
	}

	return str == rev
}

func Summation(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func combat(health, damage float64) float64 {
	if health-damage < 0 {
		return 0
	}
	return health - damage
}

func Move(position int, roll int) int {
	return position + (roll * 2)
}

func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}

func GetGrade(a, b, c int) rune {
	score := (a + b + c) / 3
	switch {
	case score >= 90:
		return 'A'
	case score >= 80:
		return 'B'
	case score >= 70:
		return 'C'
	case score >= 60:
		return 'D'
	case score < 60:
		return 'F'
	}
	return 'S'
}

func BinToDec(bin string) int {
	dec, _ := strconv.ParseInt(bin, 2, 32)
	return int(dec)
}

func SortVowels(s string) (res string) {
	for _, j := range s {
		switch string(j) {
		case "a", "i", "u", "e", "o", "A", "I", "U", "E", "O":
			res += "|" + string(j) + "\n"
		case "":
			res = ""
		default:
			res += string(j) + "|" + "\n"
		}
	}
	return strings.TrimSuffix(res, "\n")
}

func SimpleStringDivision(st string, k int) int {
	step := len(st) - k
	maxNum := 0

	for i := 0; i+step <= len(st); i++ {
		num, _ := strconv.Atoi(st[i : i+step])
		if num >= maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func ReduceFraction(fraction [2]int) [2]int {
	a, b := fraction[0], fraction[1]

	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	d := a + b

	return [2]int{fraction[0] / d, fraction[1] / d}
}

func MobileKeyboard(str string) (sum int) {
	clicks := map[string]int{
		"1": 1, "2": 1, "3": 1, "4": 1, "5": 1, "6": 1, "7": 1, "8": 1, "9": 1, "0": 1,
		"a": 2, "b": 3, "c": 4, "d": 2, "e": 3, "f": 4, "g": 2, "h": 3, "i": 4, "j": 2,
		"k": 3, "l": 4, "m": 2, "n": 3, "o": 4, "p": 2, "q": 3, "r": 4, "s": 5, "t": 2,
		"u": 3, "v": 4, "w": 2, "x": 3, "y": 4, "z": 5, "*": 1, "#": 1}
	for _, j := range str {
		sum += clicks[string(j)]
	}
	return
}

func Encode(str string, key int) (res []int) {
	k := strconv.Itoa(key)
	for len(k) < len(str) {
		k += k
	}
	for i := range str {
		res = append(res, int(str[i]-96+k[i]-48))
	}
	return
}

func DblLinear(n int) int {
	u := make([]int, n+1)
	u[0] = 1
	e2, e3 := 0, 0
	for i := 1; i <= n; i++ {
		u[i] = int(math.Min(float64(2*u[e2]+1), float64(3*u[e3]+1)))
		if u[i] == 2*u[e2]+1 {
			e2++
		}
		if u[i] == 3*u[e3]+1 {
			e3++
		}
	}
	return u[n]
}

func SortMyString(s string) string {
	s1, s2 := "", ""
	for i := range s {
		if i%2 == 0 {
			s1 += string(s[i])
		} else {
			s2 += string(s[i])
		}
	}
	return s1 + " " + s2
}

func AddLetters(letters []rune) rune {
	if len(letters) < 1 {
		return 'z'
	}

	var sum rune = 0
	for _, j := range letters {
		sum += j - 96

		if sum > 26 {
			sum -= 26
		}
	}
	return sum + 96
}

func Gimme(array [3]int) int {
	minN, maxN := array[0], array[0]
	for _, j := range array {
		if j <= minN {
			minN = j
		}
		if j >= maxN {
			maxN = j
		}
	}

	for i := range array {
		if array[i] > minN && array[i] < maxN {
			return i
		}
	}
	return 0
}

func inviteMoreWomen(L []int) bool {
	sum := 0
	for _, j := range L {
		sum += j
	}

	switch {
	case sum > 0:
		return true
	default:
		return false
	}
}

func SeriesSum(n int) string {
	if n != 0 {
		sum := 1.0
		for i := 1; i < n; i++ {
			sum += 1.0 / (3.0*float64(i) + 1.0)
		}
		return strconv.FormatFloat(sum, 'f', 2, 64)
	}
	return "0.00"
}

func SimpleStringCharacters(s string) []int {
	res := []int{0, 0, 0, 0}
	for _, j := range s {
		switch {
		case j >= 65 && j <= 90:
			res[0]++
		case j >= 97 && j <= 122:
			res[1]++
		case j >= 48 && j <= 57:
			res[2]++
		default:
			res[3]++
		}
	}
	return res
}

func Vaporcode(s string) (res string) {
	for i := range s {
		switch {
		case s[i] >= 'a' && s[i] <= 'z':
			res += string(s[i]-32) + "  "
		case s[i] == ' ':
			continue
		default:
			res += string(s[i]) + "  "
		}
	}
	return res[:len(res)-2]
}

func NameValue(my_list []string) (res []int) {
	sum := 0
	for i, j := range my_list {
		for _, k := range j {
			if k >= 96 && k <= 122 {
				sum += int(k - 96)
			}
		}
		res = append(res, sum*(i+1))
		sum = 0
	}
	return res
}

func RoundToNext5(n int) int {
	switch {
	case n > 0:
		switch n % 5 {
		case 1:
			n += 4
		case 2:
			n += 3
		case 3:
			n += 2
		case 4:
			n += 1
		}
	case n < 0:
		switch n % 5 {
		case -1:
			n += 1
		case -2:
			n += 2
		case -3:
			n += 3
		case -4:
			n += 4
		}
	}
	return n
}

func GetMiddle(s string) (res string) {
	if len(s)%2 == 0 {
		res = s[len(s)/2-1 : len(s)/2+1]
	} else {
		res = string(s[len(s)/2])
	}
	return
}

func GetSum(a, b int) (sum int) {
	if a > b {
		a, b = b, a
	}
	for i := a; i <= b; i++ {
		sum += i
	}
	return
}

func DNAStrand(dna string) (res string) {
	for _, j := range dna {
		switch j {
		case 'A':
			res += "T"
		case 'T':
			res += "A"
		case 'G':
			res += "C"
		case 'C':
			res += "G"
		}
	}
	return
}

func SumOfIntegersInString(strng string) (n int) {
	var res string

	for _, j := range strng {
		if j < 48 || j > 57 {
			res += " "
		} else {
			res += string(j)
		}
	}

	for _, j := range strings.Fields(res) {
		t, _ := strconv.Atoi(j)
		n += t
	}

	return
}

func Calc(s string) (res int) {
	for _, j := range s {
		if j%10 == 7 {
			res++
		}
		if j > 69 && j < 80 {
			res++
		}
	}
	return res * 6
}

func Add(n int) func(int) int {
	return func(i int) int {
		return n + i
	}
}

func Angle(n int) int {
	return (n - 2) * 180
}

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	}
	return 0
}

func BreakChocolate(n, m int) int {
	if n*m > 2 {
		return n*m - 1
	}
	return 0
}

func RemoveDuplicateWords(str string) string {
	words := make(map[string]bool)
	list := []string{}
	for _, w := range strings.Fields(str) {
		if _, value := words[w]; !value {
			words[w] = true
			list = append(list, w)
		}
	}
	return strings.Join(list, " ")
}

func DontGiveMeFive(start int, end int) (res int) {
	for i := start; i <= end; i++ {
		if !strings.Contains(strconv.Itoa(i), "5") {
			res++
		}
	}
	return
}

func CountRedBeads(n int) int {
	if n > 2 {
		return (n - 1) * 2
	}
	return 0
}

func Strong(n int) string {
	var sum uint64
	nums := strings.Split(strconv.Itoa(n), "")
	for _, j := range nums {
		num, _ := strconv.Atoi(j)
		sum += Factorial(uint64(num))
	}
	if sum == uint64(n) {
		return "STRONG!!!!"
	} else {
		return "Not Strong !!"
	}
}

func Factorial(n uint64) (res uint64) {
	if n > 0 {
		res = n * Factorial(n-1)
		return res
	}
	return 1
}

func RowSumOddNumbers(n int) (sum int) {
	start := n*(n-1) + 1
	for i := 0; n > 0; i += 2 {
		sum += start + i
		n -= 1
	}
	return
	// best solution: return n*n*n
}

func TwoToOne(s1 string, s2 string) string {
	letters := make(map[string]bool)
	var unique []string
	for _, i := range s1 + s2 {
		if _, l := letters[string(i)]; !l {
			letters[string(i)] = true
			unique = append(unique, string(i))
		}
	}
	sort.Strings(unique)
	return strings.Join(unique, "")
}

func Disemvowel(comment string) (res string) {
	for _, j := range comment {
		switch string(j) {
		case "a", "i", "u", "e", "o", "A", "I", "U", "E", "O":
			continue
		default:
			res += string(j)
		}
	}
	return
}

func EquableTriangle(a, b, c int) bool {
	p := a + b + c
	s := math.Sqrt(float64(p / 2 * (p/2 - a) * (p/2 - b) * (p/2 - c)))
	if p == int(s) {
		return true
	}
	return false
}

func ConsecutiveLetters(s string) bool {
	sl := strings.Split(s, "")
	sort.Strings(sl)
	slj := strings.Join(sl, "")
	res := true
	for i := 0; i < len(s)-1; i++ {
		if slj[i] != slj[i+1]-1 {
			res = false
		}
	}
	return res
}

func Capitalize(st string) []string {
	var s1, s2 string
	for i := 0; i < len(st); i++ {
		if i%2 != 0 {
			s1 += string(st[i])
			s2 += string(st[i] - 32)
		} else {
			s1 += string(st[i] - 32)
			s2 += string(st[i])
		}
	}
	return []string{s1, s2}
}

func LargestPower(n uint64) int {
	k := int(math.Log(float64(n)) / math.Log(3))
	if uint64(math.Pow(3, float64(k))) >= n {
		k--
	}
	return k
}

func ScrabbleScore(st string) int {
	if len(st) > 0 {
		sum := 0
		for _, j := range st {
			switch string(j) {
			case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t",
				"A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
				sum += 1
			case "d", "g", "D", "G":
				sum += 2
			case "b", "c", "m", "p", "B", "C", "M", "P":
				sum += 3
			case "f", "h", "v", "w", "y", "F", "H", "V", "W", "Y":
				sum += 4
			case "k", "K":
				sum += 5
			case "j", "x", "J", "X":
				sum += 8
			case "q", "z", "Q", "Z":
				sum += 10
			}
		}
		return sum
	}
	return 0
}

func SimpleStringReversalII(s string, a, b int) string {
	runes := []rune(s)

	if b >= len(runes) {
		b = len(runes) - 1
	}

	for i, j := a, b; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func LongestVowelChain(s string) int {
	counter := 0
	longest := 0
	for _, char := range s {
		switch char {
		case 'a', 'e', 'u', 'i', 'o':
			counter++
			if counter > longest {
				longest = counter
			}
		default:
			counter = 0
		}
	}
	return longest
}

func OverTheRoad(address int, n int) int {
	return n*2 - address + 1
}

func Digits(n uint64) int {
	return len(strconv.FormatUint(n, 10))
}

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	height := desiredHeight - upSpeed

	if height <= 0 {
		return 1
	}
	days := height / (upSpeed - downSpeed)

	if height%(upSpeed-downSpeed) > 0 {
		days++
	}
	return days + 1
}

func Divisions(n, divisor int) (counter int) {
	for n >= divisor {
		n /= divisor
		counter++
	}
	return
}

func MinimumPerimeter(area uint64) uint64 {
	for x := uint64(math.Sqrt(float64(area))); x > 0; x-- {
		if area%x == 0 {
			y := area / x
			return 2 * (x + y)
		}
	}
	return 0
}

func PickGrains(grains <-chan string) (good int, bad int) {
	for val := range grains {
		if val == "good" {
			good++
		} else {
			bad++
		}
	}
	return
}

func bandNameGenerator(word string) string {
	if word[0] == word[len(word)-1] {
		return strings.Title(word) + word[1:]
	}
	return "The " + strings.Title(word)
}

func ReverseWords2(str string) (res string) {
	str += " "
	tempStr := ""
	for _, j := range str {
		if j != ' ' {
			tempStr += string(j)
		} else {
			l := len(tempStr)
			if l > 1 {
				for i := l - 1; i >= 0; i-- {
					res += string(tempStr[i])
				}
			} else {
				res += tempStr
			}
			tempStr = ""
			res += " "
		}
	}
	return res[:len(res)-1]
}

func Incrementer(n []int) (res []int) {
	for i := range n {
		c := n[i] + i + 1
		if c > 9 {
			c %= 10
		}
		res = append(res, c)
	}
	return
}

func ModifyMultiply(str string, loc, num int) string {
	s := strings.Fields(str)
	if num > 0 {
		m := []string{}
		for i := 0; i < num; i++ {
			m = append(m, s[loc])
		}
		return strings.Join(m, "-")
	}
	return s[loc]
}

func SortByLength(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	return arr
}

func EncodeCd(n uint8) string {
	res := "P"
	num := fmt.Sprintf("%08b", n)
	sw := true
	w := string(res[0])
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] == '1' {
			if sw {
				w = "L"
				sw = false
			} else {
				w = "P"
				sw = true
			}
			res += w
		} else {
			res += w
		}
	}
	return res
}

func ClosestMultipleOf10(n uint32) uint32 {
	return (n + 5) / 10 * 10
}

func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2
	h := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return h
}

func ReverseList(lst []int) []int {
	for i, j := 0, len(lst)-1; i < j; i, j = i+1, j-1 {
		lst[i], lst[j] = lst[j], lst[i]
	}
	return lst
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	res := p0
	i := 0
	if percent == 0 {
		i--
	}

	for res <= p {
		res = res + int(float64(res)*percent/100) + aug
		i++
	}
	return i
}

func StantonMeasure(arr []int) int {
	n := arr[0]
	c1, c2 := 0, 0
	for _, i := range arr {
		if i == n {
			c1++
		}
	}
	for _, j := range arr {
		if j == c1 {
			c2++
		}
	}
	return c2
}

func Brightest(colors []string) string {
	var brightest int
	var b int
	for i, color := range colors {
		c, _ := hex.DecodeString(color[1:])
		n := valueOfColor(c)
		if n > brightest {
			brightest = n
			b = i
		}
	}
	return colors[b]
}

func valueOfColor(color []byte) int {
	m := color[0]
	if color[1] > m {
		m = color[1]
	}
	if color[2] > m {
		m = color[2]
	}
	return int(m)
}

func Dominator(a []int) int {
	m := make(map[int]int)
	for _, j := range a {
		if _, val := m[j]; val {
			m[j]++
		} else {
			m[j] = 1
		}
	}
	n, k := 0, 0
	for i, j := range m {
		if j > n {
			n = j
			k = i
		}
	}

	if n > len(a)/2 {
		return k
	}

	return -1
}

func Alternate(n int, firstValue string, secondValue string) (res []string) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			res = append(res, firstValue)
		} else {
			res = append(res, secondValue)
		}
	}
	return
}

func Capitalize2(st string, arr []int) (res string) {
	s := strings.Split(st, "")
	for _, j := range arr {
		if j > len(st)-1 {
			break
		}
		s[j] = strings.ToUpper(s[j])
	}
	return strings.Join(s, "")
}

func FixStringCase(str string) string {
	l, u := 0, 0
	for _, r := range str {
		if r >= 'a' && r <= 'z' {
			l++
		} else {
			u++
		}
	}

	if l >= u {
		return strings.ToLower(str)
	}
	return strings.ToUpper(str)
}

func FindNextPower(val, pow int) int {
	x, y := float64(val), float64(pow)
	r1 := int(math.Pow(math.Floor(math.Pow(x, 1.0/y))+1.0, y))
	r2 := int(math.Pow(math.Ceil(math.Pow(x, 1.0/y))+1.0, y))
	if r1 <= val {
		return r2
	}
	return r1
}

func BabySharkLyrics() string {
	w, l, h, d, n, s := []string{"Baby", "Mommy", "Daddy", "Grandma", "Grandpa"}, "", "Let's go hunt", ","+strings.Repeat(" doo", 6), "\n", " shark"
	func() {
		for _, v := range w {
			l += strings.Repeat(v+s+d+n, 3)
			l += v + s + "!" + n
		}
	}()
	return l + strings.Repeat(h+d+n, 3) + h + "!" + n + "Run away,…\n"
}

func NextPrime(n int) int {
	if n == 0 {
		return 2
	}
	n++
	for !NumIsPrime(n) {
		n++
	}
	return n
}

func NumIsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PassHash(str string) string {
	s := md5.Sum([]byte(str))
	return hex.EncodeToString(s[:])
}

func WrapPresent(height, width, length int) int {
	sizes := []int{height, width, length}
	sort.Ints(sizes)
	return sizes[0]*4 + sizes[1]*2 + sizes[2]*2 + 20
}

func XMasTree(height int) []string {
	tree := make([]string, height+2)
	halfWidth := height - 1
	h := 1
	for i := 0; i < height; i++ {
		tree[i] = strings.Repeat("_", halfWidth) + strings.Repeat("#", h) + strings.Repeat("_", halfWidth)
		h += 2
		halfWidth--

	}
	tree[len(tree)-1], tree[len(tree)-2] = tree[0], tree[0]
	return tree
}

type NBAPlayer struct {
	Team string
	Ppg  float64
}

func SumPpg(playerOne, playerTwo NBAPlayer) float64 {
	return playerOne.Ppg + playerTwo.Ppg
}

func Mirror(data []int) []int {
	res, data2 := make([]int, len(data)), make([]int, len(data))
	// Your solution should not mutate the input
	copy(data2, data) // Okay, okay
	sort.Ints(data2)
	copy(res, data2)
	for i := len(data2) - 2; i >= 0; i-- {
		res = append(res, data2[i])
	}
	return res
}

func AlphabetSymmetry(slice []string) (res []int) {
	for _, word := range slice {
		c := 0
		for i := 0; i < len(word); i++ {
			if int(word[i]) == i+97 || int(word[i]) == i+65 {
				c++
			}
		}
		res = append(res, c)
	}
	return
}

func Dative(word string) string {
	runes := []rune(word)
	for i := len(runes) - 1; i >= 0; i-- {
		l := string(runes[i])
		switch l {
		case "a", "á", "o", "ó", "u", "ú":
			return word + "nak"

		}
	}
	return word + "nek"
}

func Deemojify(emote_string string) (res string) {
	var nums []string
	words := strings.Split(emote_string, "  ")
	for _, word := range words {
		emojis := strings.Split(word, " ")
		letter := ""
		for _, emoji := range emojis {
			switch emoji {
			case ":)":
				letter += "0"
			case ":D":
				letter += "1"
			case ">(":
				letter += "2"
			case ">:C":
				letter += "3"
			case ":/":
				letter += "4"
			case ":|":
				letter += "5"
			case ":O":
				letter += "6"
			case ";)":
				letter += "7"
			case "^.^":
				letter += "8"
			case ":(":
				letter += "9"
			}
		}
		nums = append(nums, letter)
	}

	for num := range nums {
		n, _ := strconv.Atoi(nums[num])
		res += string(rune(n))

	}
	return
}

func RangeBitCount(a, b int) (res int) {
	for i := a; i <= b; i++ {
		n := fmt.Sprintf("%b", i)
		for _, j := range n {
			d, _ := strconv.Atoi(string(j))
			res += d
		}
	}
	return
}

func Is_valid_ip(ip string) bool {
	m := strings.Split(ip, ".")
	if len(m) != 4 {
		return false
	}
	for _, sec := range m {
		if len(sec) > 3 {
			return false
		}

		if strings.HasPrefix(sec, "0") && len(sec) > 1 {
			return false
		}

		n, err := strconv.Atoi(sec)
		if err != nil {
			return false
		}

		if n < 0 || n > 255 {
			return false
		}
	}
	return true
}

func HoopCount(n int) string {
	if n < 10 {
		return "Keep at it until you get it"
	}
	return "Great, now move on to tricks"
}

func Contamination(text, char string) (res string) {
	for i := 0; i < len(text); i++ {
		res += char
	}
	return
}

func Quadratic(x1, x2 int) (res [3]int) {
	return [3]int{1, -(x1 + x2), x1 * x2}
}

func ToAlternatingCase(str string) string {
	var result strings.Builder
	for _, r := range str {
		if unicode.IsLower(r) {
			result.WriteRune(unicode.ToUpper(r))
		} else if unicode.IsUpper(r) {
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func NearestSq(n int) int {
	r := math.Round(math.Sqrt(float64(n)))
	return int(r * r)
}

func TwoSort(arr []string) (res string) {
	sort.Strings(arr)
	for i := 0; i < len(arr[0])-1; i++ {
		res += string(arr[0][i]) + "***"
	}
	return res + string(arr[0][len(arr[0])-1])
}

func ExpressionMatter(a int, b int, c int) int {
	values := []int{a * (b + c), a * b * c, a + b*c, (a + b) * c, a + b + c}
	sort.Ints(values)
	return values[len(values)-1]
}

func CircleOfNumbers(n int, firstNumber int) int {
	op := n/2 + firstNumber
	if op >= n {
		return op - n
	}
	return op
}

func Accum(s string) (res string) {
	res += strings.ToUpper(string(s[0])) + "-"
	for i := 1; i < len(s); i++ {
		res += strings.ToUpper(string(s[i]))
		res += strings.Repeat(strings.ToLower(string(s[i])), i)
		res += "-"
	}
	return strings.TrimSuffix(res, "-")
}

func High(s string) string {
	words := strings.Split(s, " ")
	var scores []int

	for _, word := range words {
		score := 0
		for _, letter := range word {
			score += int(letter) - 96
		}
		scores = append(scores, score)
	}

	firstMax, index := 0, 0
	for i := range scores {
		if scores[i] > firstMax {
			firstMax = scores[i]
			index = i
		}
	}

	return words[index]
}

func TwoSum(numbers []int, target int) [2]int {
	values := make(map[int]int)
	for i := range numbers {
		_, exists := values[numbers[i]]
		if exists {
			return [2]int{values[numbers[i]], i}
		} else {
			values[target-numbers[i]] = i
		}
	}
	return [2]int{}
}

const PI = 3.141592653589793

func AreaOfPolygonInsideCircle(circleRadius float64, numberOfSides int) float64 {
	s := (float64(numberOfSides) * circleRadius * circleRadius * math.Sin(2*PI/float64(numberOfSides))) / 2
	return math.Round(s*1000) / 1000
}

func RGB(r, g, b int) (res string) {
	colors := []int{r, g, b}
	for c := range colors {
		switch {
		case colors[c] < 0:
			colors[c] = 0
		case colors[c] > 255:
			colors[c] = 255
		}
		h := fmt.Sprintf("%X", colors[c])
		if len(h) == 1 {
			res += "0" + h
		} else {
			res += h
		}
	}
	return res
}

func LastDigit(n1, n2 string) int {
	if n1 == "0" && n2 == "0" {
		return 1
	}
	bigN1 := new(big.Int)
	bigN2 := new(big.Int)
	bigN1.SetString(n1, 10)
	bigN2.SetString(n2, 10)
	lastDigitN1 := new(big.Int).Mod(bigN1, big.NewInt(10)).Int64()

	cycles := map[int][]int{
		0: {0},
		1: {1},
		2: {2, 4, 8, 6},
		3: {3, 9, 7, 1},
		4: {4, 6},
		5: {5},
		6: {6},
		7: {7, 9, 3, 1},
		8: {8, 4, 2, 6},
		9: {9, 1},
	}

	cycle := cycles[int(lastDigitN1)]
	cycleLength := len(cycle)

	if bigN2.Cmp(big.NewInt(0)) == 0 {
		return 1
	}

	bigCycleLength := big.NewInt(int64(cycleLength))
	exp := new(big.Int).Mod(bigN2, bigCycleLength).Int64() - 1

	if exp == -1 {
		exp = int64(cycleLength - 1)
	}

	return cycle[exp]
}

func Uniq(a []string) (res []string) {
	if len(a) == 0 {
		return nil
	}

	res = append(res, a[0])
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			res = append(res, a[i])
		}
	}
	return
}

func seatsInTheater(nCols int, nRows int, col int, row int) int {
	return (nCols - col + 1) * (nRows - row)
}

func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
}

func QueueTime(customers []int, n int) int {
	tills := make([]int, n)

	for _, minutes := range customers {
		mv, mi := tills[0], 0
		for i, j := range tills {
			if j <= mv {
				mv = j
				mi = i
			}
		}
		tills[mi] += minutes
	}

	mx := tills[0]
	for _, el := range tills {
		if el > mx {
			mx = el
		}
	}
	return mx
}

func FindUniq(arr []float32) float32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	if arr[1] == arr[0] && arr[1] == arr[len(arr)-1] {
		return -1
	} else if arr[1] == arr[len(arr)-1] {
		return arr[0]
	}
	return arr[len(arr)-1]
}

func ToNato(words string) (res string) {
	natoAlphabet := map[string]string{
		"A": "Alfa",
		"B": "Bravo",
		"C": "Charlie",
		"D": "Delta",
		"E": "Echo",
		"F": "Foxtrot",
		"G": "Golf",
		"H": "Hotel",
		"I": "India",
		"J": "Juliett",
		"K": "Kilo",
		"L": "Lima",
		"M": "Mike",
		"N": "November",
		"O": "Oscar",
		"P": "Papa",
		"Q": "Quebec",
		"R": "Romeo",
		"S": "Sierra",
		"T": "Tango",
		"U": "Uniform",
		"V": "Victor",
		"W": "Whiskey",
		"X": "Xray",
		"Y": "Yankee",
		"Z": "Zulu",
		".": ".",
		",": ",",
		":": ":",
		";": ";",
		"!": "!",
		"?": "?",
	}

	wrds := strings.Split(words, " ")

	for _, word := range wrds {
		for _, letter := range word {
			res += natoAlphabet[strings.ToUpper(string(letter))] + " "
		}
	}
	res = strings.TrimSuffix(res, " ")
	return
}

func PrimeFactors(n int) (res []int) {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			res = append(res, i)
			n /= i
		}
	}
	if n > 1 {
		res = append(res, n)
	}
	return res
}

func Smaller(arr []int) []int {
	res := make([]int, len(arr))

	for i, j := range arr {
		counter := 0
		for _, k := range arr[i+1:] {
			if k < j {
				counter++
			}
		}
		res[i] = counter
	}
	return res
}

func LengthOfSequence(arr []int, key int) int {
	start, stop, counter := -1, -1, 0
	for i, j := range arr {
		if j == key {
			counter++
		}

		if j == key && start == -1 {
			start = i
		} else if j == key && start != -1 {
			stop = i
		}
	}

	if counter != 2 {
		return 0
	}

	return len(arr[start : stop+1])
}

func LeastLarger(a []int, i int) int {
	res := 0
	r := int(^uint(0) >> 1)

	for idx, j := range a {
		if j > a[i] && j-a[i] < r {
			r = j - a[i]
			res = idx
		}
	}

	if r == -1 {
		return -1
	}

	return res
}

func ToJadenCase(str string) string {
	return strings.Title(str)
}

func Decode(code []int, key int) string {
	res := ""
	keys := strings.Split(strconv.Itoa(key), "")
	k := 0

	for _, j := range code {
		n, _ := strconv.Atoi(keys[k])
		res += string(rune(j - n + 96))
		k++
		if k > len(keys)-1 {
			k = 0
		}
	}

	return res
}

func SpinningRings(innerMax, outerMax int) int {
	res := 1
	for innerMax-(res-1)%(innerMax+1) != res%(outerMax+1) {
		res++
	}
	return res
}

func Histogram(results [6]int) (res string) {
	for i := 5; i >= 0; i-- {
		subs := strconv.Itoa(i+1) + "|"
		if results[i] != 0 {
			subs += strings.Repeat("#", results[i]) + " " + strconv.Itoa(results[i]) + "\n"
		} else {
			subs += "\n"
		}
		res += subs
	}
	return
}

func Recursion101(a, b int) []int {
	fmt.Println(a, b)
	if a == 0 || b == 0 {
		return []int{a, b}
	}

	if a >= 2*b {
		a = a - 2*b
		return Recursion101(a, b)
	}

	if b >= 2*a {
		b = b - 2*a
		return Recursion101(a, b)
	}

	return []int{a, b}
}

func fibStrings(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "01"
	} else {
		return fibStrings(n-1) + fibStrings(n-2)
	}
}

func alphanumeric(str string) bool {
	if len(str) < 1 {
		return false
	}

	for _, j := range str {
		if !unicode.IsLetter(j) && !unicode.IsNumber(j) {
			return false
		}
	}

	return true
}

func ToCsvText(array [][]int) string {
	if len(array) == 0 {
		return ""
	}
	var res string
	for _, i := range array {
		var row string
		for _, j := range i {
			col := strconv.Itoa(j)
			row += col + ","
		}
		res += row[:len(row)-1] + "\n"
	}
	return res[:len(res)-1]
}

func EachCons(arr []int, n int) (res [][]int) {
	for i := 0; i < len(arr)-n+1; i++ {
		t := make([]int, n)
		for j := 0; j < n; j++ {
			t[j] = arr[i+j]
		}
		res = append(res, t)
	}
	return
}

// Use the preloaded Tuple struct as return type
type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	if len(text) == 0 {
		return []Tuple{}
	}

	var res []Tuple
	m := make(map[int32]int)
	ord := ""
	for _, char := range text {
		if _, ok := m[char]; ok {
			m[char]++
		} else {
			m[char] = 1
			ord += string(char)
		}
	}

	for _, j := range ord {
		res = append(res, Tuple{j, m[j]})
	}

	return res
}

func MaxMultiple(d, b int) int {
	for i := b; i > 0; i-- {
		if i%d == 0 {
			return i
		}
	}
	return 0
}

func Number(busStops [][2]int) int {
	num := 0
	for _, stop := range busStops {
		num += stop[0] - stop[1]
	}
	return num
}

func NbDig(n int, d int) (c int) {
	for i := 0; i <= n; i++ {
		c += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
	return
}

func GuessHatColor(a, b, c, d string) int {
	if b == c {
		return 1
	}
	return 2
}

func AlexMistakes(numberOfKatas, timeLimit int) int {
	return int(math.Log2(float64((timeLimit-numberOfKatas*6)/5 + 1)))
}

func ToCamelCase2(s string) string {
	s = strings.ReplaceAll(s, "-", "_")
	words := strings.Split(s, "_")
	res := words[0]
	for _, word := range words[1:] {
		res += strings.Title(word)
	}
	return strings.TrimSuffix(res, "_")
}

func BowlingPins(slice []int) string {
	m := make(map[int]string)
	m[1] = "I"
	m[2] = "I"
	m[3] = "I"
	m[4] = "I"
	m[5] = "I"
	m[6] = "I"
	m[7] = "I"
	m[8] = "I"
	m[9] = "I"
	m[10] = "I"

	for _, j := range slice {
		if _, ok := m[j]; ok {
			m[j] = " "
		}
	}

	return m[7] + " " + m[8] + " " + m[9] + " " + m[10] + "\n" +
		" " + m[4] + " " + m[5] + " " + m[6] + " " + "\n" +
		"  " + m[2] + " " + m[3] + "  " + "\n" +
		"   " + m[1] + "   "
}

func FizzBuzzCuckooClock(time string) string {
	st := strings.Split(time, ":")
	h, _ := strconv.Atoi(st[0])
	m, _ := strconv.Atoi(st[1])

	switch {
	case m == 30:
		return "Cuckoo"
	case m == 0:
		{
			if h > 12 {
				h -= 12
			} else if h == 0 {
				h = 12
			}
			return strings.TrimSuffix(strings.Repeat("Cuckoo ", h), " ")
		}
	case m%15 == 0:
		return "Fizz Buzz"
	case m%3 == 0:
		return "Fizz"
	case m%5 == 0:
		return "Buzz"
	default:
		return "tick"
	}
}

func Persistence(n int) (counter int) {
	num := strconv.Itoa(n)
	for len(num) != 1 {
		nums := strings.Split(num, "")
		res := 1
		for _, d := range nums {
			cd, _ := strconv.Atoi(d)
			res *= cd
		}
		num = strconv.Itoa(res)
		counter++
	}
	return counter
}

func Unlock(str string) (res string) {
	m := make(map[string]string)
	m["a"] = "2"
	m["b"] = "2"
	m["c"] = "2"
	m["d"] = "3"
	m["e"] = "3"
	m["f"] = "3"
	m["g"] = "4"
	m["h"] = "4"
	m["i"] = "4"
	m["j"] = "5"
	m["k"] = "5"
	m["l"] = "5"
	m["m"] = "6"
	m["n"] = "6"
	m["o"] = "6"
	m["p"] = "7"
	m["q"] = "7"
	m["r"] = "7"
	m["s"] = "7"
	m["t"] = "8"
	m["u"] = "8"
	m["v"] = "8"
	m["w"] = "9"
	m["x"] = "9"
	m["y"] = "9"
	m["z"] = "9"

	for _, l := range str {
		c := strings.ToLower(string(l))
		res += m[string(c)]
	}

	return
}

func ShortestArrang(n int) []int {
	if n < 3 {
		return []int{-1}
	}

	for i := 1; i*(i+1)/2 <= n; i++ {
		sum := i * (i + 1) / 2
		if (n-sum)%(i+1) == 0 {
			start := (n - sum) / (i + 1)
			result := make([]int, i+1)
			for j := 0; j <= i; j++ {
				result[j] = start + i - j
			}
			return result
		}
	}

	return []int{-1}
}

func MakeValley(arr []int) []int {
	arr2 := make([]int, len(arr))
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	c1, c2 := 0, len(arr)-1
	for k, v := range arr {
		if (k % 2) == 0 {
			arr2[c1] = v
			c1++
		} else {
			arr2[c2] = v
			c2--
		}
	}
	return arr2
}

func RacePodium(blocks int) [3]int {
	var p1, p2, p3 int
	nf := float64(int(blocks / 3))
	d := int(math.Round((float64(blocks)/3 - nf) * 3))
	n := int(nf)

	switch {
	case d == 1:
		p1 = n + d + 1
		p2 = n + 1
		p3 = n - 2
	case d > 1:
		p1 = n + d
		p2 = n + 1
		p3 = n - 1
	default:
		p1 = n + 1
		p2 = n
		p3 = n - 1
	}

	if p3 == 0 {
		p2 = p2 - 1
		p3 = 1
	}

	return [3]int{p2, p1, p3}
}

func StringPrefixAndSuffix(s string) int {
	for i := len(s) / 2; i > 0; i-- {
		if s[:i] == s[len(s)-i:] {
			return i
		}
	}
	return 0
}

func Scale(s string, k, n int) (res string) {
	var w string

	for i := range s {
		if s[i] != 10 {
			for j := k; j > 0; j-- {
				w += string(s[i])
			}
		}

		if (s[i] == 10) || len(s)-1 == i {
			var w2 string

			for j := n; j > 0; j-- {
				w2 += w + string(byte(10))
			}
			if len(s)-1 == i {
				w2 = w2[:len(w2)-1]
			}
			res += w2
			w = ""
		}
	}
	return res
}

func main() {
	fmt.Println("Codewars")
	fmt.Println(Scale("Kj\nSH", 1, 2))
	fmt.Println("Kj\nKj\nSH\nSH")
	//fmt.Println(StringPrefixAndSuffix("abcabc"), 3)
	//fmt.Println(RacePodium(11), [3]int{4, 5, 2})
	//fmt.Println(RacePodium(10), [3]int{4, 5, 1})
	//fmt.Println(MakeValley([]int{17, 17, 15, 14, 8, 7, 7, 5, 4, 4, 1}))
	//fmt.Println([]int{17, 15, 8, 7, 4, 1, 4, 5, 7, 14, 17})
	//fmt.Println(ShortestArrang(14), []int{5, 4, 3, 2})
	//fmt.Println(Unlock("Nokia"), "66542")
	//fmt.Println(Persistence(999), 4)
	//fmt.Println(FizzBuzzCuckooClock("11:15"), "| Fizz Buzz")
	//fmt.Println(FizzBuzzCuckooClock("03:03"), "| Fizz")
	//fmt.Println(FizzBuzzCuckooClock("14:30"), "| Cuckoo")
	//fmt.Println(FizzBuzzCuckooClock("00:00"), "| Cuckoo 12")
	//fmt.Println(BowlingPins([]int{3, 5, 9}))
	//fmt.Println("I I   I\n I   I \n  I    \n   I   ")
	//fmt.Println(ToCamelCase2("The_Stealth_Warrior"), "TheStealthWarrior")
	//fmt.Println(AlexMistakes(11, 120), 3)
	//fmt.Println(GuessHatColor("white", "black", "white", "black"), 2)
	//fmt.Println(NbDig(550, 5), 213)
	//fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}), 17)
	//fmt.Println(MaxMultiple(37, 200), 185)
	//fmt.Println(OrderedCount("Code Wars"))
	//fmt.Println([]Tuple{Tuple{'C', 1}, Tuple{'o', 1}, Tuple{'d', 1}, Tuple{'e', 1}, Tuple{' ', 1}, Tuple{'W', 1}, Tuple{'a', 1}, Tuple{'r', 1}, Tuple{'s', 1}})
	//fmt.Println(OrderedCount("abracadabra"))
	//fmt.Println([]Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}})
	//fmt.Println(EachCons([]int{3, 5, 8, 13}, 2), [][]int{{3, 5}, {5, 8}, {8, 13}})
	//fmt.Println(EachCons([]int{3, 5, 8, 13}, 1), [][]int{{3}, {5}, {8}, {13}})
	//fmt.Println(
	//	ToCsvText([][]int{
	//		{0, 1, 2, 3, 45},
	//		{10, 11, 12, 13, 14},
	//		{20, 21, 22, 23, 24},
	//		{30, 31, 32, 33, 34}}), "0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34")
	//fmt.Println(alphanumeric("ciao\n$$_"), false)
	//fmt.Println(fibStrings(5), "0100101001001")
	//fmt.Println(Recursion101(8796203, 7556), []int{1019, 1442})
	//fmt.Println(Histogram([6]int{7, 3, 10, 1, 0, 5}))
	//fmt.Println("6|##### 5\n5|\n4|# 1\n3|########## 10\n2|### 3\n1|####### 7\n")
	//fmt.Println(SpinningRings(2, 3), 5)
	//fmt.Println(Decode([]int{14, 10, 22, 29, 6, 27, 19, 18, 6, 12, 8}, 1939), "masterpiece")
	//fmt.Println(ToJadenCase("most trees are blue"), "Most Trees Are Blue")
	//fmt.Println(LeastLarger([]int{4, 1, 3, 5, 6}, 0), 3)
	//fmt.Println(LengthOfSequence([]int{0, -3, 7, 4, 0, 3, 7, 9}, 7), 5)
	//fmt.Println(LengthOfSequence([]int{7, 1, 7, 1, 7}, 7), 0)
	//fmt.Println(Smaller([]int{5, 4, 7, 9, 2, 4, 4, 5, 6}), []int{4, 1, 5, 5, 0, 0, 0, 0, 0})
	//fmt.Println(PrimeFactors(12), []int{2, 2, 3})
	//fmt.Println(ToNato("If you can read") == "India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta")
	//fmt.Println(FindUniq([]float32{0, 0, 0.55, 0, 0}), 0.55)
	//fmt.Println(QueueTime([]int{2, 2, 3, 3, 4, 4}, 2), 9)
	//fmt.Println(Uniq([]string{"a", "a", "b", "b", "c", "a", "b", "c", "c"}), []string{"a", "b", "c", "a", "b", "c"})
	//fmt.Println(LastDigit("9", "7"), 9)
	//fmt.Println(RGB(-20, 275, 125), "00FF7D")
	//fmt.Println(RGB(0, 0, 0), "000000")
	//fmt.Println(AreaOfPolygonInsideCircle(3, 3), 11.691)
	//fmt.Println(TwoSum([]int{1234, 5678, 9012}, 14690), "[1 2]")
	//fmt.Println(High("what time are we climbing up the volcano"), "volcano")
	//fmt.Println(High("man i need a taxi up to ubud"), "taxi")
	//fmt.Println(Accum("ZpglnRxqenU"), "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu")
	//fmt.Println(CircleOfNumbers(10, 7), 2)
	//fmt.Println(ExpressionMatter(3, 5, 7), 105)
	//fmt.Println(TwoSort([]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}), "a***r***e")
	//fmt.Println(NearestSq(10), 9)
	//fmt.Println(ToAlternatingCase("HeLLo WoRLD") == "hEllO wOrld")
	//fmt.Println(Quadratic(1, 2) == [3]int{1, -3, 2})
	//fmt.Println(Contamination("abc", "z"), "zzz")
	//fmt.Println(HoopCount(3) == "Keep at it until you get it")
	//fmt.Println(HoopCount(11) == "Great, now move on to tricks")
	//fmt.Println(Is_valid_ip("abc.def.ghi.jkl"), false)
	//fmt.Println(Is_valid_ip("127.1.1.0"), true)
	//fmt.Println(Is_valid_ip("123.045.067.089"), false)
	//fmt.Println(RangeBitCount(2, 7), 11)
	//fmt.Println(Deemojify(":D :) :/  :D :) :|"), "hi")
	//fmt.Println(Deemojify(";) >(  :D :) :D  :D :) ^.^  :D :) ^.^  :D :D :D  >:C >(  :D :D :(  :D :D :D  :D :D :/  :D :) ^.^  :D :) :)  >:C >:C"), "Hello world!")
	//fmt.Println(Dative("ablak"), "ablaknak")
	//fmt.Println(Dative("tükör"), "tükörnek")
	//fmt.Println(Dative("virág"), "virágnak")
	//fmt.Println(AlphabetSymmetry([]string{"abode", "ABc", "xyzD"}), []int{4, 3, 1})
	//fmt.Println(Mirror([]int{-5, 10, 8, 10, 2, -3, 10}), []int{-5, -3, 2, 8, 10, 10, 10, 10, 10, 8, 2, -3, -5})
	//for _, j := range XMasTree(5) {
	//	fmt.Println(j)
	//}
	//fmt.Println(WrapPresent(1, 3, 1), 32)
	//fmt.Println(PassHash("password") == "5f4dcc3b5aa765d61d8327deb882cf99")
	//fmt.Println(NextPrime(0), 1)
	//fmt.Println(NextPrime(5), 7)
	//fmt.Println(NextPrime(911), 919)
	//fmt.Println(BabySharkLyrics())
	//fmt.Println(FindNextPower(12385, 3), 13824)
	//fmt.Println(FindNextPower(1245678, 5), 1419857)
	//fmt.Println(FindNextPower(11390625, 6), 16777216)
	//fmt.Println(FixStringCase("coDE"), "code")
	//fmt.Println(Capitalize2("abcdef", []int{1, 2, 5, 100}), "aBCdeF")
	//fmt.Println(Alternate(5, "true", "false"), []string{"true", "false", "true", "false", "true"})
	//fmt.Println(Dominator([]int{3, 4, 3, 2, 3, 1, 3, 3}), 3)
	//fmt.Println(Dominator([]int{1, 2, 3, 4, 5}), -1)
	//fmt.Println(Brightest([]string{"#ABCDEF", "#123456"}), "#ABCDEF")
	//fmt.Println(Brightest([]string{"#FFFFFF", "#123456", "#000000"}), "#FFFFFF")
	//fmt.Println(StantonMeasure([]int{1, 4, 3, 2, 1, 2, 3, 2}), 3)
	//fmt.Println(StantonMeasure([]int{1, 4, 3, 0, 1, 9, 3, 6}), 0)
	//fmt.Println(NbYear(1500, 5, 100, 5000), 15)
	//fmt.Println(NbYear(1500000, 0, 10000, 2000000), 50)
	//fmt.Printf("%d\n%d\n", ReverseList([]int{1, 2, 3, 4}), []int{4, 3, 2, 1})
	//fmt.Println(Heron(3.0, 4.0, 5.0), 6.0)
	//fmt.Println(ClosestMultipleOf10(22), 20)
	//fmt.Println(EncodeCd(5), "PLLPPPPPP")
	//fmt.Println(EncodeCd(16), "PPPPPLLLL")
	//fmt.Printf("%s\n%s\n", SortByLength([]string{"dog", "food", "a", "of"}), []string{"a", "of", "dog", "food"})
	//fmt.Println(ModifyMultiply("This is a string", 3, 5) == "string-string-string-string-string")
	//fmt.Println(ModifyMultiply("LOctufVMZFzLdYnd SdEQhtaSsjlgsqIhIJgQTZ", 0, 0) == "LOctufVMZFzLdYnd")
	//fmt.Println(Incrementer([]int{3, 6, 9, 8, 9}), []int{4, 8, 2, 2, 4})
	//fmt.Println(ReverseWords2("double  spaced  words") == "elbuod  decaps  sdrow")
	//fmt.Println(ReverseWords2("a b c d") == "a b c d")
	//fmt.Println(bandNameGenerator("knife"), "The Knife")
	//fmt.Println(MinimumPerimeter(45) == 28)
	//fmt.Println(Divisions(100, 2) == 6)
	//fmt.Println(GrowingPlant(100, 10, 910), 10)
	//fmt.Println(Digits(18446744073709551615) == 20)
	//fmt.Println(OverTheRoad(1, 3) == 6)
	//fmt.Println(LongestVowelChain("codewarriors") == 2)
	//fmt.Println(SimpleStringReversalII("codingIsFun", 2, 100) == "conuFsIgnid")
	//fmt.Println(ScrabbleScore("street") == 6)
	//fmt.Println(LargestPower(82))
	//fmt.Println(Capitalize("codewars"), "CoDeWaRs", "cOdEwArS")
	//fmt.Println(ConsecutiveLetters("dabc") == true)
	//fmt.Println(EquableTriangle(5, 12, 13) == true)
	//fmt.Println(Disemvowel("This website is for losers LOL!") == "Ths wbst s fr lsrs LL!")
	//fmt.Println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding") == "abcdefghilnoprstu")
	//fmt.Println(RowSumOddNumbers(4), 64)
	//fmt.Println(Strong(145), "STRONG!!!!")
	//fmt.Println(CountRedBeads(5), 8)
	//fmt.Println(DontGiveMeFive(-55, 12), 56)
	//fmt.Println(RemoveDuplicateWords("alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta"), "alpha beta gamma delta")
	//fmt.Println(BreakChocolate(5, 5), 24)
	//fmt.Println(SortNumbers([]int{1, 2, 10, 50, 5}), "[1 2 5 10 50]")
	//fmt.Println(TwoOldestAges([]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}), "[95 99]")
	//fmt.Println(Angle(3), 180)
	//fmt.Println(Add(1)(3), 4)
	//fmt.Println(Calc("F7$&QE?M"), 18)
	//fmt.Println(SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"), 3635)
	//fmt.Println(DNAStrand("ATTGC"), "TAACG")
	//fmt.Println(GetSum(321, 123), "44178")
	//fmt.Println(GetMiddle("test"), "es")
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
		fmt.Println(GetPlanetName(3)7, "Earth")
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
		fmt.Println(AbbrevName("David Mendieta"), "D.M")
		fmt.Println(OddCount(15023), 7511)
		fmt.Println(QuarterOf(3), 1)
		fmt.Println(century(1990), 20)
		fmt.Println(ReverseWords("yoda doesn't speak like this") == "this like speak doesn't yoda")
		fmt.Println(NthEven(3), 4)
		fmt.Println(countSheep(3))
		fmt.Println(Past(1, 1, 1), 3661000)
		fmt.Println(NoSpace("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd"), "88Bifk8hB8BB8BBBB888chl8BhBfd")
		fmt.Println(SubtractSum(789278917), "apple")
		fmt.Println(BonusTime(100, false), "100\u00A3")
		fmt.Println(MultiTable(5))
		fmt.Println(Derive(5, 9), "45x^8")
		fmt.Println(CalculateYears(10), [3]int{10, 56, 64})
		fmt.Println(IsPalindrome("Abba"), true)
		fmt.Println(Summation(213), 22791)
		fmt.Println(BinToDec("1001001"), 73)
		fmt.Println(SortVowels("Codewars"), "C|\\n|o\\nd|\\n|e\\nw|\\n|a\\nr|\\ns|")
		fmt.Println(SimpleStringDivision("1234", 1), 234)
		fmt.Println(ReduceFraction([2]int{80, 120}), "2, 3")
		fmt.Println(MobileKeyboard("codewars"), 26)
		fmt.Println(Encode("scout", 19391), "20, 12, 18, 30, 21")
		fmt.Println(DblLinear(10), 22)
		fmt.Println(SortMyString("CodeWars"), "CdWr oeas")
		fmt.Println(AddLetters([]rune{'v', 'b', 'f', 'b', 'h', 't', 'd'}), 108)
		fmt.Println(Gimme([3]int{5, 10, 14}), 1)
		fmt.Println(inviteMoreWomen([]int{1, -1, 1}), true)
		fmt.Println(SeriesSum(4), "1.49")
		fmt.Println(SimpleStringCharacters("bgA5<1d-tOwUZTS8yQ"), []int{7, 6, 3, 2})
		fmt.Println(Vaporcode("Why isnt my code working"), "W  H  Y  I  S  N  T  M  Y  C  O  D  E  W  O  R  K  I  N  G")
		fmt.Println(NameValue([]string{"abc", "abc", "abc", "abc"}), []int{6, 12, 18, 24})
		fmt.Println(RoundToNext5(-21), 25)
	*/
}
