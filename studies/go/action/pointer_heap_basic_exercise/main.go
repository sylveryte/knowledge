package main

type Person struct {
  FirstName string
  LastName string
  Age int
}

func MakePerson(lastName string, firstName string, age int) *Person{
  return &Person{
    FirstName: firstName,
    LastName: lastName,
    Age: age,
  }
}

func main(){
  // pointer run with go build -gflags="-m" will show what escapes to heap
  p := MakePerson("tiwari", "Jaksan", 23)
  p.Age = 4

  var pl [10000000]*Person

  for i := range 1000 {
    pl[i] = MakePerson("tiwari", "Jaksan", 23)
  }


}
