package main

import (
	"fmt"
	"strconv"
)

func main(){
	a, _ := strconv.ParseBool("t")
	fmt.Println(a)


	name := "lwx"
	str := "%"+name+"%"
	outputStr := fmt.Sprintf("%s" ,str)
	fmt.Println(outputStr)
}