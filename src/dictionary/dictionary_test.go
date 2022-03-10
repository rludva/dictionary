package dictionary

import (
	"reflect"
	"testing"
)

func TestAddItem(t *testing.T) {
	type testCase struct {
		d     Dictionary
		line  string
		nextd Dictionary
	}
	testCases := []testCase{
		{
			Dictionary{},
			"hello:ahoj",
			Dictionary{
				[]DictionaryItem{
					DictionaryItem{"hello", "ahoj"},
				},
			},
		},
	}

	for _, tc := range testCases {
		pre := tc.d
		tc.d.AddItem(tc.line)
		after := tc.d
		if !reflect.DeepEqual(after, tc.nextd) {
			t.Errorf("%v.AddItem(%v) -> %v, expected %v", pre, tc.line, after, tc.nextd)
		}
	}
}
