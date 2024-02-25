package callback

import "testing"

func TestExtractFromAddCalendarCallback(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name             string
		args             args
		wantGroupID      int64
		wantLanguageCode string
	}{
		{
			name:             "Default",
			args:             args{query: "add-calendar:77991100:en"},
			wantGroupID:      77991100,
			wantLanguageCode: "en",
		},
		{
			name:             "Zero-values",
			args:             args{query: "add-calendar:0:"},
			wantGroupID:      0,
			wantLanguageCode: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroupID, gotLanguageCode := ExtractFromAddCalendarCallback(tt.args.query)
			if gotGroupID != tt.wantGroupID {
				t.Errorf("ExtractFromAddCalendarCallback() gotGroupID = %v, want %v", gotGroupID, tt.wantGroupID)
			}
			if gotLanguageCode != tt.wantLanguageCode {
				t.Errorf("ExtractFromAddCalendarCallback() gotLanguageCode = %v, want %v", gotLanguageCode, tt.wantLanguageCode)
			}
		})
	}
}

func TestExtractFromFavouriteGroupCallback(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name           string
		args           args
		wantGroupID    int64
		wantGroupTitle string
	}{
		{
			name:           "Default",
			args:           args{query: "favourite-group:77991100:6101-020302D"},
			wantGroupID:    77991100,
			wantGroupTitle: "6101-020302D",
		},
		{
			name:           "Zero-values",
			args:           args{query: "favourite-group:0:"},
			wantGroupID:    0,
			wantGroupTitle: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroupID, gotGroupTitle := ExtractFromFavouriteGroupCallback(tt.args.query)
			if gotGroupID != tt.wantGroupID {
				t.Errorf("ExtractFromFavouriteGroupCallback() gotGroupID = %v, want %v", gotGroupID, tt.wantGroupID)
			}
			if gotGroupTitle != tt.wantGroupTitle {
				t.Errorf("ExtractFromFavouriteGroupCallback() gotGroupTitle = %v, want %v", gotGroupTitle, tt.wantGroupTitle)
			}
		})
	}
}

func TestExtractFromScheduleCallback(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name           string
		args           args
		wantGroupID    int64
		wantGroupTitle string
		wantWeek       int
		wantWeekday    int
	}{
		{
			name:           "Default",
			args:           args{query: "schedule-daily:77991100:6101-020302D:26:3"},
			wantGroupID:    77991100,
			wantGroupTitle: "6101-020302D",
			wantWeek:       26,
			wantWeekday:    3,
		},
		{
			name:           "Zero-values",
			args:           args{query: "schedule-daily:0::0:0"},
			wantGroupID:    0,
			wantGroupTitle: "",
			wantWeek:       0,
			wantWeekday:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroupID, gotGroupTitle, gotWeek, gotWeekday := ExtractFromScheduleCallback(tt.args.query)
			if gotGroupID != tt.wantGroupID {
				t.Errorf("ExtractFromScheduleCallback() gotGroupID = %v, want %v", gotGroupID, tt.wantGroupID)
			}
			if gotGroupTitle != tt.wantGroupTitle {
				t.Errorf("ExtractFromScheduleCallback() gotGroupTitle = %v, want %v", gotGroupTitle, tt.wantGroupTitle)
			}
			if gotWeek != tt.wantWeek {
				t.Errorf("ExtractFromScheduleCallback() gotWeek = %v, want %v", gotWeek, tt.wantWeek)
			}
			if gotWeekday != tt.wantWeekday {
				t.Errorf("ExtractFromScheduleCallback() gotWeekday = %v, want %v", gotWeekday, tt.wantWeekday)
			}
		})
	}
}

func TestExtractFromScheduleTodayCallback(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name           string
		args           args
		wantGroupID    int64
		wantGroupTitle string
	}{
		{
			name:           "Default",
			args:           args{query: "schedule-today:77991100:6101-020302D"},
			wantGroupID:    77991100,
			wantGroupTitle: "6101-020302D",
		},
		{
			name:           "Zero-values",
			args:           args{query: "schedule-today:0:"},
			wantGroupID:    0,
			wantGroupTitle: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroupID, gotGroupTitle := ExtractFromScheduleTodayCallback(tt.args.query)
			if gotGroupID != tt.wantGroupID {
				t.Errorf("ExtractFromScheduleTodayCallback() gotGroupID = %v, want %v", gotGroupID, tt.wantGroupID)
			}
			if gotGroupTitle != tt.wantGroupTitle {
				t.Errorf("ExtractFromScheduleTodayCallback() gotGroupTitle = %v, want %v", gotGroupTitle, tt.wantGroupTitle)
			}
		})
	}
}

func TestExtractFromSetLanguageCallback(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name             string
		args             args
		wantLanguageCode string
	}{
		{
			name:             "Default",
			args:             args{query: "set-language:en"},
			wantLanguageCode: "en",
		},
		{
			name:             "Default",
			args:             args{query: "set-language:"},
			wantLanguageCode: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLanguageCode := ExtractFromSetLanguageCallback(tt.args.query); gotLanguageCode != tt.wantLanguageCode {
				t.Errorf("ExtractFromSetLanguageCallback() = %v, want %v", gotLanguageCode, tt.wantLanguageCode)
			}
		})
	}
}
