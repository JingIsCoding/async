package async

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ResultTestSuite struct {
	suite.Suite
}

func (suite *ResultTestSuite) TestResult() {
	suite.Run("should return value", func() {
		result := Result{
			value: "ok",
		}
		suite.Equal("ok", result.Value())
	})

	suite.Run("should return error", func() {
		result := Result{
			err: errors.New("not ok"),
		}
		suite.Equal("not ok", result.Error().Error())
	})

	suite.Run("should return is not ok there is error", func() {
		result := Result{
			err: errors.New("not ok"),
		}
		suite.False(result.IsOK())
	})
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestResultTestSuite(t *testing.T) {
	suite.Run(t, new(ResultTestSuite))
}
