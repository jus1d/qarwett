package random

import (
	"reflect"
	"testing"
)

func TestChoice(t *testing.T) {
	type args[T any] struct {
		arr []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test result contains in array",
			args: args[int]{
				arr: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Choice(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Choice() = %v, want %v", got, tt.want)
			}
		})
	}
}
