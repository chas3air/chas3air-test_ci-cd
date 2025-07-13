package chas3airtestcicd

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		a, b   int
		result int
	}{
		{"positive numbers", 1, 5, 6},
		{"zero and positive", 0, 5, 5},
		{"negative and positive", -3, 7, 4},
		{"two negatives", -2, -8, -10},
		{"large numbers", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.result, Sum(tt.a, tt.b))
		})
	}
}

func BenchmarkSum(b *testing.B) {
	cases := []struct {
		a, b int
	}{
		{1, 2},
		{100, 200},
		{-50, 50},
		{0, 0},
	}
	for _, c := range cases {
		b.Run(fmt.Sprintf("%d+%d", c.a, c.b), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sum(c.a, c.b)
			}
		})
	}
}

func FuzzSum(f *testing.F) {
	f.Add(1, 4)
	f.Add(-10, 10)
	f.Add(0, 0)
	f.Add(1000000, -1000000)
	f.Fuzz(func(t *testing.T, a int, b int) {
		got := Sum(a, b)
		// Проверяем, что сумма обратно вычисляется правильно
		if got-a != b {
			t.Errorf("Sum(%d, %d) = %d; got-a != b", a, b, got)
		}
		if got-b != a {
			t.Errorf("Sum(%d, %d) = %d; got-b != a", a, b, got)
		}
	})
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
