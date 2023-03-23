package async

import (
	"context"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FuturesTestSuite struct {
	suite.Suite
}

func (suite *FuturesTestSuite) TestAwait() {
	suite.Run("should wait on future", func() {
		future := AsyncAll(context.Background(), func(res Resolve[string], rej Reject[error]) {
			time.Sleep(1000)
			res("ok")
		})
		results := future.Await()
		suite.Equal(1, len(results.Value()))
		suite.Equal("ok", *results.Value()[0])
	})

	suite.Run("should return not ok if one of the function fail", func() {
		future := AsyncAll(context.Background(), func(res Resolve[string], rej Reject[error]) {
			time.Sleep(1000)
			res("ok")
		}, func(res Resolve[string], rej Reject[error]) {
			rej(errors.New("not ok"))
		})
		results := future.Await()
		suite.False(results.IsOk())
		suite.Equal(2, len(results.Value()))
		suite.Equal("ok", *results.Value()[0])
		suite.Equal(errors.New("not ok"), *results.Error()[1])
	})

	suite.Run("should get all results even if some calls takes longers", func() {
		future := AsyncAll(context.Background(),
			func(res Resolve[string], rej Reject[error]) {
				time.Sleep(2000)
				res("ok1")
			},
			func(res Resolve[string], rej Reject[error]) {
				time.Sleep(1000)
				res("ok2")
			})
		results := future.Await()
		suite.True(results.IsOk())
		suite.Equal(2, len(results.Value()))
		suite.Equal("ok1", *results.Value()[0])
		suite.Equal("ok2", *results.Value()[1])
	})

	suite.Run("should get all job results", func() {
		jobs := []Func[int, error]{}
		for i := 0; i < 100; i++ {
			j := i
			jobs = append(jobs, func(res Resolve[int], rej Reject[error]) {
				time.Sleep(time.Duration(rand.Int63n(100)))
				res(j)
			})
		}
		future := AsyncAll(context.Background(), jobs...)
		results := future.Await()
		suite.True(results.IsOk())
		suite.Equal(100, len(results.Value()))
		for index, value := range results.Value() {
			suite.Equal(index, *value)
		}
	})
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestFuturesTestSuite(t *testing.T) {
	suite.Run(t, new(FuturesTestSuite))
}
