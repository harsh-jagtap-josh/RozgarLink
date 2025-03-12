package apperrors

import "testing"

type testCase struct {
	name           string
	warning        string
	message        string
	id             string
	expectedOutput string
}

func TestHttpErrorResponse(t *testing.T) {
	testCases := []testCase{
		{
			name:           "success",
			message:        "some message",
			warning:        "some warning",
			id:             "1",
			expectedOutput: "some warning: some message, id: 1",
		},
	}

	for _, test := range testCases {
		output := HttpErrorResponseMessage(test.warning, test.message, test.id)
		if output != test.expectedOutput {
			t.Error("Output not equal to expected output")
		}
	}
}
