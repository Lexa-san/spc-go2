package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	InputData int
	Answer    int
	Expected  int
}

var Cases []TestCase = []TestCase{
	{
		InputData: 0,
		Expected:  1,
	},
	{
		InputData: 1,
		Expected:  1,
	},
	{
		InputData: 3,
		Expected:  6,
	},
	{
		InputData: 5,
		Expected:  120,
	},
}

func TestFactorial(t *testing.T) {
	for id, test := range Cases {
		if test.Answer = factorial(test.InputData); test.Answer != test.Expected {
			t.Errorf("test case %d failed: expected %d but result %d", id, test.Expected, test.Answer)
		}
	}
}

type HTTPTestCase struct {
	Name     string //test name
	Numeric  int    //value from HTTP-request
	Expected []byte //HTTP-response expected
}

// test case for HTTP request
// each request must have specific case
var HTTPCases = []HTTPTestCase{
	{
		Name:     "zero test",
		Numeric:  0,
		Expected: []byte("1"),
	},
	{
		Name:     "first test",
		Numeric:  1,
		Expected: []byte("1"),
	},
	{
		Name:     "second test",
		Numeric:  3,
		Expected: []byte("6"),
	},
	{
		Name:     "third test",
		Numeric:  5,
		Expected: []byte("120"),
	},
}

func TestHandleFactorial(t *testing.T) {
	handler := http.HandlerFunc(HandlerFactorial)
	for _, test := range HTTPCases {
		t.Run(test.Name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handlerData := fmt.Sprintf("/factorial?num=%d", test.Numeric)
			request, err := http.NewRequest("GET", handlerData, nil)
			//data := io.Reader([]byte('{"num":5}'))
			//request, err := http.Post("http://localhost:8080"+handlerData, "application/json", data)

			if err != nil {
				t.Error(err)
			}

			handler.ServeHTTP(recorder, request) // run request and write response into recorder
			if string(recorder.Body.Bytes()) != string(test.Expected) {
				t.Errorf("test %s failed: input %v! result: %v expected %v",
					test.Name,
					test.Numeric,
					string(recorder.Body.Bytes()),
					string(test.Expected),
				)
			}
		})
	}
}
