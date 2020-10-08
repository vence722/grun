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
	Run(func(handleErr HandleErrFunc) {
		_, _, err := testFunc(1, 1)
		handleErr(err)
	}).Catch(func(err error) {
		// if error is caught, test fails
		t.FailNow()
	})

	// run code with error caught
	Run(func(handleErr HandleErrFunc) {
		// testFunc should return params successfully
		resInt, resFloat, err := testFunc(1, 1)
		handleErr(err)
		if resInt != 1 || resFloat != 1.0 {
			t.FailNow()
		}

		// testFunc should return params successfully
		resInt, resFloat, err = testFunc(2, 5)
		handleErr(err)
		if resInt != 2 || resFloat != 5.0 {
			t.FailNow()
		}

		// testFunc should throw error and go to catch block
		resInt, resFloat, err = testFunc(4, 5)
		handleErr(err)

		// following code should not be executed
		testFunc(5, 5)
		t.FailNow()
	}).Catch(func(err error) {
		if err.Error() != "error input: input1=4, input2=5.000000" {
			t.Error(err)
			t.FailNow()
		}
	})
}
