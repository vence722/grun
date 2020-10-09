package grun

import (
	"errors"
	"fmt"
	"testing"
)

func testFunc(input1 int, input2 float64) (int, float64, error) {
	// generate error
	if input1 > 3 && input2 > 4.5 {
		return 0, 0, errors.New(fmt.Sprintf("error input: input1=%d, input2=%f", input1, input2))
	}
	return input1, input2, nil
}

func TestRun(t *testing.T) {
	// run code without error
	Run(func(throw ThrowFunc) {
		_, _, err := testFunc(1, 1)
		throw("testFunc(1, 1)", err)
	}).Catch(func(err CaughtError) {
		// if error is caught, test fails
		t.FailNow()
	})

	// run code with error caught
	Run(func(throw ThrowFunc) {
		// testFunc should return params successfully
		resInt, resFloat, err := testFunc(1, 1)
		throw("testFunc(1, 1)", err)
		if resInt != 1 || resFloat != 1.0 {
			t.FailNow()
		}

		// testFunc should return params successfully
		resInt, resFloat, err = testFunc(2, 5)
		throw("testFunc(2, 5)", err)
		if resInt != 2 || resFloat != 5.0 {
			t.FailNow()
		}

		// testFunc should throw error and go to catch block
		resInt, resFloat, err = testFunc(4, 5)
		throw("testFunc(4, 5)", err)

		// following code should not be executed
		resInt, resFloat, err = testFunc(5, 5)
		throw("testFunc(5, 5)", err)
		t.FailNow()
	}).Catch(func(err CaughtError) {
		if err.Name != "testFunc(4, 5)" || err.Err.Error() != "error input: input1=4, input2=5.000000" {
			t.Error(err)
			t.FailNow()
		}
	})
}
