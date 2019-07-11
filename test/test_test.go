package main

import (
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"strconv"
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	a := getValue(1)
	b := getValue(0)
	//assert.Equal(t, 1, "1")
	fmt.Println(a == b)

	arr1 := []int{1, 3}
	arr2 := []int{1, 3}
	result := assert.IsEqual(arr1, arr2)

	fmt.Println("数组比较", result)


	strS := strings.Replace("1%32 ", "%", "\\%", -1)
	fmt.Println(strS)

	strS = "   123   "
	strS = strconv.Quote(strS)
	fmt.Println()

	fmt.Println("len nil:")




}

func getValue(arguments ...interface{}) interface{} {
	return arguments[0]
}
