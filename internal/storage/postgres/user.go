package postgres

var announcements = map[int64]string{}

// CreateUser creates a user instance in storage.
func (s *Storage) CreateUser(telegramID int64, username string, firstname string, lastname string, languageCode string) (*User, error) {
	query := "INSERT INTO users (telegram_id, username, firstname, lastname, language_code) VALUES ($1, $2, $3, $4, $5) RETURNING *"
	row := s.db.QueryRow(query, telegramID, username, firstname, lastname, languageCode)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user User
	err := row.Scan(&user.ID, &user.TelegramID, &user.Username, &user.FirstName, &user.LastName, &user.Stage, &user.LinkedGroupID, &user.LinkedGroupTitle, &user.LanguageCode, &user.IsAdmin, &user.CreatedAt)

	return &user, err
}

// DeleteUser deletes user from storage by it's telegram ID.
func (s *Storage) DeleteUser(telegramID int64) error {
	query := "DELETE FROM users WHERE telegram_id = $1"
	_, err := s.db.Exec(query, telegramID)
	return err
}

// GetUserByTelegramID returns a *User, grabbed from storage by telegramID.
func (s *Storage) GetUserByTelegramID(telegramID int64) (*User, error) {
	var user User
	err := s.db.QueryRow("SELECT * FROM users WHERE telegram_id = $1", telegramID).Scan(&user.ID, &user.TelegramID, &user.Username, &user.FirstName, &user.LastName, &user.Stage, &user.LinkedGroupID, &user.LinkedGroupTitle, &user.LanguageCode, &user.IsAdmin, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers returns an array of User, all users saved to storage.
func (s *Storage) GetAllUsers() ([]User, error) {
	var users []User
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.TelegramID, &user.Username, &user.FirstName, &user.LastName, &user.Stage, &user.LinkedGroupID, &user.LinkedGroupTitle, &user.LanguageCode, &user.IsAdmin, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUserStage finds a user by telegram ID, and update its field Stage.
func (s *Storage) UpdateUserStage(telegramID int64, stage int) error {
	_, err := s.db.Exec("UPDATE users SET stage = $2 WHERE telegram_id = $1", telegramID, stage)
	return err
}

// UpdateUserLanguage finds a user by telegram ID, and update its field LanguageCode.
func (s *Storage) UpdateUserLanguage(telegramID int64, languageCode string) error {
	_, err := s.db.Exec("UPDATE users SET language_code = $2 WHERE telegram_id = $1", telegramID, languageCode)
	return err
}

// IsUserExists returns a boolean value of user existence.
func (s *Storage) IsUserExists(telegramID int64) bool {
	user, err := s.GetUserByTelegramID(telegramID)
	return err == nil && user != nil
}

// SetAnnouncementMessage saves announcement message to user's cache. After app reload, cache will be cleared.
func (s *Storage) SetAnnouncementMessage(telegramID int64, content string) {
	announcements[telegramID] = content
}

// GetAnnouncementMessage return's an announcement message, saved for this user.
func (s *Storage) GetAnnouncementMessage(telegramID int64) (string, bool) {
	val, exists := announcements[telegramID]
	return val, exists
}

func (s *Storage) UpdateUserLinkedGroup(telegramID int64, groupID int64, groupTitle string) error {
	_, err := s.db.Exec("UPDATE users SET linked_group_id = $2, linked_group_title = $3 WHERE telegram_id = $1", telegramID, groupID, groupTitle)
	return err
}
