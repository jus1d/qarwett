package postgres

func (s *Storage) CreateTrackedCalendar(groupID int64, languageCode string) (string, error) {
	var id string

	query := "INSERT INTO calendars (group_id, language_code) VALUES ($1, $2) RETURNING id"
	row := s.db.QueryRow(query, groupID, languageCode)

	if row.Err() != nil {
		return "", row.Err()
	}

	err := row.Scan(&id)

	return id, err
}
