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
					{"hello", "ahoj"},
				},
			},
		},
		{
			Dictionary{},
			"hello:ahoj,nazdar",
			Dictionary{
				[]DictionaryItem{
					{"hello", "ahoj,nazdar"},
				},
			},
		},
		{
			Dictionary{
				[]DictionaryItem{
					{"sun", "slunce"},
				},
			},
			"hello:ahoj",
			Dictionary{
				[]DictionaryItem{
					{"sun", "slunce"},
					{"hello", "ahoj"},
				},
			},
		},
		{
			Dictionary{},
			"hello:  	ahoj  ",
			Dictionary{
				[]DictionaryItem{
					{"hello", "ahoj"},
				},
			},
		},
		{
			Dictionary{},
			"   hello       :  	ahoj  ",
			Dictionary{
				[]DictionaryItem{
					{"hello", "ahoj"},
				},
			},
		},
		{
			Dictionary{},
			"hello:",
			Dictionary{
				[]DictionaryItem{},
			},
		},
		{
			Dictionary{},
			":ahoj",
			Dictionary{
				[]DictionaryItem{},
			},
		},
		{
			Dictionary{},
			"Lorem ipsum dolor sit amet..",
			Dictionary{
				[]DictionaryItem{},
			},
		},
	}

	for _, tc := range testCases {
		if gotDictionary := AddItem(tc.d, tc.line); !reflect.DeepEqual(gotDictionary, tc.expectedd) {
			t.Errorf("%v.AddItem(%v): %v, expected %v", tc.d, tc.line, gotDictionary, tc.expectedd)
		}
	}
}
