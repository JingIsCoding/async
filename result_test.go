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
		value := "ok"
		result := Result[string, error]{
			value: &value,
			err:   nil,
		}
		suite.Equal("ok", result.Value())
	})

	suite.Run("should return value of custom type", func() {
		value := valueType{data: "yes"}
		result := Result[valueType, error]{
			value: &value,
			err:   nil,
		}
		suite.Equal("yes", result.Value().data)
	})

	suite.Run("should return error", func() {
		err := errors.New("not ok")
		result := Result[interface{}, error]{
			value: nil,
			err:   &err,
		}
		suite.Equal("not ok", result.Error().Error())
	})

	suite.Run("should return is not ok there is error", func() {
		err := errors.New("not ok")
		result := Result[interface{}, error]{
			value: nil,
			err:   &err,
		}
		suite.False(result.IsOK())
	})
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestResultTestSuite(t *testing.T) {
	suite.Run(t, new(ResultTestSuite))
}
