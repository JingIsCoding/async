# Golang Async/Await

[![Go Report Card](https://goreportcard.com/badge/github.com/JingIsCoding/async)](https://goreportcard.com/report/github.com/JingIsCoding/async)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/JingIsCoding/async)](https://pkg.go.dev/github.com/JingIsCoding/async)

#### An implementation that makes managing async tasks easier

## Installation
```
go get -u github.com/JingIsCoding/async
```

## Usage

#### Resolves the future with string type
```go
future := Async(context.Background(), func(res Resolve[string], rej Reject[error]) {
	time.Sleep(1000)
	res("ok")
})
result := future.Await()
suite.Equal("ok", *result.Value())

```
#### Resolves the future with custom type
```go
type testUserType struct {
	Name string
}
result := Async(context.Background(), func(res Resolve[testUserType], rej Reject[error]) {
	res(testUserType{Name: "John"})
}).Await()
suite.Equal("John", result.Value().Name)

```

#### Reject the future
```go
err := errors.New("something is wrong")
result := Async(context.Background(), func(res Resolve[interface{}], rej Reject[error]) {
	rej(err)
}).Await()
suite.Equal(&err, result.Error())

```

#### Context time out
```go
ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
defer cancel()
future := Async(ctx, func(res Resolve[string], rej Reject[error]) {
	time.Sleep(3 * time.Second)
	res("should not see this")
})
result := future.Await()
suite.Equal("context deadline exceeded", (*result.Error()).Error())

```

## License

Distributed under MIT License. See [LICENSE](LICENSE) file for more details.
