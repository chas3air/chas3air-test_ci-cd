package funcs

import (
	"testing"
)

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{2, 3}, 5},
		{"", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSliceNonOpt(b *testing.B) {
	var slice []int
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		slice = append(slice, i+1)
	}
	b.StopTimer()

	_ = slice
}

func BenchmarkSlice10Opt(b *testing.B) {
	var slice = make([]int, 10)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		slice = append(slice, i+1)
	}
	b.StopTimer()

	_ = slice
}

func FuzzIsZero(f *testing.F) {
	f.Add(3)
	f.Add(-1)
	f.Add(10000000)
	f.Add(0)
	f.Add(42)
	f.Add(-100000)

	f.Fuzz(func(t *testing.T, a int) {
		isZero := IsZero(a)
		if a == 0 && isZero {
		} else if a != 0 && !isZero {
		} else {
			t.Errorf("need: %v, got: %v; a = %v", a == 0, isZero, a)
		}
	})
}
