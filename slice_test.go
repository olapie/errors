package errors_test

import (
	"fmt"
	"strings"
	"testing"

	"code.olapie.com/errors"
)

func TestAppend(t *testing.T) {
	var err = errors.New("List of errors")
	var b strings.Builder
	b.WriteString("List of errors")
	for i := 1; i < 10; i++ {
		s := fmt.Errorf("#%d error", i)
		b.WriteString("; ")
		b.WriteString(s.Error())
		err = errors.Append(err, s)
		t.Log(b.String())
		if err.Error() != b.String() {
			t.Fatalf("Expect %s instead of %s", b.String(), err.Error())
		}
	}
}
