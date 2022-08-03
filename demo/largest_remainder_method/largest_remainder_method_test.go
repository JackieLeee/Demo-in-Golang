package largest_remainder_method

import "testing"

func TestGetPercentValue(t *testing.T) {
	cases := []struct {
		list             []int64
		round            int
		setMin4NoZeroVal bool
		want             []float64
	}{
		{
			list:             []int64{1, 2, 3, 4, 5},
			round:            2,
			setMin4NoZeroVal: true,
			want:             []float64{6.67, 13.33, 20, 26.67, 33.33},
		},
		{
			list:             []int64{100000, 1},
			round:            2,
			setMin4NoZeroVal: true,
			want:             []float64{99.99, 0.01},
		},
		{
			list:             []int64{100000, 1},
			round:            2,
			setMin4NoZeroVal: false,
			want:             []float64{100, 0},
		},
		{
			list:             []int64{100000, 1},
			round:            3,
			setMin4NoZeroVal: false,
			want:             []float64{99.999, 0.001},
		},
		{
			list:             []int64{139150100, 6660000, 3656900, 100},
			round:            2,
			setMin4NoZeroVal: true,
			want:             []float64{93.09, 4.45, 2.45, 0.01},
		},
	}
	for _, c := range cases {
		got := GetPercentValue(c.list, c.round, c.setMin4NoZeroVal)
		for i := range got {
			if got[i] != c.want[i] {
				t.Errorf("GetPercentValue(%v, %d) = %v, want %v", c.list, c.round, got, c.want)
				break
			}
		}
	}
}
