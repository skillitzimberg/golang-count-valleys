package countvalleys

import "testing"

func TestCountValleys(t *testing.T) {
	tests := []struct {
		want int32
		got  int32
	}{
		{want: 1, got: countingValleys(8, "DDUUUUDD")},
		{want: 2, got: countingValleys(8, "DDUUUUDDDU")},
		{want: 3, got: countingValleys(8, "DUDUDU")},
	}

	for _, test := range tests {
		if test.want != test.got {
			t.Errorf("Want %d; Got %d", test.want, test.got)
		}
	}
}
