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

Writing example for godoc path `integers/adder_test.go`

Go linter
```bash
go get -u github.com/kisielk/errcheck
errcheck .
```

Nice exemple in dependency injection to test writers to stdout.