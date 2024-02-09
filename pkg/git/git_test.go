package git

import "testing"

func TestGetLatestCommit(t *testing.T) {
	tests := []struct {
		name    string
		wantLen int
	}{
		{
			name:    "Default",
			wantLen: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLatestCommit()
			if len(got) != tt.wantLen {
				t.Errorf("len(GetLatestCommit()) got = %v, want %v", got, tt.wantLen)
			}
		})
	}
}
