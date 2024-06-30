package lib

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func GetInput() string {
	day := os.Args[1]
	var inputFileName string
	if len(os.Args) > 2 && os.Args[2] == "test" {
		inputFileName = "test"
	} else {
		inputFileName = "in"
	}
	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/cmd/%s/%s.txt", pwd, day, inputFileName)
	b, _ := os.ReadFile(path)
	return string(b)
}

var NumRegex = regexp.MustCompile(`-?\d+`)

func ParseNums(s string) []int {
	ss := NumRegex.FindAllString(s, -1)
	nums := make([]int, len(ss))
	for i, str := range ss {
		n, _ := strconv.Atoi(str)
		nums[i] = n
	}
	return nums
}

type Vector struct {
	X, Y int
}

func (v *Vector) Hash() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}

var DIRS = []Vector{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func Permutations[T any](a []T) [][]T {
	var res [][]T
	calPermutations(a, &res, 0)
	return res
}

func calPermutations[T any](arr []T, res *[][]T, k int) {
	for i := k; i < len(arr); i++ {
		arr[k], arr[i] = arr[i], arr[k]
		calPermutations(arr, res, k+1)
		arr[i], arr[k] = arr[k], arr[i]
	}
	if k == len(arr)-1 {
		*res = append(*res, slices.Clone(arr))
	}
}

func Map[E, V any](input []E, f func(E) V) []V {
	out := make([]V, 0, len(input))
	for _, e := range input {
		out = append(out, f(e))
	}
	return out
}

func Filter[E any](input []E, f func(E) bool) []E {
	out := make([]E, 0)
	for _, e := range input {
		if f(e) {
			out = append(out, e)
		}
	}
	return out
}
