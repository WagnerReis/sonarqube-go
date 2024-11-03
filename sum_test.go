package main

import "testing"

func TestSum(t *testing.T) {
    // Test case structure
    testCases := []struct {
        a, b     int
        expected int
    }{
        {2, 3, 5},      // positive numbers
        {-2, -3, -5},   // negative numbers
        {0, 5, 5},      // zero and positive
        {-2, 3, 1},     // mixed numbers
    }

    // Run all test cases
    for _, tc := range testCases {
        result := sum(tc.a, tc.b)
        if result != tc.expected {
            t.Errorf("sum(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
        }
    }
}
