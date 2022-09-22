# Golang Async/Await

[![Go Report Card](https://goreportcard.com/badge/github.com/JingIsCoding/async)](https://goreportcard.com/report/github.com/JingIsCoding/async)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/JingIsCoding/async)](https://pkg.go.dev/github.com/JingIsCoding/async)

#### An implementation that makes managing async tasks easier

## Installation
```
go get -u github.com/JingIsCoding/async
```

## Usage
#### From v1.0.1 using Golang generics to enforce typed result

#### Resolves the future with string type
```go
future := Async(func(res Resolve[string], rej Reject[error]) {
  // do something asynchronously 
  res("yes")
})
result := future.Await()
// do something with the result
fmt.Println(result.Value())
```
#### Resolves the future with custom type
```go
type User struct {
  Name string
}
// somewhere in the code
result := Async(func(res Resolve[User], rej Reject[error]) {
  res(User{Name:":"some one"})
}).Await()
// do something with the result
fmt.Println(result.Value().Name)
```

#### Reject the future
```go
result := Async(func(res Resolve[interface{}], rej Reject[error]) {
	rej(errors.New("something is wrong"))
}).Await()
fmt.Printf("%t", result.IsOK())
fmt.Printf("%e", result.Error())
// do something with the result that has error
```

#### With context
```go
ctx, _ := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	future := Async(func(res Resolve[string], rej Reject[error]) {
	time.Sleep(3 * time.Second)
  res("should not see this")
}, ctx)
result := future.Await()
// Failed to resolve because the context timeouts
```

## License

Distributed under MIT License. See [LICENSE](LICENSE) file for more details.
