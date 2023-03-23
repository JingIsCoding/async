package async

import (
	"context"
	"errors"
	"fmt"
)

type Resolve[T any] func(T)

type Reject[E error] func(E)

type Func[T any, E error] func(Resolve[T], Reject[E])

type Future[T any, E error] interface {
	Await() Result[T, E]
}

type future[T any, E error] struct {
	ctx        context.Context
	valChannel chan T
	errChannel chan E
}

func Async[T any, E error](ctx context.Context, fun Func[T, E]) Future[T, E] {
	future := &future[T, E]{
		ctx:        ctx,
		valChannel: make(chan T),
		errChannel: make(chan E),
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(string); ok {
					future.errChannel <- errors.New(err).(E)
				}
			}
			close(future.valChannel)
			close(future.errChannel)
		}()
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
		var err E
		var ok bool
		if fe := future.ctx.Err(); fe != nil {
			if err, ok = fe.(E); !ok {
				err = fmt.Errorf("incompatible context errorr transfer from %T", fe).(E)
			}
		} else {
			err = errors.New("context cancelled").(E)
		}
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

var _ Future[any, error] = &future[any, error]{}
