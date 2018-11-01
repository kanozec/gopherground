package main

import "testing"

type testSample struct {
	input  string
	expect string
}

func TestNormalize(t *testing.T) {
	cases := []testSample{
		{"0123456789", "0123456789"},
		{"418 456 7891", "4184567891"},
		{"(225) 456 7892", "2254567892"},
		{"(523) 476-7893", "5234767893"},
		{"723-656-7894", "7236567894"},
		{"(123)555-2987", "1235552987"},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			result := normalize(c.input)
			if result != c.expect {
				t.Errorf("got %s; want %s", result, c.expect)
			}
		})
	}
}
