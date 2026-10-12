package main

import "testing"

func TestApply(t *testing.T) {
	tests := []struct {
		name    string
		op      string
		l, r    float64
		want    float64
		wantErr bool
	}{
		{"add", "+", 2, 3, 5, false},
		{"sub", "-", 2, 3, -1, false},
		{"mul", "*", 2, 3, 6, false},
		{"div", "/", 6, 3, 2, false},
		{"division by zero", "/", 2, 0, 0, true},
		{"unknown operation", "%", 1, 1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeOperation(tt.op, tt.l, tt.r)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error, but got %v", err)
			}

			if got != tt.want {
				t.Fatalf("Got %v, but want %v", got, tt.want)
			}
		})
	}
}

func TestParsing(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"valid", []string{"+", "2", "3"}, false},
		{"invalid left", []string{"+", "sds", "3"}, true},
		{"invalid right", []string{"+", "2", "abc"}, true},
		{"empty", []string{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, _, err := ParseOperands(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseOperands() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
