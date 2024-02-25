package postgres

import (
	"qarwett/internal/config"
	"testing"
)

func TestStorage_User(t *testing.T) {
	incorrect := config.Postgres{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Name:     "postgres",
		Password: "incorrect",
		ModeSSL:  "disable",
	}

	_, err := New(incorrect)
	if err == nil {
		t.Log("Database should not be created")
	}

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

	defer func(storage *Storage) {
		err = storage.Close()
		if err != nil {
			t.Errorf("error closing db connection: %v", err)
		}
	}(storage)

	users, err := storage.GetAllUsers()
	if err != nil {
		t.Errorf("error getting users: %v", err)
	}
	initialUsersLength := len(users)

	user := User{
		TelegramID:   77991100,
		Username:     "username",
		FirstName:    "Name",
		LastName:     "Lastname",
		LanguageCode: "en",
	}

	exists := storage.IsUserExists(user.TelegramID)
	if exists {
		t.Errorf("user should not be exist")
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
	if got.Stage != StageNone {
		t.Errorf("user's stage is incorrect: expected: 0, got: %d", got.Stage)
	}

	exists = storage.IsUserExists(user.TelegramID)
	if !exists {
		t.Errorf("user should exist")
	}

	err = storage.UpdateUserLinkedGroup(user.TelegramID, 779900000, "6101-010101D")
	if err != nil {
		t.Errorf("error updating user's linked group: %v", err)
	}

	users, err = storage.GetAllUsers()
	if err != nil {
		t.Errorf("error getting users: %v", err)
	}
	if len(users) != initialUsersLength+1 {
		t.Errorf("expected len(users): 1, got: %d", len(users))
	}

	err = storage.UpdateUserStage(user.TelegramID, StageWaitingAnnouncementMessage)
	if err != nil {
		t.Errorf("error updating user's stage: %v", err)
	}

	got, err = storage.GetUserByTelegramID(user.TelegramID)
	if err != nil {
		t.Errorf("error getting user: %v", err)
	}
	if user.TelegramID != got.TelegramID || user.Username != got.Username || user.FirstName != got.FirstName || user.LastName != got.LastName || user.LanguageCode != got.LanguageCode {
		t.Errorf("got user is incorrect, some fields doesn't match")
	}
	if got.Stage != StageWaitingAnnouncementMessage {
		t.Errorf("user's stage is incorrect: expected: 1, got: %d", user.Stage)
	}
	if got.LinkedGroupID != 779900000 || got.LinkedGroupTitle != "6101-010101D" {
		t.Errorf("incorrect linked group: expected ID: %d, got: %d\nexpected title: %s, git: %s", 779900000, got.LinkedGroupID, "6101-010101D", got.LinkedGroupTitle)
	}

	err = storage.DeleteUser(user.TelegramID)
	if err != nil {
		t.Errorf("error deleting user: %v", err)
	}

	users, err = storage.GetAllUsers()
	if err != nil {
		t.Errorf("error getting users: %v", err)
	}
	if len(users) != initialUsersLength {
		t.Errorf("expected len(users): 0, got: %d", len(users))
	}

	got, err = storage.GetUserByTelegramID(user.TelegramID)
	if err == nil {
		t.Errorf("error should be. user was not in db")
	}

	content := "announcement message content"
	storage.SetAnnouncementMessage(77991100, content)
	gotContent, exists := storage.GetAnnouncementMessage(77991100)
	if gotContent != content || !exists {
		t.Errorf("error with saving announcement message")
	}
}
