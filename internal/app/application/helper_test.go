package application_test

import (
	"testing"
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/application"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

func TestMapRepoApplCompEmpToService(t *testing.T) {
	type testCase struct {
		name           string
		input          repo.ApplicationCompleteEmp
		expectedOutput application.ApplicationCompleteEmp
	}

	testCases := []testCase{
		{
			name: "success",
			input: repo.ApplicationCompleteEmp{
				ID:             1,
				JobID:          1,
				WorkerID:       1,
				Status:         "Pending",
				ExpectedWage:   1500,
				ModeOfArrival:  "Personal",
				PickUpLocation: 1,
				WorkerComment:  "comments",
				AppliedAt:      time.Time{},
				UpdatedAt:      time.Time{},
				Details:        "details",
				Street:         "street",
				City:           "city",
				State:          "state",
				Pincode:        411052,
				JobTitle:       "Title",
				Description:    "random description",
				SkillsRequired: "skills",
				JobSectors:     "sectors",
				JobWage:        1500,
				Vacancy:        5,
				JobDate:        "",
				WorkerName:     "John Doe",
				ContactNumber:  "9067691136",
				WorkerEmail:    "jhon@gmail.com",
				WorkerGender:   "Male",
			},
			expectedOutput: application.ApplicationCompleteEmp{
				ID:             1,
				JobID:          1,
				WorkerID:       1,
				Status:         "Pending",
				ExpectedWage:   1500,
				ModeOfArrival:  "Personal",
				PickUpLocation: 1,
				WorkerComment:  "comments",
				AppliedAt:      time.Time{},
				UpdatedAt:      time.Time{},
				Details:        "details",
				Street:         "street",
				City:           "city",
				State:          "state",
				Pincode:        411052,
				JobTitle:       "Title",
				Description:    "random description",
				SkillsRequired: "skills",
				JobSectors:     "sectors",
				JobWage:        1500,
				Vacancy:        5,
				JobDate:        "",
				WorkerName:     "John Doe",
				ContactNumber:  "9067691136",
				WorkerEmail:    "jhon@gmail.com",
				WorkerGender:   "Male",
			},
		},
	}

	for _, test := range testCases {
		if application.MapRepoApplCompEmpToService(test.input) != test.expectedOutput {
			t.Error("Expected and actual outputs are not equal.")
		}
	}
}
