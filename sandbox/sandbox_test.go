package sandbox

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc         string
		data         string
		searchString string
		expected     bool
	}{
		{
			desc: "test_found",
			data: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor ` +
				`incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud ` +
				`exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute ` +
				`irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla ` +
				`pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia ` +
				`deserunt mollit anim id est laborum.`,
			searchString: "aliquip",
			expected:     true,
		},
		{
			desc: "test_not_found",
			data: `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor ` +
				`incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud ` +
				`exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute ` +
				`irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla ` +
				`pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia ` +
				`deserunt mollit anim id est laborum.`,
			searchString: "abrakadabra",
			expected:     false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			buffer := bytes.NewBuffer([]byte(tC.data))

			result, err := Contains(buffer, []byte(tC.searchString))
			if err != nil {
				t.Error(err)
			}

			if result != tC.expected {
				t.Errorf("Unexpected result from test: %s, expected: %v, but got: %v ", tC.desc, tC.expected, result)
			}
		})
	}
}
