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
fmt.Println("should be ok ", result.IsOk())
fmt.Println(*result.Value())

```
#### Resolves the future with custom type
```go
type testUserType struct {
	Name string
}
result := Async(context.Background(), func(res Resolve[testUserType], rej Reject[error]) {
	res(testUserType{Name: "John"})
}).Await()
fmt.Println(result.Value().Name)

```

#### Reject the future
```go
err := errors.New("something is wrong")
result := Async(context.Background(), func(res Resolve[interface{}], rej Reject[error]) {
	rej(err)
}).Await()
fmt.Println((*result.Error()).Error())

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
fmt.Println((*result.Error()).Error())
```

#### Trigger multiple async jobs
```go
future := AsyncAll(context.Background(),
	func(res Resolve[string], rej Reject[error]) {
		time.Sleep(2000)
		res("ok1")
	},
	func(res Resolve[string], rej Reject[error]) {
		time.Sleep(1000)
		res("ok2")
})
results := future.Await()
fmt.Println(result.Value()[0])
fmt.Println(result.Value()[1])

```

## License

Distributed under MIT License. See [LICENSE](LICENSE) file for more details.
