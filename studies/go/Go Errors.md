# Errors

#golang

## Basic

- Go expects the developer to handle the Errors
- Idiomatic way is to return `error` as last in tuple for
  return

```go
value, err := someErrorFunc()
if err == nil {
  // there is error
}
```

- error is created by two ways
  - first is by creating new using `errors` pkg
  - second is using `fmt.Errorf` allows to add runtime info

```go
err := errors.New("denominator is zero")
err2 := fmt.Errorf("%d isn't even number",a)
```

- 🍏 Error message should be in all smallcase and does
  _not_ contain any punctuations or newline (.! etc)
- `error` is a builtin interface that defins only 1 func

```go
   type error interface {
     Error() string
   }
```

- anything that implments this `interface` is considered error

## Sentinel Errors

- These are one of few variables declared at pkg level
- They start with `Err` generally
- `zip.ErrFormat` is a Sentinel error
- They should be treated as `readonly`
- Used to indicate that you cannot start or continue
  processing
- You compare these errors with your returned error
  {/ ./action/errors_sentinel/main.go}
- 🍎 be aware once you add a Sentinel error, it is forever
  in your public API and you have to support it in
  future

## Custom Error

- Anthing that implments the error `interface` [[Go Errors#Basic]]
- `error` is an interface, so the error is a value with the
  interface
- You can add more data on error by creating a custom

```go
type Status int
type StatusErr struct {
  Status  Status
  Message string
}
func (s StatusErr) Error() string {
  return s.Message
}
```

[errors_custom](./action/errors_custom/main.go)

- We can use StatusErr to add more details
- 🏮Be sure _not_ to return uninitialized error since,
  for an instance to be nil both the value and interface has to be nil
  using `var statusError` instead use `return StatusErr{..`

## Wrapping Errors

- `fmt.ErrorF` has a special verb `%w`
  `wrappedErr := fmt.Errorf("Error this is %w", pErr)`
- use this to create a format string which includes
  formatted string of another error and contains the
  original error as well
- stdlib provides a func to /unwrap/ `errors.Unwrap`
  `unwappedErr := errors.Unwrap(wrappedErr)`
  [errors_wrap_unwrap](./action/errors_wrap_unwrap/main.go)
- 📘you don't directly call /unwrap/ but use {todo Is and As}
- {todo Custom Error} can be wrapped with method `Unwrap` implmented
  [errors_custom](./action/errors_custom/main.go)
- To create a new error that contains the message from
  another error but don't want to wrap, use `fmt.Errorf` to
  create error but use %v instead of %w

### Wrapping Multiple Errors

    - 1. `errors.Join`

```go
var errs []error
errs = append(errs, errors.New("new err"))
return errors.Join(errs...)
```

    - 2. `%w %w`

```go
errs := fmt.Errorf("1 %w, 2 %w",er1,er2)
```

    - 3. {todo Custom Error}[Custom]

## Is and As

### Is

- To check whether the returned error or any errors that
  it wraps match a specific {\*\* Sentinel Errors}[sentinel error] instance, use
  `errors.Is` ( `error`, `instance you're comparing it against`)
  `if errors.Is(err, os.ErrNotExist)`
- By default `Is` uses == to compare each wraped with the
  specific error
- if your {todo Custom Error} is a noncomparable then implement
  [errors_custom](./action/errors_custom/main.go)

```go
func (se StatusErr) Is(target error) bool {
  if se2, ok := target.(StatusErr); ok {
    return se.Status == se2.Status
  }
  return false
}
```

### As

    - To check whether a returned error(or any error it wraps)
      matches a specific type
      `errors.As`( `error`, `pointer to var of the type` OR `pointer to an interface`)

```go
var se StatusErr
errors.As(err, &se)
var coder interface {
  ...
}
errors.As(err, &coder)
```

    - 🙀 if second parameter is anything other that
      pointer to error or interface method panics

## Wrapping errors with defer

- When you have to wrap multiple errors with same msg
  (Check book) 🙃
- defer closure

## Panic and Recover

- 🚙 As soon as panic happens, the current function _exits_ immediately,
  and any _defers_ attached to the current function start running
  and then calling function and so on till `main` is reached
- Goroutine the chain of defers ends at the function used
  to launch the goroutine
- The builtin `recover` function is called from within the
  defer to check whether a panic happened
- explict panic => `panic(msg)`
- There's a /specific/ pattern for using `recover`. Register
  a function with defer to handle a potential panic.
- You must call `recover` from within the defer because
  only defers are ran after panic
- 󱅻 You'll rarely want to keep your program running
  after a panic occurs
- The reason we don't rely on panic & recover is the
  recover doesn't make it clear what could fail. It just
  ensures that if something fails, you can print out error
  and continue
- 🙅 Go doesn't provide /stack trace/

---

[[Learning Go]]
