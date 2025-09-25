# Go Tests

#golang

## Basics Rules

- Tesing file name needs to be `xxxx_test.go`
- Test function must start with `TestSomeFunction`
- Test function must only accept one argument only `t *testing.T`
- In order to use `t` argument you must import `testing`

## Basic Test

- Prepare `want` and `got` however you want
- write `if want!=got` return `t.Errorf("standard string thing")`

```go
func TestAnalyzeCompletionTrigger(t *testing.T) {
	line := "Cool new stuff"
	kind, arg := AnalyzeTriggerKind(0, line)
	if kind != CompletionNone {
		t.Errorf(fmt.Sprintf("Wanted %d got %d", CompletionNone, kind))
	}
	if arg != "" {
		t.Errorf(fmt.Sprintf("Wanted blank got %s", arg))
	}
}
```

### Subtests

- To easily group some tests

## Running test

- run with `go test`
- to run recursively `go test -v ./...`
- or to run a specific dir `go test ./utils` utils dir
- `go test ./utils/..` to run utils dir recursively

## Tip watch n test

use [[Run on files changed]]

---

[[Learn Go With Tests]]
