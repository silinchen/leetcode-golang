package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tt := range tests {
		ans := lengthOfLongestSubstring(tt.s)
		if ans != tt.ans {
			t.Errorf("got %v for input %v; expected %v", ans, tt.s, tt.ans)
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	s := "abcabcbb"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 3
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfLongestSubstring(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
