package main

import (
	"fmt"
)

func main() {
	var arr [3][4]int = [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}

	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v \t", i, j, v2)
		}
		fmt.Println()
	}
}
