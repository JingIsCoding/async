package async

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ResultTestSuite struct {
	suite.Suite
}

type valueType struct {
	data string
}

func (suite *ResultTestSuite) TestResult() {
	suite.Run("should return string value", func() {
		result := OkResult("ok")
		suite.Equal("ok", result.Value())
	})

	suite.Run("should return value of custom type", func() {
		result := OkResult(valueType{data: "yes"})
		suite.Equal("yes", result.Value().data)
	})

	suite.Run("should return error", func() {
		result := ErrorResult(errors.New("not ok"))
		suite.Equal("not ok", result.Error().Error())
	})

	suite.Run("should return is not ok there is error", func() {
		result := ErrorResult(errors.New("not ok"))
		suite.False(result.IsOK())
	})
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestResultTestSuite(t *testing.T) {
	suite.Run(t, new(ResultTestSuite))
}
