package postgres

import (
	"qarwett/internal/config"
	"testing"
)

func TestStorage_Calendar(t *testing.T) {
	cfg := config.Postgres{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Name:     "postgres",
		Password: "1234",
		ModeSSL:  "disable",
	}

	storage, err := New(cfg)
	if err != nil {
		t.Log("Can't connect to database. Skip this test...")
		return
	}

	calendars, err := storage.GetAllTrackedCalendars()
	if err != nil {
		t.Errorf("error getting all calendars: %v", err)
	}
	if len(calendars) != 0 {
		t.Errorf("expected len(calendars): 0, got: %d", len(calendars))
	}

	calendar := Calendar{
		GroupID:      755922237,
		LanguageCode: "ru",
	}

	got, err := storage.GetTrackedCalendar(calendar.GroupID, calendar.LanguageCode)
	if err == nil {
		t.Errorf("should return error")
	}

	got, err = storage.CreateTrackedCalendar(calendar.GroupID, calendar.LanguageCode)
	if err != nil {
		t.Errorf("error creating tracked calendar: %v", err)
	}

	if got.ID == "" || got.GroupID != calendar.GroupID || got.LanguageCode != calendar.LanguageCode {
		t.Errorf("calendar saved incorrectly, some fields doesn't match")
	}

	calendars, err = storage.GetAllTrackedCalendars()
	if err != nil {
		t.Errorf("error getting all calendars: %v", err)
	}
	if len(calendars) != 1 {
		t.Errorf("expected len(calendars): 1, got: %d", len(calendars))
	}

	got, err = storage.GetTrackedCalendar(calendar.GroupID, calendar.LanguageCode)
	if err != nil {
		t.Errorf("error getting calendar: %v", err)
	}

	if got.ID == "" || got.GroupID != calendar.GroupID || got.LanguageCode != calendar.LanguageCode {
		t.Errorf("got calendar is incorrect, some fields doesn't match")
	}

	err = storage.DeleteTrackedCalendar(got.ID)
	if err != nil {
		t.Errorf("error deleting tracked calendar: %v", err)
	}

	calendars, err = storage.GetAllTrackedCalendars()
	if err != nil {
		t.Errorf("error getting all calendars: %v", err)
	}
	if len(calendars) != 0 {
		t.Errorf("expected len(calendars): 0, got: %d", len(calendars))
	}
}
