package async

import "context"

type Futures[T any, E error] interface {
	Await() Results[T, E]
}

type futures[T any, E error] []Future[T, E]

func AsyncAll[T any, E error](ctx context.Context, funcs ...Func[T, E]) Futures[T, E] {
	fs := futures[T, E]{}
	for _, fun := range funcs {
		fs = append(fs, Async(ctx, fun))
	}
	return fs
}

func (futures futures[T, E]) Await() Results[T, E] {
	results := make(Results[T, E], len(futures))
	for index, future := range futures {
		results[index] = future.Await()
	}
	return results
}

var _ Futures[any, error] = &futures[any, error]{}
