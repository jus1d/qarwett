package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"qarwett/internal/config"
	"time"
)

type Storage struct {
	db *sqlx.DB
}

type User struct {
	ID            string    `db:"id"`
	TelegramID    int64     `db:"telegram_id"`
	Username      string    `db:"username"`
	FirstName     string    `db:"firstname"`
	LastName      string    `db:"lastname"`
	LinkedGroupID int64     `db:"linked_group_id"`
	LanguageCode  string    `db:"language_code"`
	CreatedAt     time.Time `db:"created_at"`
}

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

func (s *Storage) Close() error {
	return s.db.Close()
}
