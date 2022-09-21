package async_await

import "context"

type Resolve func(interface{})

type Reject func(error)

type Func func(Resolve, Reject)

type Future interface {
	Await() Result
}

type future struct {
	ctx        context.Context
	valChannel chan interface{}
	errChannel chan error
}

func (future *future) Await() Result {
	defer close(future.valChannel)
	defer close(future.errChannel)
	var value interface{}
	var err error
	select {
	case value = <-future.ctx.Done():
		return Result{
			value: nil,
			err:   future.ctx.Err(),
		}
	case value = <-future.valChannel:
		return Result{
			value: value,
			err:   nil,
		}
	case err = <-future.errChannel:
		return Result{
			value: nil,
			err:   err,
		}
	}
}

func Async(fun Func, ctxs ...context.Context) Future {
	ctx := context.Background()
	if len(ctxs) > 0 {
		ctx = ctxs[0]
	}
	future := &future{
		ctx:        ctx,
		valChannel: make(chan interface{}),
		errChannel: make(chan error),
	}
	go func() {
		fun(func(val interface{}) {
			future.valChannel <- val
		}, func(err error) {
			future.errChannel <- err
		})
	}()
	return future
}
