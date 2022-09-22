package async

import "context"

type Resolve[T interface{}] func(T)

type Reject[E error] func(E)

type Func[T interface{}, E error] func(Resolve[T], Reject[E])

type Future[T interface{}, E error] interface {
	Await() Result[T, E]
}

type future[T interface{}, E error] struct {
	ctx        context.Context
	valChannel chan T
	errChannel chan E
}

func Async[T interface{}, E error](fun Func[T, E], ctxs ...context.Context) Future[T, E] {
	ctx := context.Background()
	if len(ctxs) > 0 {
		ctx = ctxs[0]
	}
	future := &future[T, E]{
		ctx:        ctx,
		valChannel: make(chan T),
		errChannel: make(chan E),
	}
	go func() {
		defer close(future.valChannel)
		defer close(future.errChannel)
		fun(func(val T) {
			future.valChannel <- val
		}, func(err E) {
			future.errChannel <- err
		})
	}()
	return future
}

func (future *future[T, E]) Await() Result[T, E] {
	var value T
	var err E
	select {
	case <-future.ctx.Done():
		err := future.ctx.Err().(E)
		return Result[T, E]{
			value: nil,
			err:   &err,
		}
	case value = <-future.valChannel:
		return Result[T, E]{
			value: &value,
			err:   nil,
		}
	case err = <-future.errChannel:
		return Result[T, E]{
			value: nil,
			err:   &err,
		}
	}
}
