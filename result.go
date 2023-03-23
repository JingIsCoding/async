package async

type Result[T any, E error] struct {
	value *T
	err   *E
}

func (result Result[T, E]) Value() *T {
	return result.value
}

func (result Result[T, E]) Error() *E {
	return result.err
}

func (result Result[T, E]) IsOk() bool {
	return result.err == nil
}

type Results[T any, E error] []Result[T, E]

func (results Results[T, E]) Value() []*T {
	values := []*T{}
	for _, result := range results {
		values = append(values, result.value)
	}
	return values
}

func (results Results[T, E]) Error() []*E {
	errors := []*E{}
	for _, result := range results {
		errors = append(errors, result.err)
	}
	return errors
}

func (results Results[T, E]) IsOk() bool {
	for _, result := range results {
		if !result.IsOk() {
			return false
		}
	}
	return true
}
