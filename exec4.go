package main

import "fmt"

var _x int

func main() {
	fmt.Printf("%v %T\n", _x, _x)
	_x = 42
	fmt.Println(_x)
}
