package collatzconjecture

import (
	"testing"
)

func TestCollatzConjecture(t *testing.T) {
	for _, testCase := range testCases {
		steps := CollatzConjecture(testCase.input)
		if steps != testCase.expected {
			t.Fatalf("FAIL: %s\n\tCollatzConjecture(%v) expected %v, got %v",
				testCase.description, testCase.input, testCase.expected, steps)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkCollatzConjecture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			CollatzConjecture(testCase.input)
		}
	}
}
