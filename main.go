package main

import (
	"fmt"

	"github.com/csvwolf/Go-Leetcode/NumIslands"
)

//[["1","0","1","1","1"],["1","0","1","0","1"],["1","1","1","0","1"]]

func main() {
	fmt.Println(NumIslands.NumIslands([][]byte{
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '0', '1', '0', '1'},
		[]byte{'1', '1', '1', '0', '1'},
	}))
}
