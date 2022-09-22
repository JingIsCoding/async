package async

type Result[T interface{}, E error] struct {
	value *T
	err   *E
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
