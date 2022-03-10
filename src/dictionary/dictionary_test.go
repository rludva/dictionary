package dictionary

import (
	"reflect"
	"testing"
)

func TestAddItem(t *testing.T) {
	type testCase struct {
		d         Dictionary
		line      string
		expectedd Dictionary
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
		if gotDictionary := AddItem(tc.d, tc.line); !reflect.DeepEqual(gotDictionary, tc.expectedd) {
			t.Errorf("%v.AddItem(%v): %v, expected %v", tc.d, tc.line, gotDictionary, tc.expectedd)
		}
	}
}
