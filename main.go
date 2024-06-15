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

func main() {
	fmt.Println("Codewars")
	fmt.Println(Dative("ablak"), "ablaknak")
	fmt.Println(Dative("tükör"), "tükörnek")
	fmt.Println(Dative("virág"), "virágnak")
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
