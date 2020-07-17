package main

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		testName      string
		tracks        []Track
		expectedSteps []string
	}{

		{
			testName: "FourOnTheFloor",
			tracks: []Track{
				{
					Name:  "bassdrum",
					ID:    0,
					Steps: "1000100010001000",
				},
				{
					Name:  "hi-hat",
					ID:    1,
					Steps: "0010001000100010",
				},
				{
					Name:  "snare",
					ID:    2,
					Steps: "0000100000001000",
				},
			},
			expectedSteps: []string{"bassdrum", "_", "hi-hat", "_", "bassdrum+snare", "_", "hi-hat", "_", "bassdrum", "_", "hi-hat", "_", "bassdrum+snare", "_", "hi-hat", "_"},
		},
		{
			testName: "StompCrash",
			tracks: []Track{
				{
					Name:  "bassdrum",
					ID:    0,
					Steps: "1111111111111111",
				},
				{
					Name:  "hi-hat",
					ID:    1,
					Steps: "1001001001001001",
				},
				{
					Name:  "cymbal",
					ID:    3,
					Steps: "0010010010010010",
				},
			},
			expectedSteps: []string{"bassdrum+hi-hat", "bassdrum", "bassdrum+cymbal", "bassdrum+hi-hat", "bassdrum", "bassdrum+cymbal", "bassdrum+hi-hat", "bassdrum", "bassdrum+cymbal", "bassdrum+hi-hat", "bassdrum", "bassdrum+cymbal", "bassdrum+hi-hat", "bassdrum", "bassdrum+cymbal", "bassdrum+hi-hat"},
		},
		{
			testName: "paradiddle",
			tracks: []Track{
				{
					Name:  "bassdrum",
					ID:    0,
					Steps: "0000000000000000",
				},
				{
					Name:  "hi-hat",
					ID:    1,
					Steps: "0000000000000000",
				},
				{
					Name:  "snare",
					ID:    2,
					Steps: "1010110101011010",
				},
			},
			expectedSteps: []string{"snare", "_", "snare", "_", "snare", "snare", "_", "snare", "_", "snare", "_", "snare", "snare", "_", "snare", "_"},
		},
		{
			testName: "all hi-hat",
			tracks: []Track{
				{
					Name:  "hi-hat",
					ID:    1,
					Steps: "1010101010101010",
				},
				{
					Name:  "snare",
					ID:    2,
					Steps: "0000000000000000",
				},
			},
			expectedSteps: []string{"hi-hat", "_", "hi-hat", "_", "hi-hat", "_", "hi-hat", "_", "hi-hat", "_", "hi-hat", "_", "hi-hat", "_", "hi-hat", "_"},
		},
	}

	for _, tc := range testCases {

		generatedSteps := buildSong(tc.tracks)
		if len(generatedSteps) != len(tc.expectedSteps) {
			t.Errorf("Failed %s", tc.testName)
			t.Errorf("Expected %v Actual %v", generatedSteps, tc.expectedSteps)
		}
		for i, v := range tc.expectedSteps {
			if v != generatedSteps[i] {
				t.Errorf("Failed %s", tc.testName)
				t.Errorf("Expected %v Actual %v", generatedSteps, tc.expectedSteps)
			}
		}
	}
}
