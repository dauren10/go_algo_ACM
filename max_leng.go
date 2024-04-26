package main

import "fmt"

func main() {
	nstring := "00101110000110"
	result := 0
	temp := 0
	for _, v := range nstring {

		if v == '0' {
			temp += 1
		} else {
			temp = 0
		}

		if result < temp {
			result = temp
		}
	}

	fmt.Println(result)
}
