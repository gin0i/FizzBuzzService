package managers

import (
	"FizzBuzzService/models/request"
	"github.com/pkg/errors"
)

func BiggestRequest() (*request.Request, error) {
	found := request.FindWithHighestCount()
	if found == nil {
		return nil, errors.New("No request records")
	}
	return found, nil
}