package postgres

// CreateTrackedCalendar creates a new calendar in storage, that should stay updated.
func (s *Storage) CreateTrackedCalendar(groupID int64, languageCode string) (*Calendar, error) {
	query := "INSERT INTO calendars (group_id, language_code) VALUES ($1, $2) RETURNING *"
	row := s.db.QueryRow(query, groupID, languageCode)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var calendar Calendar
	err := row.Scan(&calendar.ID, &calendar.GroupID, &calendar.LanguageCode, &calendar.CreatedAt)

	return &calendar, err
}

func (s *Storage) DeleteTrackedCalendar(id string) error {
	query := "DELETE FROM calendars WHERE id = $1"
	_, err := s.db.Exec(query, id)
	return err
}

// GetTrackedCalendar returns tracked calendar by its group ID and language code.
func (s *Storage) GetTrackedCalendar(groupID int64, languageCode string) (*Calendar, error) {
	var calendar Calendar
	err := s.db.QueryRow("SELECT * FROM calendars WHERE group_id = $1 and language_code = $2", groupID, languageCode).Scan(&calendar.ID, &calendar.GroupID, &calendar.LanguageCode, &calendar.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &calendar, nil
}

// GetAllTrackedCalendars returns all tracked calendars from storage.
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
