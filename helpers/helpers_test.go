package helpers

import (
	"reflect"
	"testing"
)

const input1 = `123
456
7890
1`

func TestStringToIntSlice(t *testing.T) {
	var tests = []struct{
		name string
		in   string
		want []int
		wantErr string
	}{
		{
			name: "1 number to int slice",
			in: "123",
			want: []int{123},
		},
		{
			name: "multiple numbers to int slice",
			in: input1,
			want: []int{123, 456, 7890, 1},
		},
		{
			name: "invalid numbers",
			in: "abc",
			want: []int{},
			wantErr: "could not convert to int: abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := StringToIntSlice(tt.in)
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("StringToIntSlice(%v) = error: %v, want %v", tt.in, err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(output, tt.want) {
					t.Errorf("StringToIntSlice(%v) = %v, want %v", tt.in, output, tt.want)
				}
			}
		})
	}
}