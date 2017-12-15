package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// fmt.Println("Hello, World!")
	data := strings.SplitN("hello:1511786641124", ":", 2)
	fmt.Println(len(data))
	num, err := strconv.ParseInt(strings.TrimSpace(data[1]), 10, 64)
	if err != nil {
		fmt.Println("eeeeeeeeeeeee")
	} else {
		fmt.Println(num)
		fmt.Println(time.Unix(0, num*1000*1000))

	}

}
