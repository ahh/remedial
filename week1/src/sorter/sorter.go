package sorter

import "vec"
import "fmt"

func QSort(v *vec.Vector) {
	c := make(chan int)
	qhelp(v, 0, v.Size() - 1, c)
//	<- c
}
var debug bool = false

func swap(a, b *int) {
	*a, *b = *b, *a
}

func qhelp(v *vec.Vector, b, e int, c chan int) {
	if debug {
		fmt.Printf("qhelp(%d, %d)\n", b, e)
	}

	if (e - b) < 1 {
		return
	}
//	fmt.Printf("qhelp: %d %d\n", b, e)
	// lazy, so just pivot on first element
	pivot := *v.At(b)
	i := b + 1
	j := e
	for i < j {
		if (*v.At(i) <= pivot) {
			i++
		} else if (*v.At(j) > pivot) {
			j--
		} else {
			swap(v.At(i), v.At(j))
		}
	}
	// everything < (i == j) is <= pivot, everything > is >pivot
	if (*v.At(i) > pivot) {
		i--
	} else {
		j++
	}
	//now [b,i] is <= pivot, and
	swap(v.At(b), v.At(i))

	if (debug) {
		for p := b; p <= i; p++ {
			if *v.At(p) > pivot {
				fmt.Printf("%d(%d) > %d\n", p, *v.At(p), pivot)
			}
		}
		if i + 1 != j {
			fmt.Printf("%d +1 != %d\n", i, j)
		}
		if *v.At(i) != pivot {
			fmt.Printf("%d(%d) != %d\n", i, *v.At(i), pivot)
		}
		for p := j; p <= e; p++ {
			if *v.At(p) <= pivot {
				fmt.Printf("%d(%d) <= %d\n", p, *v.At(p), pivot)
			}
		}

		fmt.Printf("%d %d\n", i, j)
		fmt.Printf("[\n")
		for p := b; p <= e; p++ {
			if (i == p || j == p) {
				fmt.Printf("--- %d\n", pivot)
			}
			fmt.Printf("%d\n", *v.At(p))
		}
		fmt.Printf("]\n")
	}
	c1 := make(chan int)
	c2 := make(chan int)
	qhelp(v, b, i-1, c1)
	qhelp(v, j , e, c2)
//	<- c1
//	<- c2
	//c <- 1
}
