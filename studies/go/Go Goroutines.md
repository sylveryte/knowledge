# Go Goroutines Concurrency

- based on Communcating Sequential Processes (CSP)

## To use or not to use concurrency

### How to not use concurrency 🤡

- This is amazing, I'm going to put Goroutines everywhere
- My program isn't faster. I'm adding buffers to my channels
- My channels are blocking and I'm getting deadlocks.
  I'm going to use buffered channels with /really/ big buffers
- My channels are still blocking. I'm going to use mutexes
- I'm giving up

### How to use it?

- The key is understand that /concurrency is not parallelism/ ⭐
- Concurrency is to better structure the your problem
  Whether concurrency runs in parallel deps on H/W and algo
- Pro Tip: 👩 /Amdahl's law/ is formula for figuring out
  how much parallel processing can improve perf
- Use concurrency when you want to combine data from
  multiple operations that can operate independently
- 🍊 Isn't worth it when process doesn't take long time,
  since many in-memory algos are so fast that overhead
  of passing value via concurrency results negative
- 🫐 Concurrent ops are often used for I/O, network
  which are 1000x slower than all but the most complicated
  in-memory processes

## Goroutines

### What is it?

#### Process

- A _Process_ is an instance of program being run by os.
- Os associates resources exclusively like memory.

#### Threads

- A process is composed of one or more _Threads_.
- A thread is a unit of execution that is given some
  time by os to run.
- Threads within a process share access to resources.
- A CPU can execute more than one thread deps on cores.

#### Goroutines

- Think of _Goroutine_ as a /lightweight thread/.
- Go program starts, go runtime creates a number of threads
  and launches _a single goroutine_ to run your program.
- All the Goroutines created by your program including
  intial ones, are assigned to these threads
  automagically by go runtime /scheduler/ just like os
  schedules threads across CPU cores.
- _Are goroutines the just extra work over threads?_
  - _No_, becuase
  * Goroutine is created faster, since no os level res
    - It's intial stack size is smaller
    - Switiching between it is faster, since within process
  - Goroutine scheduler is context aware does better scheduling
- Allows _hundreds, thousands even 10x thousands_ of
  _simultaneous goroutines_. Threads can barely.

#### How to use?

- It can be launched by placing `go` before func invocation.
  `go someFunc();`
- 🦧 returns are _ignored_.
- It is customary in Go to launch goroutines with a
  closure that wraps business logic. The closure takes
  care of the concurrent bookkeeping.
  {/ ./action/goroutine_basic/main.go}

## Channels

- Like maps channels are ref types.
- Passing in func is passed pointer, like slice and maps
- [zv] is nil

### Reading Writing and Buffering

```go
a := <- ch //reads value from ch to a, notice <-
ch <- b // writes value in b to ch
```

- _Each value written_ to a channel can be _read only once_
- If multiple goroutines reading only one will read
- A single goroutine /rarely/ reads or writes to same channel 📙
- When assigning a channel to a variable or field, or
  passing it to a func, Use an arrow before the `chan`
  keyword `(ch <-chan int)` to indicate that goroutine only
  /reads/ from the channel.
- Use an arrow after the chan keyword `(ch chan<- int)` to
  indicate only writes to the channel
- By _default_ channels are _unbuffered_
- 🚙 When written, the writing gr pauses till ch iz read,
  likewise when read, reading gr pauses till ch iz written
- Also has _buffered_ channels with a number limited
  `ch := make(chan int,10)`
  - `len` and `cap` provide info, len=filled numb, cap=cap
  - limited number of writes without blocking
  - 🚙 when all filled and one more write,
    then writing gr pauses,
    same reading empty buffer ch blocks gr
- 💡Most of the times u should use _unbuffered_ channels

### Using for-range

```go
for v := range ch {
  fmt.Println(v)
}
```

- There is only one variable(value).
- Assigned to v when channel is open and value is there,
  and _body is executed_.
- If no value gr is paused, until a value or ch is closed
- Loop contnues until the break,return or ch is closed

### Closing a Channel

- `close(ch)`
- 🏮Once closed any attempt to _write is panic_
- _Read_ from a closed /always succeeds/
- If no value the [zv] of channel's type is returned
  - 💡use {:$/studies/go/DataTypes:\* comma ok idiom}[comma ok idiom] to check if channel is closed
- _Close_ channel req only when gr is waiting for closed
  eg is `for-range` loop
- make _writing_ goroutine to _close_ the channel 🌷
- use a `sync.WaitGroup` to address multiple grs closing

### Behavior

|       | unbug open          | unbuf closed | buf open             | buf closed                |
| ----- | ------------------- | ------------ | -------------------- | ------------------------- |
| read  | pause until written | ret zv (ok,) | pause if empty       | ret remain, then zv (ok,) |
| write | pause until read    | PANIC 🚗     | pause if full        | PANIC 🚗                  |
| close | works               | PANIC 🚗     | works, rem stl there | PANIC 🚗                  |

### Select

- You can't favor one gr ops over another else,
  you'll never process some cases & /starvation/ occurs
- The `select` keyword allows a goroutine to read from or
  write to one of a set of multiple channels.
- Each case in select is /read/ or /write/ to a channel

```go
select {
case v := <-ch:
  fmt.Println(v)
case v := <-ch2:
  fmt.Println(v)
case ch3 <- x:
  fmt.Println("wrote ",x)
case <-ch4
  fmt.Println("Ignored value on ch4")
}
```

---

[[Learning Go]]
