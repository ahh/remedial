package main

import "fmt"
import "vec"
import "sorter"
import "os"

func main() {
	stuff := vec.MakeVector(0, 0)

	sorter.Sort(stuff)
}
