package main

import (
	"fmt"
	"math"
)

type Pair[T any] struct {
	val1 T
	val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

// this [T Differ[T]] function makes any value to have Differ implemented
// plus the arguments makes sure you're only sending the Pair
// so we only use this func for a Pair with Differ implemented
func FindCloser[T Differ[T]](p1, p2 Pair[T]) Pair[T] {
	d1 := p1.val1.Diff(p1.val2)
	d2 := p2.val1.Diff(p2.val2)

	if d1 < d2 {
		return p1
	} else {
		return p2
	}
}

type twoD struct {
	a int
	b int
}

func (p twoD) Diff(v twoD) float64 {
	x := p.b - p.a
	y := v.b - v.a
	return math.Sqrt(float64(x) - float64(y))
}

func (p twoD) String() string {
	return fmt.Sprintf("{%d,%d}", p.a, p.a)
}


func main() {

	j := Pair[twoD]{
		twoD{1, 4},
		twoD{3, 9},
	}
	k := Pair[twoD]{
		twoD{4, 8},
		twoD{8, 17},
	}

  closer:=FindCloser(j,k)

	fmt.Println("closer", closer)
}
