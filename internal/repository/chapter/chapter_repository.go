package chapter

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/chapter"
)

type chapterRepository struct {
	//
}

// Saving chapter likes to database.
// Need which chapterID and userID on params.
func (cr *chapterRepository) SaveChapterLikes(ctx context.Context, tx *sql.Tx, chapterID uint64, userID uint64) error {
	query := `
		INSERT INTO chapter_likes(chapter_id, user_id)
		VALUES(?, ?)
	`

	_, err := tx.ExecContext(ctx, query, chapterID, userID)
	if err != nil {
		return err
	}

	return nil
}

// Counting how many chapters the user have.
// Need the userID params.
func (cr *chapterRepository) CountChapterByUserID(ctx context.Context, tx *sql.Tx, userID uint64) (*uint64, error) {
	query := `
		SELECT
		COUNT(*)
	FROM
		chapters
	INNER JOIN stories ON chapters.story_id = stories.id
	WHERE
		stories.user_id = ?
	`

	var counts uint64
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, userID)
	err = row.Scan(&counts)
	if err != nil {
		return nil, err
	}

	return &counts, nil
}

// Counting chapters likes based on chapterID on params.
func (cr *chapterRepository) CountChapterLikesByChapterID(ctx context.Context, tx *sql.Tx, chapterID uint64) (*uint64, error) {
	query := `
		SELECT
		COUNT(*)
	FROM
		chapter_likes
	INNER JOIN chapters ON chapter_likes.chapter_id = chapters.id
	WHERE
		chapters.id = ?
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var counts uint64
	row := stmt.QueryRowContext(ctx, chapterID)
	err = row.Scan(&counts)
	if err != nil {
		return nil, err
	}

	return &counts, nil
}

// Saving chapter to database.
func (cr *chapterRepository) Save(ctx context.Context, tx *sql.Tx, c entity.Chapter) (*entity.Chapter, error) {
	query := `
		INSERT INTO chapters(
			id,
			story_id,
			title,
			slug,
			body,
			author_comment,
			word_counts,
			reading_time,
			is_published,
			created_at,
			updated_at
		)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := tx.ExecContext(ctx, query, c.Id, c.StoryID, c.Title, c.Slug, c.Body, c.AuthorComment, c.WordCounts,
		c.ReadingTime, c.IsPublished, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Updating single chapter to database.
// Need userID, storySlug and chapterSlug on params.
// Because we don't need update wrong data, that's why we need the
// userID, storySlug, chapterSlug for validation purpose.
func (cr *chapterRepository) Update(ctx context.Context, tx *sql.Tx, userID uint64, storySlug string, chapterSlug string, c entity.Chapter) (*entity.Chapter, error) {
	query := `
		UPDATE
		chapters,
		stories,
		users
	SET
		chapters.title = ?,
		chapters.slug = ?,
		chapters.body = ?,
		chapters.author_comment = ?,
		chapters.word_counts = ?,
		chapters.reading_time = ?,
		chapters.is_published = ?,
		chapters.updated_at = ?
	WHERE
		stories.user_id = ? AND stories.slug = ? AND chapters.slug = ?
	`

	_, err := tx.ExecContext(ctx, query, c.Title, c.Slug, c.Body, c.AuthorComment, c.WordCounts, c.ReadingTime,
		c.IsPublished, c.UpdatedAt, userID, storySlug, chapterSlug)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Find single chapter by userID, storySlug and chapterSlug.
// We need to get single chapter based on userId provided.
// We don't need shows chapters that not owned by the userID provided.
func (cr *chapterRepository) FindByStorySlugAndChapterSlug(ctx context.Context, tx *sql.Tx, userID uint64, storySlug string, chapterSlug string) (*entity.Chapter, error) {
	query := `
		SELECT
		chapters.*,
		stories.*,
		users.*
	FROM
		chapters
	INNER JOIN stories ON chapters.story_id = stories.id
	INNER JOIN users ON stories.user_id = users.id
	WHERE
		stories.user_id = ? AND stories.slug = ? AND chapters.slug = ?
	`

	row := tx.QueryRowContext(ctx, query, userID, storySlug, chapterSlug)

	var u entity.User
	var s entity.Story
	var c entity.Chapter
	err := row.Scan(&c.Id, &c.StoryID, &c.Title, &c.Slug, &c.Body, &c.AuthorComment, &c.WordCounts, &c.ReadingTime, &c.IsPublished,
		&c.CreatedAt, &c.UpdatedAt,

		&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover, &s.CreatedAt,
		&s.UpdatedAt,

		&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter,
		&u.Instagram, &u.Facebook, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.ErrChapterNotFound
		}

		return nil, err
	}

	c.User = u
	c.Story = s

	return &c, nil
}

// Delete single chapter.
// We need the userID and chapterID to make sure it legit his own chapter data.
// We don't need to delete another user chapter, that's why userID is provided on params.
func (cr *chapterRepository) Delete(ctx context.Context, tx *sql.Tx, userID uint64, chapterID uint64) error {
	query := `
		DELETE
		chapters
	FROM
		chapters
	INNER JOIN stories ON chapters.story_id = stories.id
	INNER JOIN users ON stories.user_id = users.id
	WHERE
		stories.user_id = ? AND chapters.id = ?
	`

	_, err := tx.ExecContext(ctx, query, userID, chapterID)
	if err != nil {
		return err
	}

	return nil
}

// Find single chapter by id.
func (cr *chapterRepository) FindByID(ctx context.Context, tx *sql.Tx, id uint64) (*entity.Chapter, error) {
	query := `
		SELECT
		chapters.*,
		stories.*,
		users.*
	FROM
		chapters
	INNER JOIN stories ON chapters.story_id = stories.id
	INNER JOIN users ON stories.user_id = users.id
	WHERE
		chapters.id = ?
	`

	row := tx.QueryRowContext(ctx, query, id)

	var c entity.Chapter
	var s entity.Story
	var u entity.User
	err := row.Scan(&c.Id, &c.StoryID, &c.Title, &c.Slug, &c.Body, &c.AuthorComment, &c.WordCounts, &c.ReadingTime,
		&c.IsPublished, &c.CreatedAt, &c.UpdatedAt,

		&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover, &s.CreatedAt,
		&s.UpdatedAt,

		&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter, &u.Instagram,
		&u.Facebook, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.ErrChapterNotFound
		}

		return nil, err
	}

	c.Story = s
	c.User = u

	return &c, nil
}

// Counting chapter by storySlug.
// This function count how many chapters the story has.
func (cr *chapterRepository) CountChapterByStorySlug(ctx context.Context, tx *sql.Tx, storySlug string) (*uint64, error) {
	query := `
		SELECT
		COUNT(*)
	FROM
		chapters
	INNER JOIN stories ON chapters.story_id = stories.id
	WHERE
		stories.slug = ?
	`

	var counts uint64
	row := tx.QueryRowContext(ctx, query, storySlug)

	err := row.Scan(&counts)
	if err != nil {
		return nil, err
	}

	return &counts, nil
}

// Find all chapters by storySlug.
// Find all chapters related by storySlug provided on params.
func (cr *chapterRepository) FindAllChapterByStorySlug(ctx context.Context, tx *sql.Tx, storySlug string) (*[]entity.Chapter, error) {
	query := `
		SELECT
		chapters.*
	FROM
		chapters
	INNER JOIN stories ON chapters.story_id = stories.id
	WHERE
		stories.slug = ?
	`

	rows, err := tx.QueryContext(ctx, query, storySlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chapters []entity.Chapter
	for rows.Next() {
		var c entity.Chapter
		err := rows.Scan(&c.Id, &c.StoryID, &c.Title, &c.Slug, &c.Body, &c.AuthorComment, &c.WordCounts, &c.ReadingTime,
			&c.IsPublished, &c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		chapters = append(chapters, c)
	}

	return &chapters, nil
}

// Find all chapters by storyID.
// We find all chapters using storyID that provided on params.
func (cr *chapterRepository) FindAllChapterByStoryID(ctx context.Context, tx *sql.Tx, storyID uint64) (*[]entity.Chapter, error) {
	query := `
		SELECT
		chapters.*
	FROM
		chapters
	INNER JOIN stories ON chapters.story_id = stories.id
	WHERE
		stories.id = ?
	`

	rows, err := tx.QueryContext(ctx, query, storyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chapters []entity.Chapter
	for rows.Next() {
		var c entity.Chapter
		err := rows.Scan(&c.Id, &c.StoryID, &c.Title, &c.Slug, &c.Body, &c.AuthorComment, &c.WordCounts, &c.ReadingTime,
			&c.IsPublished, &c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		chapters = append(chapters, c)
	}

	return &chapters, nil
}

func NewRepository() ChapterRepository {
	return &chapterRepository{}
}
