package story

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/story"
)

type storyRepository struct {
	//
}

// Filtering latest modified chapter by category slug.
func (sr *storyRepository) FilterLatestModifiedChapterByCategorySlug(ctx context.Context, tx *sql.Tx, categorySlug string) (*[]entity.Story, error) {
	query := `
		SELECT
		stories.*,
		users.*,
		MAX(chapters.updated_at)
	FROM
		stories
	INNER JOIN chapters ON chapters.story_id = stories.id
	INNER JOIN users ON stories.user_id = users.id
	INNER JOIN categories ON stories.category_id = categories.id
	WHERE
		categories.slug = ?
	GROUP BY
		chapters.story_id
	ORDER BY
		chapters.updated_at
	DESC
	`

	rows, err := tx.QueryContext(ctx, query, categorySlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ignore := new(any)
	var stories []entity.Story
	for rows.Next() {
		var s entity.Story
		var u entity.User
		err := rows.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover, &s.CreatedAt, &s.UpdatedAt,

			&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter,
			&u.Instagram, &u.Facebook, &u.CreatedAt, &ignore)
		if err != nil {
			return nil, err
		}

		s.User = u
		stories = append(stories, s)
	}

	return &stories, nil
}

// Filter latest story based category slug.
func (sr *storyRepository) FilterLatestBasedOnCategorySlug(ctx context.Context, tx *sql.Tx, categorySlug string) (*[]entity.Story, error) {
	query := `
		SELECT
		stories.*,
		users.*
	FROM
		stories
	INNER JOIN users ON stories.user_id = users.id
	INNER JOIN categories ON stories.category_id = categories.id
	WHERE
		categories.slug = ?
	ORDER BY
		stories.created_at
	DESC
	`

	rows, err := tx.QueryContext(ctx, query, categorySlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []entity.Story
	for rows.Next() {
		var s entity.Story
		var u entity.User
		err := rows.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover,
			&s.CreatedAt, &s.UpdatedAt,

			&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter,
			&u.Instagram, &u.Facebook, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

		s.User = u
		stories = append(stories, s)
	}

	return &stories, nil
}

// Find story by category slug.
func (sr *storyRepository) FindByCategorySlug(ctx context.Context, tx *sql.Tx, categorySlug string) (*[]entity.Story, error) {
	query := `
		SELECT
		stories.*,
		users.*
	FROM
		stories
	INNER JOIN users ON stories.user_id = users.id
	INNER JOIN categories ON stories.category_id = categories.id
	WHERE
		categories.slug = ?
	`

	rows, err := tx.QueryContext(ctx, query, categorySlug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []entity.Story
	for rows.Next() {
		var s entity.Story
		var u entity.User
		err := rows.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover,
			&s.CreatedAt, &s.UpdatedAt,

			&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter,
			&u.Instagram, &u.Facebook, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

		s.User = u
		stories = append(stories, s)
	}

	return &stories, nil
}

// Counting story by category ID.
// Need the categoryID on params.
func (sr *storyRepository) CountStoryByCategoryID(ctx context.Context, tx *sql.Tx, categoryID uint64) (*uint64, error) {
	query := `
		SELECT
		COUNT(*)
	FROM
		stories
	WHERE
		category_id = ?
	`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var counts uint64
	row := stmt.QueryRowContext(ctx, categoryID)
	err = row.Scan(&counts)
	if err != nil {
		return nil, err
	}

	return &counts, nil
}

// Filtering latest updated/modified chapter.
func (sr *storyRepository) FilterLatestModifiedChapter(ctx context.Context, tx *sql.Tx) (*[]entity.Story, error) {
	query := `
		SELECT
		stories.*,
		users.*,
		MAX(chapters.updated_at)
	FROM
		stories
	JOIN chapters ON chapters.story_id = stories.id
	JOIN users ON stories.user_id = users.id
	GROUP BY
		chapters.story_id
	ORDER BY
		chapters.updated_at
	DESC
	`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ignore := new(any)
	var stories []entity.Story
	for rows.Next() {
		var s entity.Story
		var u entity.User
		err := rows.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover, &s.CreatedAt, &s.UpdatedAt,

			&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter,
			&u.Instagram, &u.Facebook, &u.CreatedAt, &ignore)
		if err != nil {
			return nil, err
		}

		s.User = u
		stories = append(stories, s)
	}

	return &stories, nil
}

// Filter latest created story.
func (sr *storyRepository) FilterLatest(ctx context.Context, tx *sql.Tx) (*[]entity.Story, error) {
	query := `
		SELECT
		stories.*,
		users.*
	FROM
		stories
	INNER JOIN users ON stories.user_id = users.id
	ORDER BY
		stories.created_at
	DESC
	`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []entity.Story
	for rows.Next() {
		var s entity.Story
		var u entity.User

		err := rows.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover,
			&s.CreatedAt, &s.UpdatedAt,

			&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter,
			&u.Instagram, &u.Facebook, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

		s.User = u
		stories = append(stories, s)
	}

	return &stories, nil
}

// Find all story by user ID.
// We get all the story based user that has the story.
// We don't need to shows another story that the story does not owned by userID provided.
// That's why we need the userID on params.
func (sr *storyRepository) FindAllByUserID(ctx context.Context, tx *sql.Tx, userID uint64) (*[]entity.Story, error) {
	query := `
		SELECT
		stories.*,
		users.id,
		users.name,
		users.email
	FROM
		stories
	INNER JOIN users ON stories.user_id = users.id
	WHERE
		user_id = ?
	`

	rows, err := tx.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []entity.Story
	for rows.Next() {
		var u entity.User
		var s entity.Story
		err := rows.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover,
			&s.CreatedAt, &s.UpdatedAt,

			&u.Id, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}

		s.User = u
		stories = append(stories, s)
	}

	return &stories, nil
}

// Find single story by ID.
// If the ID does not exists on database, throwing an err story not found.
func (sr *storyRepository) FindByID(ctx context.Context, tx *sql.Tx, id uint64) (*entity.Story, error) {
	query := `
		SELECT
		*
	FROM
		stories
	WHERE
		id = ?
	`

	var s entity.Story
	row := tx.QueryRowContext(ctx, query, id)

	err := row.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished, &s.Cover,
		&s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.ErrStoryNotFound
		}

		return nil, err
	}

	story, err := sr.load(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// Delete single story by ID.
func (sr *storyRepository) Delete(ctx context.Context, tx *sql.Tx, id uint64) error {
	query := `
		DELETE
		FROM
			stories
		WHERE
			id = ?
	`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// Find single story by slug and userID.
// If not exists, thworing an err story not found.
func (sr *storyRepository) FindBySlugAndUserID(ctx context.Context, tx *sql.Tx, slug string, userID uint64) (*entity.Story, error) {
	query := `
		SELECT
		stories.*,
		categories.*,
		users.*
	FROM
		stories
	INNER JOIN categories ON stories.category_id = categories.id
	INNER JOIN users ON stories.user_id = users.id
	WHERE
		stories.slug = ? AND stories.user_id = ?
	`

	row := tx.QueryRowContext(ctx, query, slug, userID)

	var s entity.Story
	var c entity.Category
	var u entity.User
	err := row.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult,
		&s.IsPublished, &s.Cover, &s.CreatedAt, &s.UpdatedAt,

		&c.Id, &c.Name, &c.Slug,

		&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter,
		&u.Instagram, &u.Facebook, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.ErrStoryNotFound
		}

		return nil, err
	}

	s.User = u
	s.Category = c
	return &s, nil
}

// Updating single story by slug.
func (sr *storyRepository) Update(ctx context.Context, tx *sql.Tx, slug string, s entity.Story) (*entity.Story, error) {
	query := `
		UPDATE
		stories
	SET
		user_id = ?,
		category_id = ?,
		title = ?,
		slug = ?,
		description = ?,
		is_adult = ?,
		is_published = ?,
		cover = ?,
		updated_at = ?
	WHERE
		slug = ? AND user_id = ?
	`
	_, err := tx.ExecContext(ctx, query, s.UserID, s.CategoryID, s.Title, s.Slug, s.Description, s.IsAdult, s.IsPublished,
		s.Cover, s.UpdatedAt, slug, s.UserID)
	if err != nil {
		return nil, err
	}

	story, err := sr.load(ctx, tx, s.Id)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// Saving story into database.
func (sr *storyRepository) Save(ctx context.Context, tx *sql.Tx, s entity.Story) (*entity.Story, error) {
	query := `
		INSERT INTO stories(
			id,
			user_id,
			category_id,
			title,
			slug,
			description,
			is_adult,
			is_published,
			cover,
			created_at,
			updated_at
		)
		VALUES(
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)
	`

	_, err := tx.ExecContext(ctx, query, s.Id, s.UserID, s.CategoryID, s.Title, s.Slug, s.Description, s.IsAdult,
		s.IsPublished, s.Cover, s.CreatedAt, s.UpdatedAt)
	if err != nil {
		return nil, err
	}

	story, err := sr.load(ctx, tx, s.Id)
	if err != nil {
		return nil, err
	}

	return story, nil
}

// Load relationship.
func (sr *storyRepository) load(ctx context.Context, tx *sql.Tx, storyID uint64) (*entity.Story, error) {
	query := `
		SELECT
		stories.*,
		users.id,
		users.name,
		users.email,
		categories.*
	FROM
		stories
	INNER JOIN categories ON stories.category_id = categories.id
	INNER JOIN users ON stories.user_id = users.id
	WHERE
		stories.id = ?
	`

	row, err := tx.QueryContext(ctx, query, storyID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var s entity.Story
	var u entity.User
	var c entity.Category
	for row.Next() {
		err := row.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult,
			&s.IsPublished, &s.Cover, &s.CreatedAt, &s.UpdatedAt,

			&u.Id, &u.Name, &u.Email,

			&c.Id, &c.Name, &c.Slug)
		if err != nil {
			return nil, err
		}
	}

	s.User = u
	s.Category = c

	return &s, nil
}

// Find single story by slug.
func (sr *storyRepository) FindBySlug(ctx context.Context, tx *sql.Tx, slug string) (*entity.Story, error) {
	query := `
		SELECT
		stories.*,
		categories.*,
		users.*
	FROM
		stories
	INNER JOIN categories ON stories.category_id = categories.id
	INNER JOIN users ON stories.user_id = users.id
	WHERE
		stories.slug = ?
	`
	row := tx.QueryRowContext(ctx, query, slug)

	var s entity.Story
	var u entity.User
	var c entity.Category
	err := row.Scan(&s.Id, &s.UserID, &s.CategoryID, &s.Title, &s.Slug, &s.Description, &s.IsAdult, &s.IsPublished,
		&s.Cover, &s.CreatedAt, &s.UpdatedAt,

		&c.Id, &c.Name, &c.Slug,

		&u.Id, &u.Name, &u.Username, &u.Email, &u.Password, &u.Sex, &u.Bio, &u.DateOfBirth, &u.PhoneNumber, &u.Twitter,
		&u.Instagram, &u.Facebook, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.ErrStoryNotFound
		}

		return nil, err
	}

	s.User = u
	s.Category = c

	return &s, nil
}

func NewRepository() StoryRepository {
	return &storyRepository{}
}
