package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shlmvgleb/em-task/internal/models"
)

type PostgresSongRepo struct {
	db *pgxpool.Pool
}

func NewPostgresSongRepo(db *pgxpool.Pool) *PostgresSongRepo {
	return &PostgresSongRepo{db: db}
}

func (r *PostgresSongRepo) GetWithSearchAndPagination(
	ctx context.Context,
	searchQuery string,
	limit int, offset int,
) ([]*models.Song, int, error) {
	var amount int
	query := `SELECT count(*) as amount FROM song`
	row := r.db.QueryRow(ctx, query)
	err := row.Scan(&amount)
	if err != nil {
		return nil, 0, err
	}

	var rows pgx.Rows
	if searchQuery == "" {
		query = `
			SELECT id, song, "group", "text", release_date, "link"
			FROM song
			ORDER BY created_at ASC LIMIT $1 OFFSET $2
		`
		rows, err = r.db.Query(ctx, query, limit, offset)
	} else {
		query = `
			SELECT id, song, "group", "text", release_date, "link"
			FROM song
			where to_tsvector(song || ' ' || "group" || ' ' || "text") @@ websearch_to_tsquery($1)
			ORDER BY created_at ASC LIMIT $2 OFFSET $3
		`
		rows, err = r.db.Query(ctx, query, searchQuery, limit, offset)
	}

	if err != nil {
		return nil, 0, err
	}

	songs := make([]*models.Song, 0)

	defer rows.Close()
	for rows.Next() {
		song := models.Song{}
		err = rows.Scan(&song.Id, &song.Song, &song.Group, &song.Text, &song.ReleaseDate, &song.Link)
		if err != nil {
			return nil, 0, err
		}

		songs = append(songs, &song)
	}

	return songs, amount, nil
}

func (r *PostgresSongRepo) GetById(ctx context.Context, id int64) (*models.Song, error) {
	song := models.Song{}

	query := `
		SELECT id, song, "group", "text", release_date, "link" FROM song 
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)
	err := row.Scan(&song.Id, &song.Song, &song.Group, &song.Text, &song.ReleaseDate, &song.Link)
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (r *PostgresSongRepo) Add(ctx context.Context, song *models.Song) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return
			}
		}
	}()

	query := `
		INSERT INTO song ("group", song, "text", "link", release_date)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = tx.Exec(ctx, query, song.Group, song.Song, song.Text, song.Link, song.ReleaseDate)
	if err != nil {
		return fmt.Errorf("failed to add song: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (r *PostgresSongRepo) Update(ctx context.Context, id int64, song *models.Song) (*models.Song, error) {
	prevData, err := r.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find a song to update: %w", err)
	}

	bytes, err := json.Marshal(song)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal song struct: %w", err)
	}

	err = json.Unmarshal(bytes, prevData)
	if err != nil {
		return nil, fmt.Errorf("failed to merge new data to song struct: %w", err)
	}

	song = prevData

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return
			}
		}
	}()

	query := `
		UPDATE song
		SET song = $1, "group" = $2, "link" = $3, "text" = $4, release_date = $5, updated_at = now() 
		WHERE id = $6;
	`

	_, err = tx.Exec(
		ctx,
		query,
		song.Song,
		song.Group,
		song.Link,
		song.Text,
		song.ReleaseDate,
		song.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update song: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return song, nil
}

func (r *PostgresSongRepo) Delete(ctx context.Context, id int64) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil || err != nil {
			if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
				return
			}
		}
	}()

	query := `
		DELETE FROM song WHERE id = $1
	`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete song: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
