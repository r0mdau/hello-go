# Hello-go

Learning golang.

## Notes
Unit Tests
```shell script
go test 
```

Code coverage
```shell script
go test -cover
```

Bench : use prefix `Benchmark`
```shell script
cd iteration/
#func BenchmarkRepeat
go test -bench=.
```

Race : test race condition
```shell script
go test -race
```

Writing example for godoc path `integers/adder_test.go`

Go linter
```bash
go get -u github.com/kisielk/errcheck
errcheck .
```

### Overview of folders

Nice exemple in `dependency_injection/` to test writers to stdout.

Goroutines and Channels are covered in `concurrency/`. Waiting for multiple channels is covered in `select/`.

Use a thread safe counter in `sync/`.

`context/` to help us manage long-running processes. From client cancellation to timeout, the goal is to quickly exit 
goroutines when handling http requests for example. More details : https://blog.golang.org/context