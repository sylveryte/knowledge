package main

import "fmt"

func LogOutput(message string) {
	fmt.Println(message)
}

type LogAdapter func(ms string)

func main(){

  k := LogAdapter(LogOutput)
  k("okay man")
}
