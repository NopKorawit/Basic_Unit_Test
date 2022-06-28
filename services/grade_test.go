package services_test

import (
	"fmt"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGrader1(t *testing.T) {
// 	t.Run("Grade A", func(t *testing.T) {
// 		grade := services.Grader(80)
// 		expected := "A"

// 		if grade != expected {
// 			t.Errorf("get %v expected %v", grade, expected)
// 		}
// 	})

// 	t.Run("Grade B", func(t *testing.T) {
// 		grade := services.Grader(70)
// 		expected := "B"

// 		if grade != expected {
// 			t.Errorf("get %v expected %v", grade, expected)
// 		}
// 	})

// 	t.Run("Grade C", func(t *testing.T) {
// 		grade := services.Grader(60)
// 		expected := "C"

// 		if grade != expected {
// 			t.Errorf("get %v expected %v", grade, expected)
// 		}
// 	})
// 	t.Run("Grade D", func(t *testing.T) {
// 		grade := services.Grader(50)
// 		expected := "D"

// 		if grade != expected {
// 			t.Errorf("get %v expected %v", grade, expected)
// 		}
// 	})
// 	t.Run("Grade F", func(t *testing.T) {
// 		grade := services.Grader(49)
// 		expected := "F"

// 		if grade != expected {
// 			t.Errorf("get %v expected %v", grade, expected)
// 		}
// 	})

// }

// func TestGrader2(t *testing.T) {
// 	type testCase struct {
// 		name     string
// 		score    int
// 		expected string
// 	}

// 	cases := []testCase{
// 		{name: "a", score: 80, expected: "A"},
// 		{name: "b", score: 70, expected: "B"},
// 		{name: "c", score: 60, expected: "C"},
// 		{name: "d", score: 50, expected: "D"},
// 		{name: "f", score: 0, expected: "F"},
// 	}
// 	for _, c := range cases {
// 		t.Run(c.name, func(t *testing.T) {
// 			grade := services.Grader(c.score)

// 			if grade != c.expected {
// 				t.Errorf("get %v expected %v", grade, c.expected)
// 			}
// 		})
// 	}
// }

func TestGrader3(t *testing.T) {
	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "a", score: 80, expected: "A"},
		{name: "b", score: 70, expected: "B"},
		{name: "c", score: 60, expected: "C"},
		{name: "d", score: 50, expected: "D"},
		{name: "f", score: 0, expected: "F"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.Grader(c.score)

			assert.Equal(t, c.expected, grade) //แทน3บรรทัดล่าง
			// if grade != c.expected {
			// 	t.Errorf("get %v expected %v", grade, c.expected)
			// }
		})
	}
}

// go test gotest/services -bench.
// go test gotest/services -bench. -benchmem
func BenchmarkGrader(b *testing.B) {

	for i := 0; i < b.N; i++ {
		services.Grader(80)
	}
}

func ExampleGrader() {
	grade := services.Grader(80)
	fmt.Println(grade)
	// Output: A
}
