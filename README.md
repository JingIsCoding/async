## Async/Await

### An implementation that makes managing async tasks easier

#### From v1.0.1 using Golang generics to enforce typed result
#### Resolves the future with string type
```
result := Async(func(res Resolve[string], rej Reject[error]) {
	res("yes")
}).Await()
// do something with the result
fmt.Println(result.Value())
```
#### Resolves the future with custom type
```
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
```
result := Async(func(res Resolve[interface{}], rej Reject[error]) {
	rej(errors.New("something is wrong"))
}).Await()
fmt.Printf("%t", result.IsOK())
fmt.Printf("%e", result.Error())
// do something with the result that has error
```

#### With context
```
ctx, _ := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	future := Async(func(res Resolve[string], rej Reject[error]) {
	time.Sleep(3 * time.Second)
  res("should not see this")
}, ctx)
result := future.Await()
// Failed to resolve because the context timeouts
```
