package main

import (
	"testing"
)

func TestSumExptimeInFile(t *testing.T) {
	// Test the function
	exptime, err := getExptimeInFile("andromeda.fits")


	if err != nil {
		t.Fatalf("Error processing file: %v", err)
	}

	// Check the expected EXPTIME value
	expectedExptime := 300.0
	if exptime != expectedExptime {
		t.Errorf("Expected EXPTIME: %f, got: %f", expectedExptime, exptime)
	}
}

// Add more test cases as needed for comprehensive testing

func TestSecondsToHoursMinutes(t *testing.T) {
	tests := []struct {
		seconds  float64
		expectedHours int
		expectedMinutes int
	}{
		{3600, 1, 0},
		{9000, 2, 30},
		{150, 0, 2},
	}

	for _, test := range tests {
		hours, minutes := secondsToHoursMinutes(test.seconds)
		if hours != test.expectedHours || minutes != test.expectedMinutes {
			t.Errorf("For seconds %f, expected %d hours and %d minutes, but got %d hours and %d minutes",
				test.seconds, test.expectedHours, test.expectedMinutes, hours, minutes)
		}
	}
}

// Add more test cases for other functions and scenarios

// Note: Due to the nature of the code, some parts (such as goroutines) might be challenging to test with unit tests.
// Consider using integration tests or modifying the code for better testability if needed.

