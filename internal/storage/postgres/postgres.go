package postgres

import (
	"errors"
	"fmt"

	"github.com/AndroDeMohawk/sso-app/internal/domain/models"
	"github.com/AndroDeMohawk/sso-app/internal/storage"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

type Storage struct {
	// Используем пул соединений pgxpool вместо sql.DB
	pool *pgxpool.Pool
}

// New принимает уже инициализированный пул.
// Инициализацию (pgxpool.New) лучше выносить в main.go
func New(dbPath string) *Storage {
	pool, err := pgxpool.New(context.Background(), dbPath)
	if err != nil {
		panic(err)
	}
	return &Storage{pool: pool}
}
func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = "storage.SaveUser"

	query := `INSERT INTO users (email, pash_hash) VALUES ($1, $2) RETURNING id`

	var id int64
	// Выполняем запрос напрямую через пул
	err := s.pool.QueryRow(ctx, query, email, passHash).Scan(&id)
	if err != nil {
		// Проверяем, является ли ошибка нарушением уникальности Postgres (Unique Violation)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			// Возвращаем вашу стандартную бизнес-ошибку, которую ждет сервис
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const op = "storage.User"

	query := `SELECT id, email, pash_hash FROM users WHERE email = $1`

	var user models.User
	err := s.pool.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		// В pgx ошибка отсутствия строк называется pgx.ErrNoRows
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "storage.IsAdmin"

	query := `SELECT is_admin FROM users WHERE id = $1`

	var isAdmin bool
	err := s.pool.QueryRow(ctx, query, userID).Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}

func (s *Storage) App(ctx context.Context, appID int) (models.App, error) {
	const op = "storage.App"

	query := `SELECT id, name, secret FROM apps WHERE id = $1`

	var app models.App
	err := s.pool.QueryRow(ctx, query, appID).Scan(&app.ID, &app.Name, &app.Secret)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}
