## Async/Await

### An implementation that makes managing async tasks easier
#### Resolves the future
```
result := Async(func(res Resolve, rej Reject) {
  res("yes")
}).Await()
// do something with the result
fmt.Println(result.Value())
```

#### Reject the future
```
result := Async(func(res Resolve, rej Reject) {
  rej(errors.New("something is wrong"))
}).Await()
fmt.Printf("%t", result.IsOK())
fmt.Printf("%e", result.Error())
// do something with the result that has error
```


#### With context
```
ctx, _ := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
future := Async(func(res Resolve, rej Reject) {
  time.Sleep(3 * time.Second)
  res("time out on 3 seconds")
}, ctx)
result := future.Await()
// Failed to resolve because the context timeouts
```
