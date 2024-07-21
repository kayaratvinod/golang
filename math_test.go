package math
import "testing"



// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
    arg1, arg2, expected int
}

type subtractTest struct {
    arg1, arg2, expected int
}

var addTests = []addTest{
    {2, 3, 5},
    {4, 8, 12},
    {6, 9, 15},
    {3, 10, 13},
}

var subtractTests = []subtractTest{
    {5, 3, 2},
    {8, 4, 4},
    {15, 9, 6},
    {13, 10, 3},
}

func TestAdd(t *testing.T) {
    for _, test := range addTests {
        if output := Add(test.arg1, test.arg2); output != test.expected {
            t.Errorf("Output %d not equal to expected %d", output, test.expected)
        }
    }
}

func TestSubtract(t *testing.T) {
    for _, test := range subtractTests {
        if output := Subtract(test.arg1, test.arg2); output != test.expected {
            t.Errorf("Output %d not equal to expected %d", output, test.expected)
        }
    }
}

