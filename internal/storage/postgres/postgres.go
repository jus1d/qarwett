package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"qarwett/internal/config"
	"time"
)

// Storage provides a data structure to access the storage.
type Storage struct {
	db *sqlx.DB
}

const (
	StageNone = iota
	StageWaitingAnnouncementMessage
)

// User provides a data structure of bot's user.
type User struct {
	ID            string    `db:"id"`
	TelegramID    int64     `db:"telegram_id"`
	Username      string    `db:"username"`
	FirstName     string    `db:"firstname"`
	LastName      string    `db:"lastname"`
	Stage         int       `db:"stage"`
	LinkedGroupID int64     `db:"linked_group_id"`
	LanguageCode  string    `db:"language_code"`
	IsAdmin       bool      `db:"is_admin"`
	CreatedAt     time.Time `db:"created_at"`
}

// New creates a new instance of Storage, and returns a pointer to it.
func New(cfg config.Postgres) (*Storage, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password, cfg.ModeSSL))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, err
}

// Close is implementation for graceful shutdown. Closes a database connection.
func (s *Storage) Close() error {
	return s.db.Close()
}
