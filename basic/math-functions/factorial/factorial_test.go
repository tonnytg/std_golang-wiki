package factorial

import "testing"

func TestFactorial(t *testing.T) {

    var tests = []struct {
        input int
        expected int
    }{
        {0, 1},
        {1, 1},
        {2, 2},
        {3, 6},
        {4, 24},
        {5, 120},
        {6, 720},
        {7, 5040},
        {8, 40320},
        {9, 362880},
        {10, 3628800},
    }

    for _, test := range tests {
        if output := Factorial(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
        }
    }
}