package str

import "testing"

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Empty string",
			s:    "",
			want: "",
		},
		{
			name: "No letters",
			s:    "12 0?..",
			want: "12 0?..",
		},
		{
			name: "First symbol isn't letter",
			s:    "1word",
			want: "1word",
		},
		{
			name: "English letters",
			s:    "english",
			want: "English",
		},
		{
			name: "Russian letters",
			s:    "русский",
			want: "Русский",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.s); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
