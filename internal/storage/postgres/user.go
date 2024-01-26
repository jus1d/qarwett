package postgres

import "fmt"

func (s *Storage) CreateUser(telegramID int64, username string, firstname string, lastname string, languageCode string) (string, error) {
	var id string

	query := "INSERT INTO users (telegram_id, username, firstname, lastname, language_code) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := s.db.QueryRow(query, telegramID, username, firstname, lastname, languageCode)

	if row.Err() != nil {
		return "", row.Err()
	}

	err := row.Scan(&id)

	return id, err
}

func (s *Storage) GetUserByTelegramID(telegramID int64) (*User, error) {
	var user User
	err := s.db.QueryRow("SELECT * FROM users WHERE telegram_id = $1", telegramID).Scan(&user.ID, &user.TelegramID, &user.Username, &user.FirstName, &user.LastName, &user.LinkedGroupID, &user.LanguageCode, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	return &user, nil
}

func (s *Storage) IsUserExists(telegramID int64) bool {
	user, err := s.GetUserByTelegramID(telegramID)
	return err == nil && user != nil
}
