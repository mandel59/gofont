package otf

import "testing"

func TestMaxPowerOf2(t *testing.T) {
	type testSet struct{ in, out USHORT }
	var test = []testSet{
		{1, 0},
		{4, 2},
		{6, 2},
		{0, 0},
	}
	for i, v := range test {
		if x := maxPowerOf2(v.in); x != v.out {
			t.Errorf("Test %v maxPowerOf2(`%v`) = %v, want %v", i, v.in, x, v.out)
		}
	}
}

func TestCalcCheckSum(t *testing.T) {
	type testSet struct {
		in  string
		out ULONG
	}
	var test = []testSet{
		{"abcd", 1633837924},
		{"abcdxyz", 3655064932},
	}
	for i, v := range test {
		if x := calcCheckSum([]byte(v.in)); x != v.out {
			t.Errorf("Test %v calcCheckSum([]byte(`%v`)) = %v, want %v", i, v.in, x, v.out)
		}
	}
}
