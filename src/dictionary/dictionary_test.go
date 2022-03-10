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
		if gotDictionary := AddItem(tc.d, tc.line); !reflect.DeepEqual(gotDictionary, tc.nextd) {
			t.Errorf("%v.AddItem(%v): %v, expected %v", tc.d, tc.line, gotDictionary, tc.nextd)
		}
	}
}
