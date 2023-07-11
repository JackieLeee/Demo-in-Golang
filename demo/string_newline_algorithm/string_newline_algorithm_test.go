package string_newline_algorithm

import "testing"

func TestSplitLineAndKeepWord(t *testing.T) {
	cases := []struct {
		str    string
		maxLen int
		expect []string
	}{
		{
			str:    "!!!This is a test. This only a test.",
			maxLen: 14,
			expect: []string{
				"!!!This is a",
				"test. This",
				"only a test.",
			},
		},
	}
	for _, c := range cases {
		lines := WrapTextToLines(c.str, c.maxLen)
		if len(lines) != len(c.expect) {
			t.Errorf("len(lines) != len(c.expect), len(lines)=%d, len(c.expect)=%d", len(lines), len(c.expect))
		}
		for i := range lines {
			if len(lines[i]) > c.maxLen {
				t.Errorf("len(lines[%d]) > c.maxLen, len(lines[%d])=%d, c.maxLen=%d", i, i, len(lines[i]), c.maxLen)
			}
			if lines[i] != c.expect[i] {
				t.Errorf("lines[%d] != c.expect[%d], lines[%d]=%s, c.expect[%d]=%s", i, i, i, lines[i], i, c.expect[i])
			}
		}
	}
}
