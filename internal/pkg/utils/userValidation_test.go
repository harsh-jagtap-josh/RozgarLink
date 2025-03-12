package utils

import (
	"testing"
)

func TestValidateName(t *testing.T) {
	type testCase struct {
		name          string
		input         string
		expectedError bool
	}
	testCases := []testCase{
		{
			name:          "success",
			input:         "Harsh Jagtap",
			expectedError: false,
		},
		{
			name:          "validation error",
			input:         "123",
			expectedError: true,
		},
	}

	for _, test := range testCases {
		err := ValidateName(test.input)
		if test.expectedError == (err == nil) {
			t.Error("Error doesn't match")
		}
	}
}

func TestValidateEmail(t *testing.T) {
	type testCase struct {
		name          string
		input         string
		expectedError bool
	}
	testCases := []testCase{
		{
			name:          "success",
			input:         "harsh@gmail.com",
			expectedError: false,
		},
		{
			name:          "validation error",
			input:         "harshgmail.com",
			expectedError: true,
		},
	}

	for _, test := range testCases {
		err := ValidateEmail(test.input)
		if test.expectedError == (err == nil) {
			t.Error("Error doesn't match")
		}
	}
}

func TestValidateMobileNumber(t *testing.T) {
	type testCase struct {
		name          string
		input         string
		expectedError bool
	}
	testCases := []testCase{
		{
			name:          "success",
			input:         "9067691363",
			expectedError: false,
		},
		{
			name:          "validation error",
			input:         "906769136",
			expectedError: true,
		},
	}

	for _, test := range testCases {
		err := ValidateMobileNumber(test.input)
		if test.expectedError == (err == nil) {
			t.Error("Error doesn't match")
		}
	}
}

func TestValidatePassword(t *testing.T) {
	type testCase struct {
		name          string
		input         string
		expectedError bool
	}
	testCases := []testCase{
		{
			name:          "success",
			input:         "harsh@123",
			expectedError: false,
		},
		{
			name:          "validation error",
			input:         "hey",
			expectedError: true,
		},
	}

	for _, test := range testCases {
		err := ValidatePassword(test.input)
		if test.expectedError == (err == nil) {
			t.Error("Error doesn't match")
		}
	}
}

type validateUserInput struct {
	name     string
	email    string
	mobile   string
	password string
}

type validateUpdateUserInput struct {
	name   string
	email  string
	mobile string
}

func TestValidateUser(t *testing.T) {
	type testCase struct {
		testName      string
		input         validateUserInput
		expectedError bool
	}
	testCases := []testCase{
		{
			testName: "success",
			input: validateUserInput{
				name:     "Harsh Jagtap",
				email:    "harsh@gmail.com",
				mobile:   "9067691363",
				password: "Harsh@123",
			},
			expectedError: false,
		},
		{
			testName: "invalid name",
			input: validateUserInput{
				name:     "Harsh Jagtap1",
				email:    "harsh@gmail.com",
				mobile:   "9067691363",
				password: "Harsh@123",
			},
			expectedError: true,
		},
		{
			testName: "invalid email",
			input: validateUserInput{
				name:     "Harsh Jagtap",
				email:    "harsh.com",
				mobile:   "9067691363",
				password: "Harsh@123",
			},
			expectedError: true,
		},
		{
			testName: "invalid mobile",
			input: validateUserInput{
				name:     "Harsh Jagtap",
				email:    "harsh@gmail.com",
				mobile:   "90676913",
				password: "Harsh@123",
			},
			expectedError: true,
		},
		{
			testName: "invalid password",
			input: validateUserInput{
				name:     "Harsh Jagtap",
				email:    "harsh@gmail.com",
				mobile:   "9067691363",
				password: "harsh",
			},
			expectedError: true,
		},
	}

	for _, test := range testCases {
		err := ValidateUser(test.input.name, test.input.mobile, test.input.email, test.input.password)
		if test.expectedError == (err == nil) {
			t.Error("Error doesn't match")
		}
	}
}

func TestValiUpdateUser(t *testing.T) {
	type testCase struct {
		testName      string
		input         validateUpdateUserInput
		expectedError bool
	}
	testCases := []testCase{
		{
			testName: "success",
			input: validateUpdateUserInput{
				name:   "Harsh Jagtap",
				email:  "harsh@gmail.com",
				mobile: "9067691363",
			},
			expectedError: false,
		},
		{
			testName: "invalid name",
			input: validateUpdateUserInput{
				name:   "Harsh Jagtap1",
				email:  "harsh@gmail.com",
				mobile: "9067691363",
			},
			expectedError: true,
		},
		{
			testName: "invalid email",
			input: validateUpdateUserInput{
				name:   "Harsh Jagtap",
				email:  "harsh.com",
				mobile: "9067691363",
			},
			expectedError: true,
		},
		{
			testName: "invalid mobile",
			input: validateUpdateUserInput{
				name:   "Harsh Jagtap",
				email:  "harsh@gmail.com",
				mobile: "90676913",
			},
			expectedError: true,
		},
	}

	for _, test := range testCases {
		err := ValidateUpdateUser(test.input.name, test.input.mobile, test.input.email)
		if test.expectedError == (err == nil) {
			t.Error("Error doesn't match")
		}
	}
}
