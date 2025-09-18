# Initialize project/module

- `go mod init example/hello`
- create `hello.go`

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

## Make

- best way is to use Makefile

  ```make
  .DEFAULT_GOAL := build

  .PHONY:fmt vet build run
  fmt:
  go fmt
  vet: fmt
  go vet
  build: vet
  go build
  run: build
  ./hello #main.go
  ```

- run using command `make run`

---

[[Learning Go]]
