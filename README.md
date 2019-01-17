# go-race

```
❯ go run -race main.go
value is 0
==================
WARNING: DATA RACE
Write at 0x00c000098008 by goroutine 6:
  main.main.func1()
      /Users/nakayama/go/src/github.com/bassaer/go-race/main.go:11 +0x4e

Previous read at 0x00c000098008 by main goroutine:
  main.main()
      /Users/nakayama/go/src/github.com/bassaer/go-race/main.go:14 +0x88

Goroutine 6 (running) created at:
  main.main()
      /Users/nakayama/go/src/github.com/bassaer/go-race/main.go:10 +0x7a
==================
Found 1 data race(s)
exit status 66
```
