package intutils

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IniMin(2, -2) = %d, want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	tests := []struct{
		a, b int
		want int
	} {
		{ 1, 1, 1 },
		{ 1, 2, 1 },
		{ 2, 1, 1 },
		{ -1, 0, -1 },
		{ 0, -1, -1 },
	}

	for _, tt  := range tests {
		name := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(name, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Fatalf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
