package postgres

func (s *Storage) CreateUser(telegramID int64, username string, firstname string, lastname string, languageCode string) (string, error) {
	var id string

	query := "insert into users (telegram_id, username, firstname, lastname, language_code) values ($1, $2, $3, $4, $5) returning id"
	row := s.db.QueryRow(query, telegramID, username, firstname, lastname, languageCode)

	if row.Err() != nil {
		return "", row.Err()
	}

	err := row.Scan(&id)

	return id, err
}
