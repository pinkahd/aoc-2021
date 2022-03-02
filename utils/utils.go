package utils

import (
	"fmt"
	"os"
	"strconv"
)

func MakeIntArr(s []string) []int {
	arr := make([]int, len(s))
	for i, e := range s {
		arr[i], _ = strconv.Atoi(e)
	}
	return arr
}

func Bail(any interface{}) {
	fmt.Printf("Bailing out 👋\n%v\n", any)
	os.Exit(1)
}

func ReadFile(f string) string {
	data, err := os.ReadFile(f)
	if err != nil {
		Bail(err)
	}
	return string(data)
}

func Reduce(acc interface{}, arr []interface{}, fn func(interface{}, interface{}) interface{}) interface{} {
	val := acc
	for _, e := range arr {
		val = fn(val, e)
	}
	return val
}

func Sum(acc interface{}, val interface{}) interface{} {
	return acc.(int) + val.(int)
}

func Itoifa(t []int) []interface{} {
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}
	return s
}
