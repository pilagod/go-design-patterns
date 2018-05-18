package pipeline

import "testing"

func TestLaunchPipeline(t *testing.T) {
	testCases := [][]int{
		{3, 14},
		{5, 55},
	}
	for _, testCase := range testCases {
		res := LaunchPipeline(testCase[0])

		if res != testCase[1] {
			t.Fatalf("expected: %d, but got: %d", testCase[1], res)
		}
		t.Logf("%d == %d\n", res, testCase[1])
	}
}
