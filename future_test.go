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
		future := Async(func(res Resolve[string], rej Reject[error]) {
			res("ok")
		})
		time.Sleep(1000)
		result := future.Await()
		suite.Equal("ok", result.Value())
	})

	suite.Run("should return custom value", func() {
		result := Async(func(res Resolve[testUserType], rej Reject[error]) {
			res(testUserType{Name: "John"})
		}).Await()
		suite.Equal("John", result.Value().Name)
	})

	suite.Run("should return value", func() {
		result := Async(func(res Resolve[string], rej Reject[error]) {
			res("yes")
		}).Await()
		suite.Equal("yes", result.Value())
	})

	suite.Run("should return error", func() {
		result := Async(func(res Resolve[interface{}], rej Reject[error]) {
			rej(errors.New("something is wrong"))
		}).Await()
		suite.Equal("something is wrong", result.Error().Error())
	})

	suite.Run("should error on context deadline exceeded", func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
		defer cancel()
		future := Async(func(res Resolve[string], rej Reject[error]) {
			time.Sleep(3 * time.Second)
			res("should not see this")
		}, ctx)
		result := future.Await()
		suite.Equal("context deadline exceeded", result.Error().Error())
	})

	suite.Run("should handle panic", func() {
		result := Async(func(res Resolve[interface{}], rej Reject[error]) {
			panic("something is deadly wrong..")
		}).Await()
		suite.False(result.IsOK())
		suite.Equal("something is deadly wrong..", result.Error().Error())
	})
}

func (suite *FutureTestSuite) TestPanic() {
	result := Async(func(res Resolve[interface{}], rej Reject[error]) {
		panic("something is deadly wrong..")
	}).Await()
	suite.False(result.IsOK())
	suite.Equal("something is deadly wrong..", result.Error().Error())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestFutureTestSuite(t *testing.T) {
	suite.Run(t, new(FutureTestSuite))
}
