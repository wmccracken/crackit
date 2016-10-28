package util

import (
	"errors"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func TestContainsLessThan(t *testing.T) {
	cases := []struct {
		in   []int
		max  int
		want bool
	}{
		{[]int{0, 0, 0}, 4, true},
		{[]int{5, 5, 5}, 4, false},
		{[]int{0, 5, 0}, 4, true},
		{[]int{0, 0, 5}, 4, true},
		{[]int{5, 0, 0}, 4, true},
	}
	for _, c := range cases {
		got := ContainsLessThan(c.in[:], c.max)
		if got != c.want {
			t.Errorf("ContainsLessThan(%q, %d) == %t, want %t", c.in, c.max, got, c.want)
		}
	}
}

func (s *MySuite) TestCheck(c *C) {
	var e = errors.New("ACK")
	c.Assert(func() { Check(e) }, Panics, errors.New("ACK"))
}
