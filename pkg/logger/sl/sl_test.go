package sl

import (
	"errors"
	"log/slog"
	"reflect"
	"testing"
)

func TestErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want slog.Attr
	}{
		{
			name: "Check error",
			args: args{err: errors.New("some error")},
			want: slog.Attr{Key: "error", Value: slog.StringValue("some error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Err() = %v, want %v", got, tt.want)
			}
		})
	}
}
