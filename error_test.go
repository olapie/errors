package errors_test

import (
	"testing"

	"code.olapie.com/errors"
)

func TestFromString(t *testing.T) {
	err := errors.FromString("code=1, message=hello")
	if err == nil || err.Code != 1 || err.Message != "hello" {
		t.Fail()
		return
	}

	err = errors.FromString("code=, message=")
	if err != nil {
		t.Fail()
		return
	}

	err = errors.FromString("code=1, message=")
	if err == nil || err.Code != 1 || err.Message != "" {
		t.Fail()
		return
	}

	err = errors.FromString("code=1.1, message=")
	if err != nil {
		t.Fail()
		return
	}

	err = errors.FromString("code=10, message=ha ha")
	if err == nil || err.Code != 10 || err.Message != "ha ha" {
		t.Fail()
		return
	}
}
