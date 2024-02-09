package postgres

import (
	"qarwett/internal/config"
	"testing"
)

func TestStorage_User(t *testing.T) {
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

	user := User{
		TelegramID:   77991100,
		Username:     "username",
		FirstName:    "Name",
		LastName:     "Lastname",
		LanguageCode: "en",
	}

	got, err := storage.CreateUser(user.TelegramID, user.Username, user.FirstName, user.LastName, user.LanguageCode)
	if err != nil {
		t.Errorf("error creating user: %v", err)
	}

	if got.ID == "" {
		t.Errorf("incorrect user's uuid: '%s'", got.ID)
	}

	if user.TelegramID != got.TelegramID || user.Username != got.Username || user.FirstName != got.FirstName || user.LastName != got.LastName || user.LanguageCode != got.LanguageCode {
		t.Errorf("user saved incorrectly, some fields doesn't match")
	}

	got, err = storage.GetUserByTelegramID(user.TelegramID)
	if err != nil {
		t.Errorf("error getting user: %v", err)
	}

	if user.TelegramID != got.TelegramID || user.Username != got.Username || user.FirstName != got.FirstName || user.LastName != got.LastName || user.LanguageCode != got.LanguageCode {
		t.Errorf("got user is incorrect, some fields doesn't match")
	}

	err = storage.DeleteUser(user.TelegramID)
	if err != nil {
		t.Errorf("error deleting user: %v", err)
	}

	got, err = storage.GetUserByTelegramID(user.TelegramID)
	if err == nil {
		t.Errorf("error should be. user was not in db")
	}
}
