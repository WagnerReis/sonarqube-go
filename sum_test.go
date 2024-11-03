package main

import "testing"

func TestSum(t *testing.T) {
    // Test case structure
    testCases := []struct {
        name     string
        a, b     int
        expected int
    }{
        {
            name:     "positive numbers",
            a:        2,
            b:        3,
            expected: 5,
        },
        {
            name:     "negative numbers",
            a:        -2,
            b:        -3,
            expected: -5,
        },
        {
            name:     "zero and positive",
            a:        0,
            b:        5,
            expected: 5,
        },
        {
            name:     "mixed numbers",
            a:        -2,
            b:        3,
            expected: 1,
        },
    }

    // Run all test cases
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := sum(tc.a, tc.b)
            if result != tc.expected {
                t.Errorf("sum(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
            }
        })
    }
}

