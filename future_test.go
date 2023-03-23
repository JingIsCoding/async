package async

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FutureTestSuite struct {
	suite.Suite
}

type testUserType struct {
	Name string
}

func (suite *FutureTestSuite) TestAwait() {
	suite.Run("should wait on future", func() {
		future := Async(context.Background(), func(res Resolve[string], rej Reject[error]) {
			time.Sleep(1000)
			res("ok")
		})
		result := future.Await()
		suite.Equal("ok", *result.Value())
	})

	suite.Run("should return custom value", func() {
		result := Async(context.Background(), func(res Resolve[testUserType], rej Reject[error]) {
			res(testUserType{Name: "John"})
		}).Await()
		suite.Equal("John", result.Value().Name)
	})

	suite.Run("should return value", func() {
		result := Async(context.Background(), func(res Resolve[string], rej Reject[error]) {
			res("yes")
		}).Await()
		suite.Equal("yes", *result.Value())
	})

	suite.Run("should return error", func() {
		err := errors.New("something is wrong")
		result := Async(context.Background(), func(res Resolve[interface{}], rej Reject[error]) {
			rej(err)
		}).Await()
		suite.Equal(&err, result.Error())
	})

	suite.Run("should error on context deadline exceeded", func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
		defer cancel()
		future := Async(ctx, func(res Resolve[string], rej Reject[error]) {
			time.Sleep(3 * time.Second)
			res("should not see this")
		})
		result := future.Await()
		suite.Equal("context deadline exceeded", (*result.Error()).Error())
	})

	suite.Run("should handle panic", func() {
		result := Async(context.Background(), func(res Resolve[interface{}], rej Reject[error]) {
			panic("something is deadly wrong..")
		}).Await()
		suite.False(result.IsOk())
		suite.Equal("something is deadly wrong..", (*result.Error()).Error())
	})
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestFutureTestSuite(t *testing.T) {
	suite.Run(t, new(FutureTestSuite))
}
