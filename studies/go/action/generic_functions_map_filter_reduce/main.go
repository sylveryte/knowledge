package main

import (
	"fmt"
	"strings"
)

func Map[t1, t2 any](s []t1, f func(t1) t2) []t2 {
	r := make([]t2, len(s))

	for i, v := range s {
		r[i] = f(v)
	}

	return r
}

func Filter[t any](s []t, f func(a t) bool) []t {
	r := []t{}
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[t1, t2 any](s []t1, f func(newValue t1, reducedValue t2) t2, initial t2) t2 {
	r := initial
	for _, v := range s {
		r = f(v, r)
	}

	return r
}

func main() {
	s := []string{"sola", "kattu", "kola", "puttu", "mola", "jola"}
	fmt.Println("s", s)
	// s [sola kola mola jola]
	j := Map(s, func(a string) int { return len(a) })
	fmt.Println("j", j)
	// j [4 5 4 5 4 4]
	fl := Filter(s, func(a string) bool {
		return strings.HasSuffix(a, "ola")
	})
	fmt.Println("fl", fl)
	// fl [sola kola mola jola]
	rd := Reduce(s, func(a string, b string) string { return a + b }, "")
	fmt.Println("rd", rd)
	// rd jolamolaputtukolakattusola
}
