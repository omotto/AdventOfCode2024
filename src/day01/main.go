package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hola")
	i, err := strconv.Atoi("243")
	s := strings.ToUpper("fwwre")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
	fmt.Println(s)
}
