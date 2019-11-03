package main

import "testing"

var trueVal bool = true
var falseVal bool = false

func Test_not(t *testing.T) {
	type args struct {
		o *bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Check not(true)", args{&trueVal}, falseVal},
		{"Check not(false)", args{&falseVal}, trueVal},
		{"Check not(nil)", args{nil}, trueVal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := not(tt.args.o); *got != tt.want {
				t.Errorf("not() = %v, want %v", *got, tt.want)
			}
		})
	}
}
