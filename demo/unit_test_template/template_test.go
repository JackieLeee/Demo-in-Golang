package unit_test_template

import (
	"math"
	"testing"
)

/**
 * @Author  Flagship
 * @Date  2022/4/27 15:47
 * @Description
 */

func TestSum(t *testing.T) {
	cases := []struct {
		name   string
		a      int
		b      int
		target int
	}{
		{
			name:   "zero add zero",
			a:      0,
			b:      0,
			target: 0,
		},
		{
			name:   "one add one",
			a:      1,
			b:      1,
			target: 2,
		},
		{
			name:   "int64_max add int64_max",
			a:      math.MaxInt64,
			b:      math.MaxInt64,
			target: -2,
		},
		{
			name:   "int32_max add int32_max",
			a:      math.MaxInt32,
			b:      math.MaxInt32,
			target: math.MaxInt32 << 1,
		},
	}

	for _, c := range cases {
		res := sum(c.a, c.b)
		if res != c.target {
			t.Errorf("name %s, sum(%d,%d) got %d, expected %d",
				c.name, c.a, c.b, res, c.target)
		}
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
