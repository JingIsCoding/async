package async_await

type Result struct {
	value interface{}
	err   error
}

func (result Result) Value() interface{} {
	return result.value
}

func (result Result) Error() error {
	return result.err
}

func (result Result) IsOK() bool {
	return result.err == nil
}
