package main

import (
	"github.com/elliotchance/testify-stats/assert"
	"testing"
)

var calculateTests =  []struct {
	expression string
	expected    int
	err        error
}{
	{"3+2*2",
	7,
	nil,
	},
}

func TestCalculation(t *testing.T) {
	for _, test := range calculateTests {
		t.Run("", func(t *testing.T) {
			newRepl := repl{}
			result, err := newRepl.calculate(test.expression)
            assert.Equal(t, test.expected, result)
			assert.Equal(t, test.err, err)
		})
	}
}