package vec

type Vector struct {
	elements []int
}

func (v *Vector) At(i int) *int {
	return &v.elements[i]
}

func (v *Vector) Push_back(x int) {
	pos := len(v.elements)
	if len(v.elements) < cap(v.elements) {
		v.elements = v.elements[:(pos+1)]
	} else {
		v.elements = growslice(v.elements)
	}

	v.elements[pos] = x
}

func (v *Vector) Size() int {
	return len(v.elements)
}

func (v *Vector) Capacity() int {
	return cap(v.elements)
}

func (v *Vector) CopyFrom(rhs *Vector) {
	v.Resize(rhs.Size())
	for p := 0; p < v.Size(); p++ {
		v.elements[p] = rhs.elements[p]
	}
}

func (v *Vector) Resize(n int) {
	if n < v.Capacity()/2 {
		e := make([]int, n, n)
		for k := range e {
			e[k] = v.elements[k]
		}
		v.elements = e
	} else if n <= v.Capacity() {
		v.elements = v.elements[:n]
	} else {
		m := n
		if (m < v.Capacity()*2) {
			m = v.Capacity() * 2
		}
		e := make([]int, m, m)
		for k := range v.elements {
			e[k] = v.elements[k]
		}
		v.elements = e
	}
}

func (v *Vector) Empty() bool {
	return len(v.elements) == 0;
}

func (v *Vector) Pop_back() {
	l := len(v.elements)
	if l * 2 < cap(v.elements) {
		v.elements = shrinkslice(v.elements)
	} else {
		v.elements = v.elements[:(l-1)]
	}
}

func (v *Vector) Items() (retv chan int) {
	retv = make(chan int)
	go dump(retv, v.elements, len(v.elements))
	return
}


func MakeVector(size int, val int) (ret Vector) {
	ret = Vector{make([]int, size, size)}
	for i := 0; i < size; i++ {
		ret.elements[i] = val
	}
	return
}

func growslice(s []int) (ret []int) {
	c := cap(s)
	nc := 2 * c
	if nc < c+1 {
		nc = c + 1
	}
	ret = make([]int, c+1, nc)
	for k,v := range s {
		ret[k] = v
	}
	return
}

func shrinkslice(s []int) (ret []int) {
        c := cap(s)
	sz := len(s)
	ret = make([]int, sz-1, c/2)
	for k := range ret {
		ret[k] = s[k]
	}
	return
}

func dump(c chan int, elem [] int, count int) {
	for i := 0; i < count; i++ {
		c <- elem[i]
	}
	close(c)
}
