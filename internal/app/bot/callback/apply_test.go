package callback

import "testing"

func TestApplyAddCalendarMask(t *testing.T) {
	type args struct {
		groupID      int64
		languageCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default",
			args: args{
				groupID:      77991100,
				languageCode: "en",
			},
			want: "add-calendar:77991100:en",
		},
		{
			name: "Zero-values",
			args: args{
				groupID:      0,
				languageCode: "",
			},
			want: "add-calendar:0:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyAddCalendarMask(tt.args.groupID, tt.args.languageCode); got != tt.want {
				t.Errorf("ApplyAddCalendarMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyFavouriteGroupMask(t *testing.T) {
	type args struct {
		groupID    int64
		groupTitle string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default",
			args: args{
				groupID:    77991100,
				groupTitle: "6101-020302D",
			},
			want: "favourite-group:77991100:6101-020302D",
		},
		{
			name: "Zero-values",
			args: args{
				groupID:    0,
				groupTitle: "",
			},
			want: "favourite-group:0:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyFavouriteGroupMask(tt.args.groupID, tt.args.groupTitle); got != tt.want {
				t.Errorf("ApplyFavouriteGroupMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyScheduleMask(t *testing.T) {
	type args struct {
		groupID    int64
		groupTitle string
		week       int
		weekday    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default",
			args: args{
				groupID:    77991100,
				groupTitle: "6101-020302D",
				week:       26,
				weekday:    3,
			},
			want: "schedule-daily:77991100:6101-020302D:26:3",
		},
		{
			name: "Zero-values",
			args: args{
				groupID:    0,
				groupTitle: "",
				week:       0,
				weekday:    0,
			},
			want: "schedule-daily:0::0:0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyScheduleMask(tt.args.groupID, tt.args.groupTitle, tt.args.week, tt.args.weekday); got != tt.want {
				t.Errorf("ApplyScheduleMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplyScheduleTodayMask(t *testing.T) {
	type args struct {
		groupID    int64
		groupTitle string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default",
			args: args{
				groupID:    77991100,
				groupTitle: "6101-020302D",
			},
			want: "schedule-today:77991100:6101-020302D",
		},
		{
			name: "Zero-values",
			args: args{
				groupID:    0,
				groupTitle: "",
			},
			want: "schedule-today:0:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplyScheduleTodayMask(tt.args.groupID, tt.args.groupTitle); got != tt.want {
				t.Errorf("ApplyScheduleTodayMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplySetLanguageMask(t *testing.T) {
	type args struct {
		languageCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default",
			args: args{
				languageCode: "en",
			},
			want: "set-language:en",
		},
		{
			name: "Zero-values",
			args: args{
				languageCode: "",
			},
			want: "set-language:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ApplySetLanguageMask(tt.args.languageCode); got != tt.want {
				t.Errorf("ApplySetLanguageMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
