package utils

import (
	"testing"
)

func TestGetRequest(t *testing.T) {
	tableTest := []struct {
		name           string
		req            map[string]string
		expectedResult string
	}{
		{
			name: "TEST convert map string",
			req: map[string]string{
				"apikey": "faf7e5bb",
				"s":      "batman",
				"p":      "2",
			},
			expectedResult: `apikey="faf7e5bb", s="batman", p="2", `,
		},
	}

	for _, test := range tableTest {
		resp := ConvertMapToString(test.req)

		if resp == "" {
			t.Error(test.name + " Invalid result")
		}
	}

}
