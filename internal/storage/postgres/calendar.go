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

func (s *Storage) GetTrackedCalendar(groupID int64, languageCode string) (*Calendar, error) {
	var calendar Calendar
	err := s.db.QueryRow("SELECT * FROM calendars WHERE group_id = $1 and language_code = $1", groupID, languageCode).Scan(&calendar.ID, &calendar.GroupID, &calendar.LanguageCode, &calendar.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &calendar, nil
}

func (s *Storage) GetAllTrackedCalendars() ([]Calendar, error) {
	var calendars []Calendar
	rows, err := s.db.Query("SELECT * FROM calendars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var calendar Calendar
		err = rows.Scan(&calendar.ID, &calendar.GroupID, &calendar.LanguageCode, &calendar.LanguageCode)
		if err != nil {
			return nil, err
		}
		calendars = append(calendars, calendar)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return calendars, nil
}