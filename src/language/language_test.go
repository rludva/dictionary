package language

import "testing"

func TestNoAccent(t *testing.T) {
  type testCase struct {
    input, expectedOutput string
  }
  testCases := []testCase {
    {
      "dobrý den",
      "dobry den",
    },
  }

  for _, tc := range testCases {
    if got := NoAccent(tc.input); got != tc.expectedOutput {
      t.Errorf("NoAccent(`%v`): `%v`, expected `%v`", tc.input, got, tc.expectedOutput)
    }
  }
}
