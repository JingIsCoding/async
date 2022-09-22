package async

type Result[T interface{}, E error] struct {
	value *T
	err   *E
}

func OkResult[T interface{}](value T) Result[T, error] {
	return Result[T, error]{
		value: &value,
		err:   nil,
	}
}

func ErrorResult[E error](err E) Result[interface{}, E] {
	return Result[interface{}, E]{
		value: nil,
		err:   &err,
	}
}

func (result Result[T, E]) Value() T {
	return *result.value
}

func (result Result[T, E]) Error() E {
	return *result.err
}

func (result Result[T, E]) IsOK() bool {
	return result.err == nil
}
