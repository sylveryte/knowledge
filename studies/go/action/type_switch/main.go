package main

import "fmt"

type MyInt int

func doThings(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Println("i is nil, type of j is any", j)
	case int:
		fmt.Println("type is int", j)
	case MyInt:
		fmt.Println("type is MyInt", j)
	case bool, rune:
		fmt.Println("type is bool or rune", j)
  default:
    fmt.Println("no idea bro")
	}
}

func main() {
  var k any;
  k = 12;
  doThings(k)
  k = true;
  doThings(k)
  k = nil;
  doThings(k)
  k = "somtth";
  doThings(k)
  k = MyInt(234);
  doThings(k)

}
