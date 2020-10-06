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
	Run(func(try TryFunc) {
		// testFunc should return params successfully
		res := try(testFunc(1, 1))
		if len(res) != 3 || res[0] != 1 || res[1] != 1.0 {
			t.Error(res)
			t.FailNow()
		}
		// testFunc should return params successfully
		res = try(testFunc(2, 5))
		if len(res) != 3 || res[0] != 2 || res[1] != 5.0 {
			t.Error(res)
			t.FailNow()
		}
		// testFunc should throw error and go to catch block
		res = try(testFunc(4, 5))
		// should not come to here
		t.Error(res)
		t.FailNow()
	}).Catch(func(err error) {
		if err.Error() != "error input: input1=4, input2=5.000000" {
			t.Error(err)
			t.FailNow()
		}
	})
}
