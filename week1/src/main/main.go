package main

import "fmt"
import "vec"
import "sorter"
import "os"

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()
	v := vec.MakeVector(0, 0)
	for {
		var i int
	        n, _ := fmt.Fscanf(f, "%d,", &i)
		if n == 0 {
			n, _ = fmt.Fscanf(f, "%d", &i)
			if n == 0 {
				break
			}
		}
		// fmt.Printf("Got: %d\n", i)
		v.Push_back(i)
	}
	sorter.MSort(&v)
	last := -1
	for v := range v.Items() {
		if last > v {
			fmt.Printf("%d > %d ???\n", last, v)
		}
		last = v
		fmt.Printf("%d\n", v);
	}
}
