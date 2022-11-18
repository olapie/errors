package errors_test

import (
	"testing"

	"code.olapie.com/errors"
)

func TestFromString(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		cases := map[string]*errors.Error{
			"code:1, message:hello": &errors.Error{
				Code:    1,
				Message: "hello",
			},
			" code:1, message:": &errors.Error{
				Code: 1,
			},
			"code:11": &errors.Error{
				Code: 11,
			},
			"code:10, message:ha ha": &errors.Error{
				Code:    10,
				Message: "ha ha",
			},
		}

		for s, e := range cases {
			err := errors.FromString(s)
			if err == nil {
				t.Fatal(s)
			}

			t.Log(err)

			if err.Code != e.Code {
				t.Fatalf("%s %v", s, err)
			}

			if err.Message != e.Message {
				t.Fatalf("%s %v", s, err)
			}
		}
	})

	t.Run("bad", func(t *testing.T) {
		cases := []string{
			"code:, message:",
			"code:1.1, message:",
			"s code:1, message:",
		}

		for _, s := range cases {
			err := errors.FromString(s)
			if err != nil {
				t.Fatal(s)
			}
		}
	})
}
