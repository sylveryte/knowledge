package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main(){
  d := []byte("This is not a zip file")
  notAZipFile := bytes.NewReader(d)

  // zip.ErrFormat = errors.New("Just messing")
  fmt.Println(zip.ErrFormat)

  _, err := zip.NewReader(notAZipFile,int64(notAZipFile.Len()))
  fmt.Println(err)

  if err == zip.ErrFormat {
    fmt.Println("Told you",zip.ErrFormat)
  }
}
