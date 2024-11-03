package main

import (
    "bytes"
    "io"
    "os"
    "testing"
)

// TestSum tests the sum function
func TestSumMain(t *testing.T) {
    testCases := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {"positive numbers", 1, 2, 3},
        {"negative numbers", -1, -2, -3},
        {"mixed numbers", -1, 2, 1},
        {"zeros", 0, 0, 0},
        {"large numbers", 999999, 1, 1000000},
    }

    for _, tc := range testCases {
        got := sum(tc.a, tc.b)
        if got != tc.expected {
            t.Errorf("%s: sum(%d, %d) = %d; want %d", 
                tc.name, tc.a, tc.b, got, tc.expected)
        }
    }
}

// TestMain tests the main function by capturing stdout
func TestMain(t *testing.T) {
    // Backup original stdout
    oldStdout := os.Stdout

    // Create a pipe to capture stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Call main()
    main()

    // Close writer and restore stdout
    w.Close()
    os.Stdout = oldStdout

    // Read captured output
    var buf bytes.Buffer
    io.Copy(&buf, r)
    output := buf.String()

    // Check the expected output
    expected := "Sum: 3\n"
    if output != expected {
        t.Errorf("main() output = %q; want %q", output, expected)
    }
}
