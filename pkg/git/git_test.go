package git

import "testing"

func TestGetLatestCommit(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			name:    "Default",
			want:    7,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLatestCommit()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestCommit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("len(GetLatestCommit()) got = %v, want %v", got, tt.want)
			}
		})
	}
}
