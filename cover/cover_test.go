package cover

import "testing"

func TestTriangle(t *testing.T) {
	type args struct {
		base   float64
		height float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"a", args{2, 5}, 5},
		{"b", args{2, 2}, 2},
		{"c", args{11, 4}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Triangle(tt.args.base, tt.args.height); got != tt.want {
				t.Errorf("Triangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
