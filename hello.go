package main

import (
	"fmt"
	//	"math"
)

var i int
var c, python, java bool

/*

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 の別名

rune // int32 の別名
     // Unicode のコードポイントを表す

	 float32 float64

	 complex64 complex128

*/

func main() {
	//	fmt.Println(add(10,15))
	a, b := swap("hoge", "fuga")
	fmt.Println(a, b)
	fmt.Println(split(10))
	fmt.Println(i, c, python, java)
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
