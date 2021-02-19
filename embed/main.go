package main

import (
	_ "embed"
)

//go:embed test.html
var global string

func main() {
	println(global)
}
