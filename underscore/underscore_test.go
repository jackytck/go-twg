package underscore

import "testing"

func TestCamel(t *testing.T) {
	testCases := []struct {
		arg  string
		want string
	}{
		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
		{"with a space", "with a space"},
		{"endsWithA", "ends_with_a"},
	}
	for _, tc := range testCases {
		got := Camel(tc.arg)
		if got != tc.want {
			t.Errorf("Camel(%q) = %q; want %q\n", tc.arg, got, tc.want)
		}

	}
}
