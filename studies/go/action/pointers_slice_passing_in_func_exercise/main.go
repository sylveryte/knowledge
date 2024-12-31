package main

import "fmt"

func UpdateSlice(sl []string, s string) {
	sl[len(sl)-1] = s
}

func GrowSlice(sl []string, s string) {
	sl = append(sl, s)
}


func main() {
	sl := []string{"ram", "jam", "kam"}
	s := "kenny"
	fmt.Println("update before", sl)
  UpdateSlice(sl,s)
	fmt.Println("update after", sl)


  GrowSlice(sl,"magpie")
	fmt.Println("grow after", sl)
}
