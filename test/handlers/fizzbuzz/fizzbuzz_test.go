package capacites

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"FizzBuzzService/handlers"
	"FizzBuzzService/models/request"
	"FizzBuzzService/routers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	externalRouter = routers.NewRouter()
)

type TestSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupTest() {
	request.DeleteAll()
}

func (suite *TestSuite) TearDownTest() {
	request.DeleteAll()
}

func (suite *TestSuite) TestBasicFizzBuzz() {
	url := "/fizzbuzz/api/launch"
	params := handlers.RawFizzBuzzRequest{
		MulA:  3,
		MulB:  5,
		Limit: 100,
		StrA:  "Fizz",
		StrB:  "Buzz",
	}
	payload, err := json.Marshal(params)


	assert.Nil(suite.T(), err)
	jsonStr := []byte(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	externalRouter.ServeHTTP(resp, req)
	fmt.Println("RESP is", resp.Code, resp.Body.String())
	assert.Equal(suite.T(), http.StatusOK, resp.Code, resp.Body.String())

	var resultContent []string
	err = json.NewDecoder(resp.Body).Decode(&resultContent)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "76", resultContent[75])
	assert.Equal(suite.T(), "Fizz", resultContent[8])
	assert.Equal(suite.T(), "FizzBuzz", resultContent[14])
	assert.Equal(suite.T(), "Buzz", resultContent[49])
}

func (suite *TestSuite) TestBigFizzBuzz() {
	url := "/fizzbuzz/api/launch"
	params := handlers.RawFizzBuzzRequest{
		MulA:  3,
		MulB:  5,
		Limit: 10000,
		StrA:  "???",
		StrB:  "!!!!!!",
	}
	payload, err := json.Marshal(params)
	assert.Nil(suite.T(), err)

	jsonStr := []byte(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	externalRouter.ServeHTTP(resp, req)
	fmt.Println("RESP is", resp.Code, resp.Body.String())
	assert.Equal(suite.T(), http.StatusOK, resp.Code, resp.Body.String())

	var resultContent []string
	err = json.NewDecoder(resp.Body).Decode(&resultContent)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "536", resultContent[535])
	assert.Equal(suite.T(), "???", resultContent[8])
	assert.Equal(suite.T(), "???!!!!!!", resultContent[8414])
	assert.Equal(suite.T(), "!!!!!!", resultContent[7999])

}

func (suite *TestSuite) TestBadLimitsFizzBuzz() {
	url := "/fizzbuzz/api/launch"
	params := handlers.RawFizzBuzzRequest{
		MulA:  3,
		MulB:  5,
		Limit: -100,
		StrA:  "Fizz",
		StrB:  "Buzz",
	}
	payload, err := json.Marshal(params)
	assert.Nil(suite.T(), err)

	jsonStr := []byte(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	externalRouter.ServeHTTP(resp, req)
	log.Println("RESP is", resp.Code, resp.Body.String())
	assert.Equal(suite.T(), http.StatusBadRequest, resp.Code, resp.Body.String())
}

func (suite *TestSuite) TestBadMultipleFizzBuzz() {
	url := "/fizzbuzz/api/launch"
	resp := httptest.NewRecorder()


	// Test with A 0
	params := handlers.RawFizzBuzzRequest{
		MulA:  0,
		MulB:  10,
		Limit: 100,
		StrA:  "Fizz",
		StrB:  "Buzz",
	}
	payload, err := json.Marshal(params)

	assert.Nil(suite.T(), err)
	jsonStr := []byte(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")


	// Same but with B Negative
	externalRouter.ServeHTTP(resp, req)
	log.Println("RESP is", resp.Code, resp.Body.String())
	assert.Equal(suite.T(), http.StatusBadRequest, resp.Code, resp.Body.String())

	params.MulA = 0
	params.MulB = -3
	payload, err = json.Marshal(params)

	assert.Nil(suite.T(), err)
	jsonStr = []byte(payload)

	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")


	externalRouter.ServeHTTP(resp, req)
	log.Println("RESP is", resp.Code, resp.Body.String())
	assert.Equal(suite.T(), http.StatusBadRequest, resp.Code, resp.Body.String())
}


